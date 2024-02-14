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

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		length: len(input),
		tokens: make([]Token, 0),
		pos:    0,
	}
}

func (l *Lexer) Tokenize() []Token {
	for l.pos < l.length {
		current := l.peek(0)
		if l.isDigit(current) {
			l.tokenizeNumber()
		} else if isLetter(current) {
			l.tokenizeWord()
		} else if current == '#' {
			l.next()
			l.tokenizeHexNumber()
		} else if strings.ContainsRune("+-*/()", current) {
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

func (l *Lexer) tokenizeHexNumber() {
	var buffer strings.Builder
	current := l.peek(0)
	for l.isDigit(current) {
		buffer.WriteRune(current)
		current = l.next()
	}
	l.addToken(HEX_NUMBER, buffer.String())
}

func (l *Lexer) isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isLetter(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func (l *Lexer) tokenizeOperator() {
	operators := "+-*/()"
	position := strings.IndexRune(operators, l.peek(0))
	tokenTypes := []TokenType{PLUS, MINUS, STAR, SLASH, LPAREN, RPAREN}
	l.addToken(tokenTypes[position], "")
	l.next()
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
	l.addToken(WORD, buffer.String())
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
