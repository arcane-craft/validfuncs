package maps

import (
	"fmt"

	"github.com/arcane-craft/validfuncs"
	"github.com/arcane-craft/validfuncs/generic"
	"github.com/arcane-craft/validfuncs/utils"
)

type Entry[A, B any] struct {
	V1 A
	V2 B
}

func NewEntry[A, B any](v1 A, v2 B) *Entry[A, B] {
	return &Entry[A, B]{
		V1: v1,
		V2: v2,
	}
}

func Each[T map[K]V, K comparable, V any](expr validfuncs.Expr[*Entry[K, V]], left ...validfuncs.Expr[*Entry[K, V]]) validfuncs.Expr[T] {
	return AEach[T, K, V](expr, left...)
}

func AEach[T ~map[K]V, K comparable, V any](expr validfuncs.Expr[*Entry[K, V]], left ...validfuncs.Expr[*Entry[K, V]]) validfuncs.Expr[T] {
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
					for k, v := range v {
						if !validfuncs.WithSelector(fmt.Sprintf("[%v]", k), instrs).Exec(ctx, NewEntry(k, v)) {
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

func EachAnd[T ~map[K]V, K comparable, V any](exprs *utils.VarArgs[utils.N1, validfuncs.Expr[*Entry[K, V]]], next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(AEach[T](exprs.FixedArgs()[0], exprs.LeftArgs()...), next)
}

func Value[T map[K]V, K comparable, V any](key K, expr validfuncs.Expr[V], left ...validfuncs.Expr[V]) validfuncs.Expr[T] {
	return AValue[T, K, V](key, expr, left...)
}

func AValue[T ~map[K]V, K comparable, V any](key K, expr validfuncs.Expr[V], left ...validfuncs.Expr[V]) validfuncs.Expr[T] {
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
				func(ctx *validfuncs.Context, m T) bool {
					v, ok := m[key]
					if !ok {
						return false
					}
					return validfuncs.WithSelector(fmt.Sprintf("[\"%v\"]", key), instrs).Exec(ctx, v)
				},
				func() string {
					return fmt.Sprintf("key %v", key)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("One", expr)
		},
	)
}

func ValueAnd[T ~map[K]V, K comparable, V any](key K, exprs *utils.VarArgs[utils.N1, validfuncs.Expr[V]], next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(AValue[T](key, exprs.FixedArgs()[0], exprs.LeftArgs()...), next)
}
