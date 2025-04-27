package ast

import "github.com/antlr4-go/antlr/v4"

type DeclarationStatement struct {
	Token   antlr.Token
	Type    string
	Name    *IdentifierExpression
}

func (ds *DeclarationStatement) statementNode() {}
func (ds *DeclarationStatement) TokenLiteral() string { return ds.Token.GetText() }

type AssignmentStatement struct {
	Token antlr.Token
	Name  *IdentifierExpression
	Value Expression
}

func (as *AssignmentStatement) statementNode() {}
func (as *AssignmentStatement) TokenLiteral() string { return as.Token.GetText() }

type DeclarationAssignmentStatement struct {
	Token   antlr.Token
	Type    string
	Name    *IdentifierExpression
	Value   Expression
}

func (das *DeclarationAssignmentStatement) statementNode() {}
func (das *DeclarationAssignmentStatement) TokenLiteral() string { return das.Token.GetText() }

