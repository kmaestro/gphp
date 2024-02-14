package ast

import (
	"fmt"
)

type binaryExpression struct {
	expr1, expr2 Expression
	operation    rune
}

func NewBinaryExpression(operation rune, expr1 Expression, expr2 Expression) *binaryExpression {
	return &binaryExpression{
		operation: operation,
		expr1:     expr1,
		expr2:     expr2,
	}
}

func (be *binaryExpression) Evel() float32 {
	switch be.operation {
	case '-':
		return be.expr1.Evel() - be.expr2.Evel()
	case '*':
		return be.expr1.Evel() * be.expr2.Evel()
	case '/':
		return be.expr1.Evel() / be.expr2.Evel()
	default:
		return be.expr1.Evel() + be.expr2.Evel()
	}
}

func (be *binaryExpression) ToString() string {
	return fmt.Sprintf("[%s %c %s]", be.expr1.ToString(), be.operation, be.expr2.ToString())
}
