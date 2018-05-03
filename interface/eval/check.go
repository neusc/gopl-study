package eval

import (
	"strings"
	"fmt"
)

// 每种具体类型进行操作符和操作数是否合法的检测

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Check(vars map[Var]bool) error {
	// 检测一元操作符的有效性
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected unary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok { // 检测函数名称的有效性
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity { // 检测函数传参个数是否有效
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

// 函数调用中每种函数对应的参数个数
var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}
