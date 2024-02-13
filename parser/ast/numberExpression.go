package ast

type NumberExpression struct {
	value float32
}

func NewNumberExpression(value float32) *NumberExpression {
	return &NumberExpression{value: value}
}

func (ne *NumberExpression) Evel() float32 {
	return ne.value
}
