package parser

import "fmt"

type TokenType rune

const (
	ILLEGAL TokenType = iota
	NUMBER
	HEX_NUMBER

	PLUS
	MINUS
	STAR
	SLASH

	LPAREN
	RPAREN

	EQ

	WORD

	EOF
)

func tokenTypeToString(t TokenType) string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case NUMBER:
		return "NUMBER"
	case HEX_NUMBER:
		return "HEX_NUMBER"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case STAR:
		return "STAR"
	case SLASH:
		return "SLASH"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case EQ:
		return "EQ"
	case WORD:
		return "WORD"
	case EOF:
		return "EOF"
	default:
		return "Unknown"
	}
}

type Token struct {
	tokenType TokenType
	text      string
}

func NewToken(tokenType TokenType, text string) *Token {
	return &Token{
		text:      text,
		tokenType: tokenType,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s", tokenTypeToString(TokenType(t.tokenType)), t.text)
}
