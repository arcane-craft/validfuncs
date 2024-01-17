package slices

import (
	"fmt"

	"github.com/arcane-craft/validfuncs"
	"github.com/arcane-craft/validfuncs/utils"
)

func Len[T []E, E any](size int) validfuncs.Expr[T] {
	return ALen[T, E](size)
}

func ALen[T ~[]E, E any](size int) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					return len(v) == size
				},
				func() string {
					return fmt.Sprintf("length %v", size)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Len", size)
		},
	)
}

func LenAnd[T ~[]E, E any](size int, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(ALen[T](size), next)
}

func RangeLen[T []E, E any](min int, max int) validfuncs.Expr[T] {
	return ARangeLen[T, E](min, max)
}

func ARangeLen[T ~[]E, E any](min int, max int) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr[T](
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

func RangeLenAnd[T ~[]E, E any](min int, max int, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(ARangeLen[T](min, max), next)
}

func MinLen[T []E, E any](min int) validfuncs.Expr[T] {
	return AMinLen[T, E](min)
}

func AMinLen[T ~[]E, E any](min int) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					return len(v) >= min
				},
				func() string {
					return fmt.Sprintf("min length %v", min)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("MinLen", min)
		},
	)
}

func MinLenAnd[T ~[]E, E any](min int, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(AMinLen[T](min), next)
}

func MaxLen[T []E, E any](max int) validfuncs.Expr[T] {
	return AMaxLen[T, E](max)
}

func AMaxLen[T ~[]E, E any](max int) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					return len(v) <= max
				},
				func() string {
					return fmt.Sprintf("max length %v", max)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("MaxLen", max)
		},
	)
}

func MaxLenAnd[T ~[]E, E any](max int, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(AMaxLen[T](max), next)
}
