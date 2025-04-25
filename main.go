package main

import (
	"fmt"
	"log" // Import log for potential errors

	// *** IMPORTANT: Adjust this import path to match your project structure ***
	// If RuneLexer.g4 is in your main package:
	// "./"
	// If RuneLexer.g4 is in a sub-package like "parser":
	// "./parser"
	// If using Go modules, it will be something like "your_module_path/your_package_containing_g4_files"
	// For example: "myproject/parser" if myproject is your module name and parser is the directory with the g4 file.
	"github.com/antlr4-go/antlr"
	runelexer "github.com/hard265/rune/parser" // <-- REPLACE THIS IMPORT PATH
)

func main() {
	// Test input with varying indentation
	inputString := `
int main():
    x = 10
    if x > 5:
        print(x)
    else:
        y = 20
        if y < 30:
            print(y)
    print("Done")
    another = true

final_value = 100.5

` // Added some more structure for better indentation test

	// Create input stream from the string
	input := antlr.NewInputStream(inputString)

	// Create the lexer instance using the generated constructor
	lexer := runelexer.NewRuneLexer(input)

	// *** IMPORTANT: Call the SetupLexer method you added in @lexer::members ***
	// This initializes the indentation stack and pending token list.
	fmt.Println("Lexer Setup Complete.")


	// Iterate through the tokens and print them
	fmt.Println("\nTokens:")
	for {
		token := lexer.NextToken() // Calls your overridden NextToken method

		// Stop when the End-Of-File token is reached
		if token.GetTokenType() == antlr.TokenEOF {
			// The overridden NextToken should handle final DEDENTs before EOF,
			// so we just print EOF and break.
			tokenTypeName := lexer.SymbolicNames[token.GetTokenType()]
			fmt.Printf("  Type: %s, Text: '%s'\n", tokenTypeName, token.GetText())
			break
		}

		// Get the token type name using the generated SymbolicNames array
		// This is safer than relying on RuneLexerTokens map for synthetic tokens like INDENT/DEDENT
		tokenTypeName := "UNKNOWN" // Default value
        if token.GetTokenType() >= 0 && token.GetTokenType() < len(lexer.SymbolicNames) {
            tokenTypeName = lexer.SymbolicNames[token.GetTokenType()]
        }


		// Print the token details (excluding hidden tokens like WS and comments)
		// Check if the channel is default channel (not hidden)
        // TokenDefaultChannel is 0, TokenHiddenChannel is 1
		if token.GetChannel() == antlr.TokenDefaultChannel {
			fmt.Printf("  Type: %s, Line: %d, Column: %d, Text: '%s'\n",
				tokenTypeName,
				token.GetLine(),
				token.GetColumn(),
				token.GetText(),
			)
		} else {
             // Optional: Print hidden tokens if you want to see them
             // fmt.Printf("  HIDDEN Type: %s, Line: %d, Column: %d, Text: '%s'\n",
             //     tokenTypeName,
             //     token.GetLine(),
             //     token.GetColumn(),
             //     token.GetText(),
             // )
        }
	}

	fmt.Println("\nFinished tokenizing.")
}
