package ast

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
