package parser

import (
	"errors"
	"fmt"
	"php/parser/ast"
	"strconv"
)

var eof Token = Token{text: "", tokenType: int8(EOF)}

type Parser struct {
	tokens []Token
	pos    int
	size   int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, size: len(tokens)}
}

func (p *Parser) Parse() []ast.Expression {
	var result []ast.Expression
	for !p.match(EOF) {
		result = append(result, p.expression())
	}
	fmt.Println(result)
	return result
}

func (p *Parser) expression() ast.Expression {
	return p.additive()
}

func (p *Parser) additive() ast.Expression {
	var result ast.Expression = p.multiplicative()

	for true {
		if p.match(PLUS) {
			result = ast.NewBinaryExpression("+", result, p.multiplicative())
			continue
		}
		if p.match(MINUS) {
			result = ast.NewBinaryExpression("-", result, p.multiplicative())
			continue
		}
		break
	}
	return result
}

func (p *Parser) multiplicative() ast.Expression {
	var result ast.Expression = p.unary()

	for true {
		if p.match(STAR) {
			result = ast.NewBinaryExpression("*", result, p.unary())
			continue
		}
		if p.match(SLASH) {
			result = ast.NewBinaryExpression("/", result, p.unary())
			continue
		}
		break
	}
	return result
}

func (p *Parser) unary() ast.Expression {

	if p.match(MINUS) {
		expression, _ := p.primary()
		return ast.NewUnaryExpression("*", expression)
	}
	if p.match(PLUS) {
		expression, _ := p.primary()
		return expression
	}

	expression, _ := p.primary()
	return expression
}

func (p *Parser) primary() (ast.Expression, error) {
	current := p.get(0)

	if p.match(NUMBER) {
		s, _ := strconv.ParseFloat(current.text, 32)
		return ast.NewNumberExpression(float32(s)), nil
	}

	if p.match(HEX_NUMBER) {
		s, _ := strconv.ParseFloat(current.text, 32)
		return ast.NewNumberExpression(float32(s)), nil
	}

	if p.match(LPAREN) {
		var result ast.Expression = p.expression()
		p.match(RPAREN)
		return result, nil
	}

	return nil, errors.New("Unknown expression")
}

func (p *Parser) match(tokenType TokenType) bool {
	currency := p.get(0)

	if tokenType != TokenType(currency.tokenType) {
		return false
	}
	p.pos++
	return true
}

func (p *Parser) get(relativePosition int) Token {
	position := p.pos + relativePosition
	if position >= int(p.size) {
		return eof
	}
	return p.tokens[p.pos]
}
