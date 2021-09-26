package parser

import (
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	// At this point, cursor's curToken is the first token in a vien
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken()
	p.peekToken = p.l.NextToken() // move the lexer's position
}
