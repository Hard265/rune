package ast

import "github.com/antlr4-go/antlr/v4"

type DeclarationStatement struct {
	Token   antlr.Token
	Type    string
	Name    *Identifier
}

func (ds *DeclarationStatement) statementNode() {}
func (ds *DeclarationStatement) TokenLiteral() string { return ds.Token.GetText() }

type AssignmentStatement struct {
	Token antlr.Token
	Name  *Identifier
	Value Expression
}

func (as *AssignmentStatement) statementNode() {}
func (as *AssignmentStatement) TokenLiteral() string { return as.Token.GetText() }

type DeclarationAssignmentStatement struct {
	Token   antlr.Token
	Type    string
	Name    *Identifier
	Value   Expression
}

func (das *DeclarationAssignmentStatement) statementNode() {}
func (das *DeclarationAssignmentStatement) TokenLiteral() string { return das.Token.GetText() }

