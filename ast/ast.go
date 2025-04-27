package ast

type Node interface {
	TokenLiteral() string
}

type ProgramNode struct {
	Statements []Statement
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

