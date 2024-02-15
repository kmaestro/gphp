package ast

type Expression interface {
	Eval() float32
	String() string
}
