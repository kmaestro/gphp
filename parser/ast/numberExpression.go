package ast

import (
	"fmt"
)

type NumberExpression struct {
	value float32
}

func NewNumberExpression(value float32) *NumberExpression {
	return &NumberExpression{value: value}
}

func (ne *NumberExpression) Eval() float32 {
	return ne.value
}

func (ne *NumberExpression) String() string {
	return fmt.Sprint(ne.value)
}
