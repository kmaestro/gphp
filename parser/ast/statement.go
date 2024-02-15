package ast

type Statement interface {
	Execute()
	String() string
}
