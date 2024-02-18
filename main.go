package main

import (
	"php/parser"
)

func main() {
	lexer := parser.NewLexer("$word = 2+2; $word2 = PI + $word; echo $word2")
	tokens := lexer.Tokenize()

	// for _, token := range tokens {
	// 	fmt.Printf("%s\n", token.String())
	// }
	parser := parser.NewParser(tokens)

	statements := parser.Parse()

	// for _, statement := range statements {
	// 	fmt.Printf("%s \n", statement.String())
	// }

	for _, statement := range statements {
		statement.Execute()
	}
	// fmt.Println(valiables.Get("word"))
	// fmt.Println(valiables.Get("word2"))
}
