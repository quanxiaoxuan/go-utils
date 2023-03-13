package codex

import (
	"strings"

	"github.com/quanxiaoxuan/go-utils/stringx"
)

// 构建更新语句
func BuildUpdateSql(table string, fieldList FieldList) string {
	sb := strings.Builder{}
	sb.WriteString("update ")
	sb.WriteString(table)
	sb.WriteString(LineSep)
	sb.WriteString("   set ")
	for i, field := range fieldList {
		low := stringx.LowerCamelCase(field.Name)
		if i > 0 {
			sb.WriteString(LineSep)
			sb.WriteString("       ")
		}
		sb.WriteString(field.Name)
		sb.WriteString(" = #{updateParam.")
		sb.WriteString(low)
		sb.WriteString("},")
	}
	sb.WriteString(",")
	sb.WriteString(LineSep)
	sb.WriteString(" where id = #{updateParam.id} ")
	sb.WriteString(LineSep)
	return strings.ReplaceAll(sb.String(), ",,", "")
}
