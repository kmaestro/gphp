package ast

type unaryExpression struct {
	operation string
	expr1     Expression
}

func NewUnaryExpression(operation string, expr1 Expression) *unaryExpression {
	return &unaryExpression{operation: operation, expr1: expr1}
}

func (ue *unaryExpression) Evel() float32 {
	switch ue.operation {
	case "-":
		return -ue.expr1.Evel()
	default:
		return ue.expr1.Evel()
	}
}
