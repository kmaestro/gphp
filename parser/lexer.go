package parser

import (
	"strconv"
	"strings"
)

const OPRATOR_CHARS = "+-*/()"

var OPERATOR_TOKENS = map[string]int8{
	"+": int8(PLUS),
	"-": int8(MINUS),
	"*": int8(STAR),
	"/": int8(SLASH),
	"(": int8(LPAREN),
	")": int8(RPAREN),
}

type Lexer struct {
	input  string
	length int64
	tokens []Token
	pos    int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, length: int64(len(input))}
}

func (l *Lexer) Tokenize() []Token {
	for l.pos < int(l.length) {
		currency := l.peek(0)
		_, error := strconv.ParseInt(currency, 8, 8)
		if error == nil {
			l.tokenizeNumber()
		} else if strings.Contains(OPRATOR_CHARS, currency) {
			l.tokenizeOperator()
		} else {
			l.next()
		}
	}
	return l.tokens
}

func (l *Lexer) tokenizeNumber() {
	currency := l.peek(0)
	l.tokens = append(l.tokens, *NewToken(NUMBER, currency))
	l.next()
}

func (l *Lexer) tokenizeOperator() {
	currency := l.peek(0)
	l.tokens = append(l.tokens, *NewToken(TokenType(OPERATOR_TOKENS[currency]), currency))
	l.next()
}

func (l *Lexer) peek(relativePosition int) string {
	position := l.pos + relativePosition
	if position >= int(l.length) {
		return string(ILLEGAL)
	}
	return string(l.input[position])
}

func (l *Lexer) next() string {
	l.pos++
	return l.peek(0)
}
