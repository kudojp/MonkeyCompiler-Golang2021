package ast

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode() // just for improving productivity
}

type Expression interface {
	Node
	expressionNode() // just for improving productivity
}
