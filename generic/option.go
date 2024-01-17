package generic

import (
	"github.com/arcane-craft/validfuncs"
	"github.com/arcane-craft/validfuncs/utils"
)

func Ref[P *V, V any](expr validfuncs.Expr[V], left ...validfuncs.Expr[V]) validfuncs.Expr[P] {
	if len(left) > 0 {
		expr = All(expr, left[0], left[1:]...)
	}
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[P], error) {
			instrs, err := expr.Interpret()
			if err != nil {
				return nil, err
			}
			var nilPointer bool
			return validfuncs.NewInstr[P](
				func(ctx *validfuncs.Context, v P) bool {
					if v == nil {
						nilPointer = true
						return false
					}
					return instrs.Exec(ctx, *v)
				},
				func() string {
					if nilPointer {
						return "not nil"
					}
					return instrs.String()
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Ref", expr)
		},
	)
}

func OmitNil[P *V, V any](expr validfuncs.Expr[P]) validfuncs.Expr[P] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[P], error) {
			instrs, err := expr.Interpret()
			if err != nil {
				return nil, err
			}
			return validfuncs.NewInstr[P](
				func(ctx *validfuncs.Context, v P) bool {
					if v == nil {
						return true
					}
					return instrs.Exec(ctx, v)
				},
				func() string {
					return instrs.String()
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("OmitNil", expr)
		},
	)
}

func OmitEmpty[T comparable](expr validfuncs.Expr[T], left ...validfuncs.Expr[T]) validfuncs.Expr[T] {
	if len(left) > 0 {
		expr = All(expr, left[0], left[1:]...)
	}
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			instrs, err := expr.Interpret()
			if err != nil {
				return nil, err
			}
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					var empty T
					if v == empty {
						return true
					}
					return instrs.Exec(ctx, v)
				},
				func() string {
					return instrs.String()
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("OmitEmpty", expr)
		},
	)
}
