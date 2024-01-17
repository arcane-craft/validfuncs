package generic

import (
	"fmt"

	"github.com/arcane-craft/validfuncs"
	"github.com/arcane-craft/validfuncs/utils"
)

func Not[T any](expr validfuncs.Expr[T], left ...validfuncs.Expr[T]) validfuncs.Expr[T] {
	if len(left) > 0 {
		expr = All(expr, left[0], left[1:]...)
	}
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			instrs, err := expr.Interpret()
			if err != nil {
				return nil, err
			}
			return validfuncs.NewInstr(
				instrs.Exec,
				func() string {
					return "not" + instrs.String()
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Not", expr)
		},
	)
}

func All[T any](expr1, expr2 validfuncs.Expr[T], left ...validfuncs.Expr[T]) validfuncs.Expr[T] {
	exprs := validfuncs.Exprs[T](append([]validfuncs.Expr[T]{expr1, expr2}, left...))
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			instrs, err := exprs.Interpret()
			if err != nil {
				return nil, err
			}
			return validfuncs.NewInstr(
				instrs.Exec,
				func() string {
					return instrs.String()
				},
			), nil
		},
		func() string {
			return exprs.String()
		},
	)
}

func Any[T any](expr1, expr2 validfuncs.Expr[T], left ...validfuncs.Expr[T]) validfuncs.Expr[T] {
	exprs := validfuncs.Exprs[T](append([]validfuncs.Expr[T]{expr1, expr2}, left...))
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			instrs, err := exprs.Interpret()
			if err != nil {
				return nil, err
			}
			return validfuncs.NewInstr(
				func(ctx *validfuncs.Context, v T) bool {
					for _, instr := range instrs {
						if instr.Exec(ctx, v) {
							return true
						}
					}
					return false
				},
				func() string {
					return fmt.Sprintf("any of {%s}", instrs.String())
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Any", exprs)
		},
	)
}
