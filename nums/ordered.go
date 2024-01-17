package nums

import (
	"fmt"

	"github.com/arcane-craft/validfuncs"
	"github.com/arcane-craft/validfuncs/utils"
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](min T) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					return v >= min
				},
				func() string {
					return fmt.Sprintf("min %v", min)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Min", min)
		},
	)
}

func MinAnd[T constraints.Ordered](min T, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(Min(min), next)
}

func Max[T constraints.Ordered](max T) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					return v <= max
				},
				func() string {
					return fmt.Sprintf("max %v", max)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("Max", max)
		},
	)
}

func MaxAnd[T constraints.Ordered](max T, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(Max(max), next)
}

func EQ[T constraints.Ordered](value T) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					return v == value
				},
				func() string {
					return fmt.Sprintf("equal %v", value)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("EQ", value)
		},
	)
}

func EQAnd[T constraints.Ordered](value T, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(EQ(value), next)
}

func LT[T constraints.Ordered](rhs T) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					return v < rhs
				},
				func() string {
					return fmt.Sprintf("less than %v", rhs)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("LT", rhs)
		},
	)
}

func LTAnd[T constraints.Ordered](rhs T, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(LT(rhs), next)
}

func GT[T constraints.Ordered](rhs T) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					return v > rhs
				},
				func() string {
					return fmt.Sprintf("greater than %v", rhs)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("GT", rhs)
		},
	)
}

func GTAnd[T constraints.Ordered](rhs T, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(GT(rhs), next)
}

func LTE[T constraints.Ordered](rhs T) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					return v <= rhs
				},
				func() string {
					return fmt.Sprintf("less than and equal %v", rhs)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("LTE", rhs)
		},
	)
}

func LTEAnd[T constraints.Ordered](rhs T, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(LTE(rhs), next)
}

func GTE[T constraints.Ordered](rhs T) validfuncs.Expr[T] {
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[T], error) {
			return validfuncs.NewInstr[T](
				func(ctx *validfuncs.Context, v T) bool {
					return v >= rhs
				},
				func() string {
					return fmt.Sprintf("greater than and equal %v", rhs)
				},
			), nil
		},
		func() string {
			return utils.StringifyExpr("GTE", rhs)
		},
	)
}

func GTEAnd[T constraints.Ordered](rhs T, next validfuncs.Expr[T]) validfuncs.Expr[T] {
	return validfuncs.NewCPSExpr(GTE(rhs), next)
}
