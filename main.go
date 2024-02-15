package main

import (
	"fmt"
	"php/parser"
)

func main() {
	lexer := parser.NewLexer("(2.6 + 2) * PI")
	tokens := lexer.Tokenize()

	for _, token := range tokens {
		fmt.Printf("%s\n", token.String())
	}
	parser := parser.NewParser(tokens)

	expressions := parser.Parse()

	for _, expression := range expressions {
		fmt.Printf("%s = %s \n", expression.String(), fmt.Sprint(expression.Eval()))
	}
}
