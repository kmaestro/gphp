package main

import (
	"fmt"
	"log"
	"os"
	"php/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Error getting file information: %v", err)
	}
	fileSize := fileInfo.Size()

	content := make([]byte, fileSize)

	_, err = file.Read(content)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	lexer := parser.NewLexer(string(content))
	tokens := lexer.Tokenize()

	for _, token := range tokens {
		fmt.Printf("%s\n", token.String())
	}
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
