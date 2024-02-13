package main

import (
	"fmt"
	"php/parser"
)

func main() {
	lexer := parser.NewLexer("1 + 3 * 4")
	token := lexer.Tokenize()
	parser := parser.NewParser(token)

	expressions := parser.Parse()

	for _, expression := range expressions {
		fmt.Println(expression.Evel())
	}
}
