package codex

import (
	"strings"

	"github.com/quanxiaoxuan/go-utils/stringx"
)

// 构建插入语句
func BuildInsertSql(table string, fieldList FieldList) string {
	sb := strings.Builder{}
	iv := strings.Builder{}
	sb.WriteString("insert")
	sb.WriteString(" into ")
	sb.WriteString(table)
	sb.WriteString(LineSep)
	sb.WriteString("  (")
	for i, field := range fieldList {
		low := stringx.LowerCamelCase(field.Name)
		if i > 0 {
			sb.WriteString(LineSep)
			sb.WriteString("   ")
			iv.WriteString(LineSep)
			iv.WriteString("   ")
		}
		sb.WriteString(field.Name)
		sb.WriteString(",")
		iv.WriteString("#{addParam.")
		iv.WriteString(low)
		iv.WriteString("},")
	}
	sb.WriteString(",)")
	sb.WriteString(LineSep)
	sb.WriteString("values")
	sb.WriteString(LineSep)
	sb.WriteString("  (")
	sb.WriteString(iv.String())
	sb.WriteString(",)")
	sb.WriteString(LineSep)
	return strings.ReplaceAll(sb.String(), ",,", "")
}
