// Code generated from grammar/RuneParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // RuneParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by RuneParserParser.
type RuneParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RuneParserParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by RuneParserParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by RuneParserParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by RuneParserParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by RuneParserParser#declarationAssignment.
	VisitDeclarationAssignment(ctx *DeclarationAssignmentContext) interface{}

	// Visit a parse tree produced by RuneParserParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by RuneParserParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by RuneParserParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by RuneParserParser#binaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by RuneParserParser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by RuneParserParser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by RuneParserParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by RuneParserParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by RuneParserParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by RuneParserParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by RuneParserParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by RuneParserParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
