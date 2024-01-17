package slices

import (
	"fmt"

	"github.com/arcane-craft/validfuncs"
	"github.com/arcane-craft/validfuncs/generic"
	"github.com/arcane-craft/validfuncs/utils"
)

func Each[T []E, E any](expr validfuncs.Expr[E], left ...validfuncs.Expr[E]) validfuncs.Expr[T] {
	return AEach[T, E](expr, left...)
}

func AEach[T ~[]E, E any](expr validfuncs.Expr[E], left ...validfuncs.Expr[E]) validfuncs.Expr[T] {
	if len(left) > 0 {
		expr = generic.All(expr, left[0], left[1:]...)
	}
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			instrs, err := expr.Interpret()
			if err != nil {
				return nil, err
			}
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					for i, v := range v {
						if !validfuncs.WithSelector(fmt.Sprintf("[%d]", i), instrs).Exec(ctx, v) {
							return false
						}
					}
					return true
				},
				func() string {
					return instrs.String()
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Each", expr)
		},
	)
}

func EachAnd[T ~[]E, E any](exprs *utils.VarArgs[utils.N1, validfuncs.Expr[E]], next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(AEach[T](exprs.FixedArgs()[0], exprs.LeftArgs()...), next)
}

func One[T []E, E any](idx int, expr validfuncs.Expr[E], left ...validfuncs.Expr[E]) validfuncs.Expr[T] {
	return AOne[T, E](idx, expr, left...)
}

func AOne[T ~[]E, E any](idx int, expr validfuncs.Expr[E], left ...validfuncs.Expr[E]) validfuncs.Expr[T] {
	if len(left) > 0 {
		expr = generic.All(expr, left[0], left[1:]...)
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
					return validfuncs.WithSelector(fmt.Sprintf("[%d]", idx), instrs).Exec(ctx, v[idx])
				},
				func() string {
					return fmt.Sprintf("at index %d", idx)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("One", expr)
		},
	)
}

func OneAnd[T ~[]E, E any](idx int, exprs *utils.VarArgs[utils.N1, validfuncs.Expr[E]], next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(AOne[T](idx, exprs.FixedArgs()[0], exprs.LeftArgs()...), next)
}
