package generic

import "github.com/arcane-craft/validfuncs"

type DefRet[T any] func(validfuncs.Expr[T], ...validfuncs.Expr[T])

func Define[T any](decl func(v T, R DefRet[T])) validfuncs.Expr[T] {
	var zero T
	var ret validfuncs.Expr[T]
	decl(zero, func(expr validfuncs.Expr[T], left ...validfuncs.Expr[T]) {
		if len(left) > 0 {
			expr = All(expr, left[0], left[1:]...)
		}
		ret = expr
	})
	return ret
}
