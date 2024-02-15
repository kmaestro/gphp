package ast

import (
	"fmt"
	valiables "php/lib/variables"
)

type variabletExpression struct {
	name string
}

func NewVariabletExpression(name string) *variabletExpression {
	return &variabletExpression{name: name}
}

func (ve *variabletExpression) Eval() float32 {
	if !valiables.IsExists(ve.name) {
		panic("Constant does not exists")
	}

	return valiables.Get(ve.name)
}

func (ve *variabletExpression) String() string {
	return fmt.Sprintf("%s", ve.name)
}
