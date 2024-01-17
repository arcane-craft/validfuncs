package strings

import (
	"fmt"

	"github.com/arcane-craft/validfuncs"
	"github.com/arcane-craft/validfuncs/generic"
	"github.com/arcane-craft/validfuncs/utils"
)

func Charset(expr validfuncs.Expr[rune], left ...validfuncs.Expr[rune]) validfuncs.Expr[string] {
	return ACharset[string](expr, left...)
}

func ACharset[T ~string](expr validfuncs.Expr[rune], left ...validfuncs.Expr[rune]) validfuncs.Expr[T] {
	exprs := validfuncs.Exprs[rune](append([]validfuncs.Expr[rune]{expr}, left...))
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			instrs, err := exprs.Interpret()
			if err != nil {
				return nil, err
			}
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					for _, v := range v {
						var ok bool
						for _, instr := range instrs {
							ok = instr.Exec(ctx, v)
							if ok {
								break
							}
						}
						if !ok {
							return false
						}
					}
					return true
				},
				func() string {
					return fmt.Sprintf("charset {%s}", instrs)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Charset", expr)
		},
	)
}

func CharsetAnd[T ~string](exprs *utils.VarArgs[utils.N1, validfuncs.Expr[rune]], next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(ACharset[T](exprs.FixedArgs()[0], exprs.LeftArgs()...), next)
}

func Char(idx int, expr validfuncs.Expr[rune], left ...validfuncs.Expr[rune]) validfuncs.Expr[string] {
	return AChar[string](idx, expr, left...)
}

func AChar[T ~string | []rune](idx int, expr validfuncs.Expr[rune], left ...validfuncs.Expr[rune]) validfuncs.Expr[T] {
	if len(left) > 0 {
		expr = generic.Any(expr, left[0], left[1:]...)
	}
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			instrs, err := expr.Interpret()
			if err != nil {
				return nil, err
			}
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					if idx >= len(v) {
						return false
					}
					return validfuncs.WithSelector(fmt.Sprintf("[%d]", idx), instrs).Exec(ctx, []rune(v)[idx])
				},
				func() string {
					return fmt.Sprintf("char at index %d", idx)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("One", expr)
		},
	)
}

func CharAnd[T ~string | []rune](idx int, exprs *utils.VarArgs[utils.N1, validfuncs.Expr[rune]], next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(AChar[T](idx, exprs.FixedArgs()[0], exprs.LeftArgs()...), next)
}
