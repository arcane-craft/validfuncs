package strings

import (
	"unicode"

	"github.com/arcane-craft/validfuncs"
	"github.com/arcane-craft/validfuncs/utils"
)

func Space() validfuncs.Expr[rune] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[rune], error) {
			return validfuncs.NewInstr(
				func(ctx *validfuncs.Context, v rune) bool {
					return unicode.IsSpace(v)
				},
				func() string {
					return "space"
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Space")
		},
	)
}

func Number() validfuncs.Expr[rune] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[rune], error) {
			return validfuncs.NewInstr(
				func(ctx *validfuncs.Context, v rune) bool {
					return (v >= '0' && v <= '9')
				},
				func() string {
					return "number"
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Number")
		},
	)
}

func AlphaNum() validfuncs.Expr[rune] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[rune], error) {
			return validfuncs.NewInstr(
				func(ctx *validfuncs.Context, v rune) bool {
					return (v >= '0' && v <= '9') ||
						(v >= 'A' && v <= 'Z') ||
						(v >= 'a' && v <= 'z')
				},
				func() string {
					return "alpha numeric"
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("AlphaNum")
		},
	)
}

func PrintASCII() validfuncs.Expr[rune] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[rune], error) {
			return validfuncs.NewInstr(
				func(ctx *validfuncs.Context, v rune) bool {
					return v >= 0x21 && v <= 0x7e
				},
				func() string {
					return "printable ASCII"
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("PrintASCII")
		},
	)
}

func ExcludesChars(chars []rune) validfuncs.Expr[rune] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[rune], error) {
			return validfuncs.NewInstr(
				func(ctx *validfuncs.Context, v rune) bool {
					for _, c := range chars {
						if c == v {
							return true
						}
					}
					return false
				},
				func() string {
					return "without chars"
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("ExcludesChar", chars)
		},
	)
}
