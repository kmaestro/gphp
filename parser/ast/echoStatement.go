package ast

import "fmt"

type echoStatement struct {
	expression Expression
}

func NewEchoStatement(expression Expression) *echoStatement {
	return &echoStatement{expression: expression}
}

func (es *echoStatement) Execute() {
	fmt.Print(es.expression.Eval())
}
