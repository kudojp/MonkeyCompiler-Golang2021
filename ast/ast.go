package ast

import (
	"monkey/token"
)

type Node interface {
	TokenLiteral() string // a corresponding literal in a source code
}

type Statement interface {
	Node
	statementNode() // just for improving productivity
}

type Expression interface {
	Node
	expressionNode() // just for improving productivity
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // = token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {} // TODO??
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
type Identifier struct {
	Token token.Token //=  token.IDENT
	Value string
}

// In this abstraction tree, we take Identifier as an expression node. This is for simplicity.
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
