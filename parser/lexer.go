package parser

import (
	"strings"
)

type Lexer struct {
	input  string
	length int
	tokens []Token
	pos    int
}

var operators = map[rune]TokenType{
	'+': PLUS,
	'-': MINUS,
	'*': STAR,
	'/': SLASH,
	'(': LPAREN,
	')': RPAREN,
	'=': EQ,
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		length: len(input),
		tokens: make([]Token, 0),
		pos:    0,
	}
}

func (l *Lexer) Tokenize() []Token {
	isCode := false
	var buffer strings.Builder
	for l.pos < l.length {
		current := l.peek(0)
		if buffer.String() != "<?php" && !isCode {
			buffer.WriteRune(current)
			l.next()
			continue
		} else {
			isCode = true
		}
		if l.isDigit(current) {
			l.tokenizeNumber()
		} else if current == '$' {
			l.tokenizeVar()
		} else if isLetter(current) {
			l.tokenizeWord()
		} else if _, operator := operators[current]; operator {
			l.tokenizeOperator()
		} else {
			// whitespaces
			l.next()
		}
	}
	return l.tokens
}

func (l *Lexer) tokenizeNumber() {
	var buffer strings.Builder
	current := l.peek(0)
	for {
		if current == '.' {
			if strings.Contains(buffer.String(), ".") {
				panic("Invalid float number")
			}
		} else if !l.isDigit(current) {
			break
		}
		buffer.WriteRune(current)
		current = l.next()
	}
	l.addToken(NUMBER, buffer.String())
}

func (l *Lexer) isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isLetter(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func (l *Lexer) tokenizeOperator() {
	operators := "+-*/()="
	position := strings.IndexRune(operators, l.peek(0))
	tokenTypes := []TokenType{PLUS, MINUS, STAR, SLASH, LPAREN, RPAREN, EQ}
	l.addToken(tokenTypes[position], "")
	l.next()
}

func (l *Lexer) tokenizeVar() {
	var buffer strings.Builder
	l.next()
	current := l.peek(0)
	for {
		if !isLetter(current) && !l.isDigit(current) && current != '_' && current != '$' {
			break
		}
		buffer.WriteRune(current)
		current = l.next()
	}
	l.addToken(VARIABLE, buffer.String())
}

func (l *Lexer) tokenizeWord() {
	var buffer strings.Builder
	current := l.peek(0)
	for {
		if !isLetter(current) && !l.isDigit(current) && current != '_' && current != '$' {
			break
		}
		buffer.WriteRune(current)
		current = l.next()
	}
	switch buffer.String() {
	case "if":
		l.addToken(IF, "")
	case "echo":
		l.addToken(ECHO, "")
	default:
		l.addToken(CONSTANT, buffer.String())
	}
}

func (l *Lexer) next() rune {
	l.pos++
	return l.peek(0)
}

func (l *Lexer) peek(relativePosition int) rune {
	position := l.pos + relativePosition
	if position >= l.length {
		return '\x00'
	}
	return rune(l.input[position])
}

func (l *Lexer) addToken(tokenType TokenType, text string) {
	l.tokens = append(l.tokens, *NewToken(tokenType, text))
}
