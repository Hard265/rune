package tests

import (
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hard265/rune/ast"
	"github.com/hard265/rune/parser"
)

func TestBuild(t *testing.T) {
	var input = antlr.NewInputStream("int a = 5 + 6")
	var lexer = parser.NewRuneLexer(input)
	var tokensStream = antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewRuneParserParser(tokensStream)

	tree := p.Program();

	builder := &ast.ASTBuilder{}
	ast := builder.Visit(tree)
	if ast == nil{
		fmt.Println("ast is nil")
	}
}
