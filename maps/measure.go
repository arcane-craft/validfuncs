package maps

import (
	"fmt"

	"github.com/arcane-craft/validfuncs"
	"github.com/arcane-craft/validfuncs/utils"
)

func Len[T map[K]V, K comparable, V any](size int) validfuncs.Expr[T] {
	return ALen[T, K, V](size)
}

func ALen[T ~map[K]V, K comparable, V any](size int) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr(
				func(ctx *validfuncs.Context, v T) bool {
					return len(v) == size
				},
				func() string {
					return fmt.Sprintf("length %d", size)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Len", size)
		},
	)
}

func LenAnd[T ~map[K]V, K comparable, V any](size int, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(ALen[T](size), next)
}

func RangeLen[T map[K]V, K comparable, V any](min int, max int) validfuncs.Expr[T] {
	return ARangeLen[T, K, V](min, max)
}

func ARangeLen[T ~map[K]V, K comparable, V any](min int, max int) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr(
				func(ctx *validfuncs.Context, v T) bool {
					return len(v) >= min && len(v) <= max
				},
				func() string {
					return fmt.Sprintf("length between [%d, %d]", min, max)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("RangeLen", min, max)
		},
	)
}

func RangeLenAnd[T ~map[K]V, K comparable, V any](min int, max int, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(ARangeLen[T](min, max), next)
}

func MinLen[T map[K]V, K comparable, V any](min int) validfuncs.Expr[T] {
	return AMinLen[T, K, V](min)
}

func AMinLen[T ~map[K]V, K comparable, V any](min int) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr(
				func(ctx *validfuncs.Context, v T) bool {
					return len(v) >= min
				},
				func() string {
					return fmt.Sprintf("min length %d", min)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Min", min)
		},
	)
}

func MinLenAnd[T ~map[K]V, K comparable, V any](min int, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(AMinLen[T](min), next)
}

func MaxLen[T map[K]V, K comparable, V any](max int) validfuncs.Expr[T] {
	return AMaxLen[T, K, V](max)
}

func AMaxLen[T ~map[K]V, K comparable, V any](max int) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr(
				func(ctx *validfuncs.Context, v T) bool {
					return len(v) <= max
				},
				func() string {
					return fmt.Sprintf("max length %d", max)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Max", max)
		},
	)
}

func MaxLenAnd[T ~map[K]V, K comparable, V any](max int, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(AMaxLen[T](max), next)
}
