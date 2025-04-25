package parser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type syntaxError struct {
	line, col int
	msg       string
}

func (e *syntaxError) Error() string {
	return fmt.Sprintf("Syntax error at line %d:%d - %s", e.line, e.col, e.msg)
}

type runeErrorListener struct {
	*antlr.DefaultErrorListener
	err error
}

func newRuneErrorListener() *runeErrorListener {
	return &runeErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}

func (l *runeErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol any,
	line, col int,
	msg string,
	ex antlr.RecognitionException,
) {
	panic(&syntaxError{line, col, msg})
}

func Parse(input *antlr.InputStream) (tree IProgramContext, err error) {
	defer func() {
		if r := recover(); r != nil {
			if rErr, ok := r.(*syntaxError); ok {
				tree, err = nil, rErr
			} else {
				panic(r)
			}
		}
	}()

	lexer := NewRuneLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(newRuneErrorListener())

	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	parser := NewRuneParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(newRuneErrorListener())

	return parser.Program(), nil
}

func ParseString(input string) (antlr.ParseTree, error) {
	return Parse(antlr.NewInputStream(input))
}
