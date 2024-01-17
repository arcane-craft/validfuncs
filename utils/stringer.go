package utils

import (
	"fmt"
	"strings"
)

func stringifyArgs(args ...any) []string {
	var strArgs []string
	for _, a := range args {
		switch v := any(a).(type) {
		case rune:
			strArgs = append(strArgs, fmt.Sprintf("'%c'", v))
		case []rune:
			strArgs = append(strArgs, stringifyArgs(v)...)
		case string:
			strArgs = append(strArgs, fmt.Sprintf("\"%s\"", v))
		case []string:
			strArgs = append(strArgs, v...)
		case fmt.Stringer:
			strArgs = append(strArgs, v.String())
		case []fmt.Stringer:
			strArgs = append(strArgs, stringifyArgs(v)...)
		default:
			strArgs = append(strArgs, fmt.Sprintf("%v", v))
		}
	}
	return strArgs
}

func StringifyExpr(name string, args ...any) string {
	strArgs := stringifyArgs(args...)
	if len(strArgs) > 0 {
		return fmt.Sprintf("(%s %s)", name, strings.Join(strArgs, " "))
	}
	return fmt.Sprintf("(%s)", name)
}
