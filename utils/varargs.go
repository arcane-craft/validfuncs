package utils

type Natural interface {
	~struct{ asNatural struct{} }
	Eval() int
}

type N0 struct {
	asNatural struct{}
}

func (N0) Eval() int {
	return 0
}

type Succ[P Natural] struct{ asNatural struct{} }

func (Succ[P]) Eval() int {
	var pre P
	return pre.Eval() + 1
}

type N1 = Succ[N0]
type N2 = Succ[N1]
type N3 = Succ[N2]

type VarArgs[MIN Natural, T any] struct {
	args []T
}

func (va *VarArgs[MIN, T]) Eval() []T {
	var min MIN
	if len(va.args) < min.Eval() {
		panic("args not enough")
	}
	return va.args
}

func (va *VarArgs[MIN, T]) FixedArgs() []T {
	size := MIN{}.Eval()
	if len(va.args) < size {
		panic("args not enough")
	}
	return va.args[:size]
}

func (va *VarArgs[MIN, T]) LeftArgs() []T {
	size := MIN{}.Eval()
	if len(va.args) < size {
		panic("args not enough")
	}
	return va.args[size:]
}

func VA[T any](args ...T) *VarArgs[N0, T] {
	return &VarArgs[N0, T]{
		args: args,
	}
}

func VA1[T any](arg T, left ...T) *VarArgs[N1, T] {
	return &VarArgs[N1, T]{
		args: append([]T{arg}, left...),
	}
}

func VA2[T any](arg1, arg2 T, left ...T) *VarArgs[N2, T] {
	return &VarArgs[N2, T]{
		args: append([]T{arg1, arg2}, left...),
	}
}

func VA3[T any](arg1, arg2, arg3 T, left ...T) *VarArgs[N3, T] {
	return &VarArgs[N3, T]{
		args: append([]T{arg1, arg2, arg3}, left...),
	}
}
