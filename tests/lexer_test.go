package tests

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/hard265/rune/parser"
)

func TestLexerTokens(t *testing.T) {
	input := antlr.NewInputStream("int x = 42\n")
	lexer := parser.NewRuneLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	tokenStream.Fill()

	expectedTokens := []string{
		"INT", "IDENTIFIER", "EQUAL", "INTEGER_LITERAL", "NEWLINE", "EOF",
	}
	tokens := tokenStream.GetAllTokens()

	for i, token := range tokens {
		if i >= len(expectedTokens) {
			t.Fatalf("Unexpected number of tokens")
			break
		}
		if token.GetTokenType() == antlr.TokenEOF {
			break
		}
		expected := expectedTokens[i]
		actual := lexer.SymbolicNames[token.GetTokenType()]

		if expected != actual {
			t.Fatalf("Expected token %s but got %s", expected, actual)
		}
	}

}
