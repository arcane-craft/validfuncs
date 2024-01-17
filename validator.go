package validfuncs

import (
	"fmt"
)

type Validation[T any] interface {
	Validate(T) error
}

type validationImpl[T any] struct {
	instrs Instruction[T]
}

func (imp *validationImpl[T]) Validate(value T) error {
	var ctx Context
	if !imp.instrs.Exec(&ctx, value) {
		return fmt.Errorf("%s", ctx.Failure())
	}
	return nil
}

func Compile[T any](expr Expr[T]) (Validation[T], error) {
	instrs, err := expr.Interpret()
	if err != nil {
		return nil, fmt.Errorf("interpret expression failed: %w", err)
	}
	return &validationImpl[T]{
		instrs: instrs,
	}, nil
}

func MustCompile[T any](expr Expr[T]) Validation[T] {
	validation, err := Compile(expr)
	if err != nil {
		panic(err.Error())
	}
	return validation
}
