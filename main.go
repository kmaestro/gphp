package main

import (
	"fmt"
	"php/parser"
)

func main() {
	lexer := parser.NewLexer("(2.6 + 2) * PI")
	token := lexer.Tokenize()
	parser := parser.NewParser(token)

	expressions := parser.Parse()

	for _, expression := range expressions {
		fmt.Printf("%s = %s \n", expression.String(), fmt.Sprint(expression.Eval()))
	}
}
