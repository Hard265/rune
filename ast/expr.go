package ast

import "github.com/antlr4-go/antlr/v4"

type IdentifierExpression struct {
	Token antlr.Token
	Name string
}

func (i *IdentifierExpression) expressionNode()      {}
func (i *IdentifierExpression) TokenLiteral() string { return i.Token.GetText() }

type IntegerLiteral struct {
	Token antlr.Token
	Value string
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.GetText() }

type StringLiteral struct {
	Token antlr.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.GetText() }

type BooleanLiteral struct {
	Token antlr.Token
	Value string
}

func (sl *BooleanLiteral) expressionNode()      {}
func (sl *BooleanLiteral) TokenLiteral() string { return sl.Token.GetText() }

type DoubleLiteral struct {
	Token antlr.Token
	Value string
}

func (sl *DoubleLiteral) expressionNode()      {}
func (sl *DoubleLiteral) TokenLiteral() string { return sl.Token.GetText() }

type UnaryExpression struct {
	Operator string
	Token    antlr.Token
	Right    Expression
}

func (ue *UnaryExpression) expressionNode()      {}
func (ue *UnaryExpression) TokenLiteral() string { return ue.Token.GetText() }

type LogicalAndExpression struct {
	Token antlr.Token
	Left  Expression
	Right Expression
}

func (lae *LogicalAndExpression) expressionNode()      {}
func (lae *LogicalAndExpression) TokenLiteral() string { return lae.Token.GetText() }

type LogicalOrExpression struct {
	Token antlr.Token
	Left  Expression
	Right Expression
}

func (loe *LogicalOrExpression) expressionNode()      {}
func (loe *LogicalOrExpression) TokenLiteral() string { return loe.Token.GetText() }

type EqualityExpression struct {
	Token antlr.Token
	Left  Expression
	Right Expression
}

func (ee *EqualityExpression) expressionNode()      {}
func (ee *EqualityExpression) TokenLiteral() string { return ee.Token.GetText() }

type RelationalExpression struct {
	Left  Expression
	Token antlr.Token
	Right Expression
}

func (re *RelationalExpression) expressionNode()      {}
func (re *RelationalExpression) TokenLiteral() string { return re.Token.GetText() }

type AdditiveExpression struct {
	Left  Expression
	Token antlr.Token
	Right Expression
}

func (ae *AdditiveExpression) expressionNode()      {}
func (ae *AdditiveExpression) TokenLiteral() string { return ae.Token.GetText() }

type MultiplicativeExpression struct {
	Left  Expression
	Token antlr.Token
	Right Expression
}

func (me *MultiplicativeExpression) expressionNode()      {}
func (me *MultiplicativeExpression) TokenLiteral() string { return me.Token.GetText() }
