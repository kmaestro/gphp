package main

import (
	"fmt"
	"php/parser"
)

func main() {
	lexer := parser.NewLexer("(2.6 + 2) * 4")
	token := lexer.Tokenize()
	parser := parser.NewParser(token)

	expressions := parser.Parse()

	for _, expression := range expressions {
		fmt.Printf("%s = %s \n", expression.ToString(), fmt.Sprint(expression.Evel()))
	}
}
