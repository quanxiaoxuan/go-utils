package codex

import (
	"strings"

	"github.com/quanxiaoxuan/go-utils/stringx"
)

// 不带字段别名
func BuildSelectSql(table string, fieldList FieldList) string {
	sb := strings.Builder{}
	sb.WriteString("select ")
	for i, field := range fieldList {
		if i > 0 {
			sb.WriteString(LineSep)
			sb.WriteString("       ")
		}
		sb.WriteString(field.Name)
		sb.WriteString(",")
	}
	sb.WriteString(",")
	sb.WriteString(LineSep)
	sb.WriteString("  from ")
	sb.WriteString(table)
	sb.WriteString(LineSep)
	return strings.ReplaceAll(sb.String(), ",,", "")
}

// 带字段别名
func BuildSelectSqlAlias(table string, columns FieldList) string {
	sb := strings.Builder{}
	sb.WriteString("select ")
	for i, field := range columns {
		low := stringx.LowerCamelCase(field.Name)
		if i > 0 {
			sb.WriteString(LineSep)
			sb.WriteString("       ")
		}
		sb.WriteString(field.Name)
		sb.WriteString(" as ")
		sb.WriteString(low)
		sb.WriteString(",")
	}
	sb.WriteString(",")
	sb.WriteString(LineSep)
	sb.WriteString("  from ")
	sb.WriteString(table)
	sb.WriteString(LineSep)
	return strings.ReplaceAll(sb.String(), ",,", "")
}
