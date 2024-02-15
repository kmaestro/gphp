package ast

import (
	"fmt"
	valiables "php/lib/variables"
)

type AssignmentStatement struct {
	variable   string
	expression Expression
}

func NewAssignmentStatement(variable string, expression Expression) *AssignmentStatement {
	return &AssignmentStatement{
		variable:   variable,
		expression: expression,
	}
}

func (a *AssignmentStatement) Execute() {
	result := a.expression.Eval()
	valiables.Set(a.variable, result)
}

func (a *AssignmentStatement) String() string {
	return fmt.Sprintf("%s = %s", a.variable, a.expression)
}
