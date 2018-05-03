package eval

type Expr interface {
	Eval(env Env) float64          //在environment变量中计算表达式的值
	Check(vars map[Var]bool) error //检测表达式操作符以及操作数的有效性，并将表达式中的变量聚集为一个slice
}

type Var string

type literal float64

type unary struct {
	op rune
	x  Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}
