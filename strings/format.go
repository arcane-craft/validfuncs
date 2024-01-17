package strings

import (
	"net/mail"
	"strings"

	"github.com/arcane-craft/validfuncs"
	"github.com/arcane-craft/validfuncs/utils"
)

func Email() validfuncs.Expr[string] {
	return AEmail[string]()
}

func AEmail[T ~string]() validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					if strings.TrimSpace(string(v)) != string(v) {
						return false
					}
					_, err := mail.ParseAddress(string(v))
					return err == nil
				},
				func() string {
					return "email"
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Email")
		},
	)
}

func EmailAnd[T ~string](next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(AEmail[T](), next)
}

func Phone() validfuncs.Expr[string] {
	return APhone[string]()
}

func APhone[T ~string]() validfuncs.Expr[T] {
	return RangeLenAnd(3, 15, ACharset[T](Number()))
}

func PhoneAnd[T ~string](next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(APhone[T](), next)
}
