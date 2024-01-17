package strings

import (
	"fmt"

	"github.com/arcane-craft/validfuncs"
	"github.com/arcane-craft/validfuncs/utils"
)

func Len(size int) validfuncs.Expr[string] {
	return ALen[string](size)
}

func ALen[T ~string](size int) validfuncs.Expr[T] {
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

func LenAnd[T ~string](size int, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(ALen[T](size), next)
}

func RangeLen(min int, max int) validfuncs.Expr[string] {
	return ARangeLen[string](min, max)
}

func ARangeLen[T ~string](min int, max int) validfuncs.Expr[T] {
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

func RangeLenAnd[T ~string](min int, max int, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(ARangeLen[T](min, max), next)
}

func MinLen(min int) validfuncs.Expr[string] {
	return AMinLen[string](min)
}

func AMinLen[T ~string](min int) validfuncs.Expr[T] {
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

func MinLenAnd[T ~string](min int, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(AMinLen[T](min), next)
}

func MaxLen(max int) validfuncs.Expr[string] {
	return AMaxLen[string](max)
}

func AMaxLen[T ~string](max int) validfuncs.Expr[T] {
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

func MaxLenAnd[T ~string](max int, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(AMaxLen[T](max), next)
}
