package parser

import (
	"fmt"
	"php/parser/ast"
	"strconv"
)

type Parser struct {
	tokens []Token
	size   int
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
		size:   len(tokens),
		pos:    0,
	}
}

func (p *Parser) Parse() []ast.Statement {
	result := make([]ast.Statement, 0)
	for !p.match(EOF) {
		result = append(result, p.statement())
	}
	return result
}

func (p *Parser) statement() ast.Statement {
	return p.assignmentStatement()
}

func (p *Parser) assignmentStatement() ast.Statement {
	current := p.get(0)
	if p.match(ECHO) {
		return ast.NewEchoStatement(p.expression())
	}
	if (p.match(VARIABLE) || p.match(CONSTANT)) && p.get(0).tokenType == EQ {
		variable := current.text
		p.consume(EQ)
		return ast.NewAssignmentStatement(variable, p.expression())
	}
	panic("Unknown statement")
}

func (p *Parser) consume(tokenType TokenType) Token {
	current := p.get(0)
	if tokenType != TokenType(current.tokenType) {
		panic(fmt.Sprintf("Token %s doesn't match %s", current.text, tokenTypeToString(tokenType)))
	}
	p.pos++
	return current
}
func (p *Parser) expression() ast.Expression {
	return p.additive()
}

func (p *Parser) additive() ast.Expression {
	result := p.multiplicative()

	for {
		if p.match(PLUS) {
			result = ast.NewBinaryExpression(
				'+',
				result,
				p.multiplicative(),
			)
			continue
		}
		if p.match(MINUS) {
			result = ast.NewBinaryExpression(
				'-',
				result,
				p.multiplicative(),
			)
			continue
		}
		break
	}

	return result
}

func (p *Parser) multiplicative() ast.Expression {
	result := p.unary()

	for {
		if p.match(STAR) {
			result = ast.NewBinaryExpression(
				'*',
				result,
				p.unary(),
			)
			continue
		}
		if p.match(SLASH) {
			result = ast.NewBinaryExpression(
				'/',
				result,
				p.unary(),
			)
			continue
		}
		break
	}

	return result
}

func (p *Parser) unary() ast.Expression {
	if p.match(MINUS) {
		return ast.NewUnaryExpression('-', p.primary())
	}
	if p.match(PLUS) {
		return p.primary()
	}
	return p.primary()
}

func (p *Parser) primary() ast.Expression {
	current := p.get(0)
	if p.match(NUMBER) {
		val, _ := strconv.ParseFloat(current.text, 64)
		return ast.NewNumberExpression(float32(val))
	}
	if p.match(HEX_NUMBER) {
		val, _ := strconv.ParseInt(current.text, 16, 64)
		return ast.NewNumberExpression(float32(val))
	}
	if p.match(VARIABLE) {
		return ast.NewVariabletExpression(current.text)
	}
	if p.match(CONSTANT) {
		return ast.NewVariabletExpression(current.text)
	}
	if p.match(LPAREN) {
		result := p.expression()
		p.match(RPAREN)
		return result
	}
	panic("Unknown expression")
}

func (p *Parser) match(tokenType TokenType) bool {
	current := p.get(0)
	if tokenType != TokenType(current.tokenType) {
		return false
	}
	p.pos++
	return true
}

func (p *Parser) get(relativePosition int) Token {
	position := p.pos + relativePosition
	if position >= p.size {
		return *NewToken(EOF, "")
	}
	return p.tokens[position]
}
