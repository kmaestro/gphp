package parser

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

type Token struct {
	tokenType int8
	text      string
}

func NewToken(tokenType TokenType, text string) *Token {
	return &Token{
		text:      text,
		tokenType: int8(tokenType),
	}
}
