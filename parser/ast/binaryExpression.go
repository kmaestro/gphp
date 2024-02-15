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

func (be *binaryExpression) Eval() float32 {
	switch be.operation {
	case '-':
		return be.expr1.Eval() - be.expr2.Eval()
	case '*':
		return be.expr1.Eval() * be.expr2.Eval()
	case '/':
		return be.expr1.Eval() / be.expr2.Eval()
	default:
		return be.expr1.Eval() + be.expr2.Eval()
	}
}

func (be *binaryExpression) String() string {
	return fmt.Sprintf("[%s %c %s]", be.expr1.String(), be.operation, be.expr2.String())
}
