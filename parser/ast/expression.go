package ast

type Expression interface {
	Evel() float32
	ToString() string
}
