package ast

import "fmt"

type echoStatement struct {
	expression Expression
}

func NewEchoStatement(expression Expression) *echoStatement {
	return &echoStatement{expression: expression}
}

func (es *echoStatement) Execute() {
	fmt.Println(es.expression.Eval())
}

func (es *echoStatement) String() string {
	fmt.Println(es.expression.Eval())
	return ""
}
