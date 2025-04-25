package tests

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hard265/rune/parser"
)

func TestParserBasicProgram(t *testing.T) {
	input := antlr.NewInputStream("int a\n")
	lexer := parser.NewRuneLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewRuneParser(tokenStream)

	tree := p.Program()
	
	if tree == nil {
		t.Errorf("Parsing produced nil tree.")
	}
}
