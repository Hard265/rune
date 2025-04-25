package ast

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/hard265/rune/parser"
)

type ASTBuilder struct {
	*parser.BaseRuneParserVisitor
}

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{}
}

func (v *ASTBuilder) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *ASTBuilder) VisitDeclaration(ctx *parser.DeclarationContext) any {
	typ := ctx.Type_().GetText()
	name := ctx.IDENTIFIER().GetSymbol()

	return &DeclarationStatement{
		Token: name,
		Type:  typ,
		Name: &IdentifierExpression{
			Token: name,
			Name:  name.GetText(),
		},
	}
}

func (v *ASTBuilder) VisitAssignment(ctx *parser.AssignmentContext) any {
	name := ctx.IDENTIFIER().GetSymbol()
	expr := ctx.Expression().Accept(v).(Expression)

	return &AssignmentStatement{
		Token: ctx.EQUAL().GetSymbol(),
		Name: &IdentifierExpression{
			Token: name,
			Name:  name.GetText(),
		},
		Value: expr,
	}
}

func (v *ASTBuilder) VisitDeclarationAssignment(ctx *parser.DeclarationAssignmentContext) any {
	typ := ctx.Type_().GetText()
	name := ctx.IDENTIFIER().GetSymbol()
	expr := ctx.Expression().Accept(v).(Expression)

	return &DeclarationAssignmentStatement{
		Token: ctx.Type_().GetStart(),
		Type:  typ,
		Name: &IdentifierExpression{
			Token: name,
			Name: name.GetText(),
		},
		Value: expr,
	}
}

func (v *ASTBuilder) VisitExpression(ctx *parser.ExpressionContext) any {
	return ctx.UnaryExpression().Accept(v)
}

func (v *ASTBuilder) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) any {
	if ctx.BinaryExpression() != nil {
		return ctx.BinaryExpression().Accept(v)
	}

	right := ctx.UnaryExpression().Accept(v).(Expression)
	return &UnaryExpression{
		Token: ctx.GetStart(),
		Right: right,
	}
}

func (v *ASTBuilder) VisitBinaryExpression(ctx *parser.BinaryExpressionContext) any {
	return ctx.LogicalOrExpression().Accept(v)
}

func (v *ASTBuilder) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) any {
	left := ctx.LogicalAndExpression(0).Accept(v).(Expression)

	for i := 1; i < len(ctx.AllLogicalAndExpression()); i++ {
		tok := ctx.GetChild(2*i - 1).(antlr.TerminalNode).GetSymbol()

		right := ctx.LogicalAndExpression(i).Accept(v).(Expression)
		left = &LogicalOrExpression{
			Token: tok,
			Right: right,
			Left:  left,
		}
	}
	return left
}

func (v *ASTBuilder) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) any {
	left := ctx.EqualityExpression(0).Accept(v).(Expression)

	for i := 1; i < len(ctx.AllEqualityExpression()); i++ {
		tok := ctx.GetChild(2*i - 1).(antlr.TerminalNode).GetSymbol()

		right := ctx.EqualityExpression(i).Accept(v).(Expression)
		left = &LogicalAndExpression{
			Token: tok,
			Left:  left,
			Right: right,
		}
	}
	return left
}

func (v *ASTBuilder) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) any {
	left := ctx.RelationalExpression(0).Accept(v).(Expression)

	for i := 1; i < len(ctx.AllRelationalExpression()); i++ {
		tok := ctx.GetChild(2*i - 1).(antlr.TerminalNode).GetSymbol() // EQEQUAL or NOTEQUAL token
		right := ctx.RelationalExpression(i).Accept(v).(Expression)

		left = &EqualityExpression{
			Token: tok,
			Left:  left,
			Right: right,
		}
	}
	return left
}

func (v *ASTBuilder) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) any {
	left := ctx.AdditiveExpression(0).Accept(v).(Expression)

	for i := 1; i < len(ctx.AllAdditiveExpression()); i++ {
		token := ctx.GetChild(2*i - 1).(antlr.TerminalNode).GetSymbol()
		right := ctx.AdditiveExpression(i).Accept(v).(Expression)

		left = &RelationalExpression{
			Left:  left,
			Token: token,
			Right: right,
		}
	}

	return left
}

func (v *ASTBuilder) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) any {
	left := ctx.MultiplicativeExpression(0).Accept(v).(Expression)

	for i := 1; i < len(ctx.AllMultiplicativeExpression()); i++ {
		token := ctx.GetChild(2*i - 1).(antlr.TerminalNode).GetSymbol() // PLUS or MINUS token
		right := ctx.MultiplicativeExpression(i).Accept(v).(Expression)

		left = &AdditiveExpression{
			Left:  left,
			Token: token,
			Right: right,
		}
	}

	return left
}

func (v *ASTBuilder) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) any {
	left := ctx.PrimaryExpression(0).Accept(v).(Expression)

	for i := 1; i < len(ctx.AllPrimaryExpression()); i++ {
		token := ctx.GetChild(2*i - 1).(antlr.TerminalNode).GetSymbol() // STAR or SLASH token
		right := ctx.PrimaryExpression(i).Accept(v).(Expression)

		left = &MultiplicativeExpression{
			Left:  left,
			Token: token,
			Right: right,
		}
	}
	return left
}

func (v *ASTBuilder) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) any {
	switch {
	case ctx.Literal() != nil:
		return ctx.Literal().Accept(v)
	case ctx.IDENTIFIER() != nil:
		return &IdentifierExpression{
			Name: ctx.IDENTIFIER().GetText(),
		}
	case ctx.Expression() != nil:
		return ctx.Expression().Accept(v)
	default:
		panic("unknown primaryExpression type")
	}
}

func (v *ASTBuilder) VisitLiteral(ctx *parser.LiteralContext) any {
	switch {
	case ctx.INTEGER_LITERAL() != nil:
		return &IntegerLiteral{
			Value: ctx.INTEGER_LITERAL().GetText(),
		}
	case ctx.BOOLEAN_LITERAL() != nil:
		return &BooleanLiteral{
			Value: ctx.BOOLEAN_LITERAL().GetText(),
		}
	case ctx.STRING_LITERAL() != nil:
		return &StringLiteral{
			Value: ctx.STRING_LITERAL().GetText(),
		}
	case ctx.DOUBLE_LITERAL() != nil:
		return &DoubleLiteral{
			Value: ctx.DOUBLE_LITERAL().GetText(),
		}
	default:
		panic("unknown literal type")
	}
}
