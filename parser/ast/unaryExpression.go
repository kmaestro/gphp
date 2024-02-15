package ast

import "fmt"

type unaryExpression struct {
	operation rune
	expr1     Expression
}

func NewUnaryExpression(operation rune, expr1 Expression) *unaryExpression {
	return &unaryExpression{operation: operation, expr1: expr1}
}

func (ue *unaryExpression) Eval() float32 {
	switch ue.operation {
	case '-':
		return -ue.expr1.Eval()
	default:
		return ue.expr1.Eval()
	}
}

func (ue *unaryExpression) String() string {
	return fmt.Sprint("%c %s", ue.operation, ue.expr1.String())
}
