package validfuncs

import (
	"fmt"
	"strings"
)

type Failure struct {
	Selector    string
	Constraints string
}

func (f Failure) String() string {
	return fmt.Sprintf("%s not satisfy constraints: %s", f.Selector, f.Constraints)
}

type Context struct {
	selectorNames []string
	constraints   string
}

func (c *Context) Failure() Failure {
	return Failure{
		Selector:    strings.Join(c.selectorNames, "."),
		Constraints: c.constraints,
	}
}

type Instruction[T any] interface {
	Exec(*Context, T) bool
	fmt.Stringer
}

type Instructions[T any] []Instruction[T]

func (is Instructions[T]) Exec(ctx *Context, v T) bool {
	for _, instr := range is {
		if !instr.Exec(ctx, v) {
			return false
		}
	}
	return true
}

func (is Instructions[T]) String() string {
	var ret string
	for i, instr := range is {
		ret += instr.String()
		if i < len(is)-1 {
			ret += ", "
		}
	}
	return ret
}

type simpleInstr[T any] struct {
	exec     func(*Context, T) bool
	stringer func() string
}

func (i simpleInstr[T]) Exec(ctx *Context, v T) bool {
	return i.exec(ctx, v)
}

func (i simpleInstr[T]) String() string {
	return i.stringer()
}

func NewInstr[T any](
	exec func(ctx *Context, v T) bool,
	stringer func() string,
) Instruction[T] {
	return simpleInstr[T]{exec, stringer}
}

func WithSelector[T any](name string, instr Instruction[T]) Instruction[T] {
	return NewInstr(
		func(ctx *Context, v T) bool {
			ctx.selectorNames = append(ctx.selectorNames, name)
			if !instr.Exec(ctx, v) {
				if len(ctx.constraints) <= 0 {
					ctx.constraints = instr.String()
				}
				return false
			}
			ctx.selectorNames = ctx.selectorNames[:len(ctx.selectorNames)-1]
			return true
		},
		func() string {
			return ""
		},
	)
}
