package validfuncs

import (
	"fmt"
	"strings"

	"github.com/arcane-craft/validfuncs/utils"
)

type Expr[T any] interface {
	Interpret() (Instruction[T], error)
	fmt.Stringer
}

type Exprs[T any] []Expr[T]

func (s Exprs[T]) Interpret() (Instructions[T], error) {
	var ret Instructions[T]
	for _, expr := range s {
		instrs, err := expr.Interpret()
		if err != nil {
			return nil, err
		}
		ret = append(ret, instrs)
	}
	return ret, nil
}

func (s Exprs[T]) String() string {
	var ret []string
	for _, expr := range s {
		ret = append(ret, expr.String())
	}
	return strings.Join(ret, " ")
}

type simpleExpr[T any] struct {
	interpret func() (Instruction[T], error)
	stringer  func() string
}

func (m simpleExpr[T]) Interpret() (Instruction[T], error) {
	return m.interpret()
}

func (m simpleExpr[T]) String() string {
	return m.stringer()
}

func NewExpr[T any](interpret func() (Instruction[T], error), stringer func() string) Expr[T] {
	return simpleExpr[T]{
		interpret: interpret,
		stringer:  stringer,
	}
}

type cpsExpr[T any] struct {
	current, next Expr[T]
}

func NewCPSExpr[T any](current, next Expr[T]) cpsExpr[T] {
	return cpsExpr[T]{
		current: current,
		next:    next,
	}
}

func (c cpsExpr[T]) Interpret() (Instruction[T], error) {
	currentInstr, err := c.current.Interpret()
	if err != nil {
		return nil, err
	}
	nextInstr, err := c.next.Interpret()
	if err != nil {
		return nil, err
	}
	return NewInstr[T](
		func(ctx *Context, v T) bool {
			if !currentInstr.Exec(ctx, v) {
				return false
			}
			return nextInstr.Exec(ctx, v)
		},
		func() string {
			return currentInstr.String() + ", " + nextInstr.String()
		},
	), nil
}

func (c cpsExpr[T]) String() string {
	return utils.StringifyExpr("CPS", c.current, c.next)
}
