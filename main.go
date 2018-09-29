package lex

import (
	"fmt"
	"os"

	"github.com/scottshotgg/express-token"
)

// Lexemes ...
var (
	Lexemes = []string{
		"var",
		"int",
		"float",
		"string",
		"bool",
		"char",
		"object",

		":",
		"=",
		"+",
		"-",
		"*",
		"/",
		"(",
		")",
		"{",
		"}",
		"[",
		"]",
		"\"",
		"'",
		";",
		",",
		"#",
		"!",
		"<",
		">",
		"@",
		"\\",
		// "„",
		" ",
		"\n",
		"\t",

		// "select",
		// "SELECT",
		// "FROM",
		// "WHERE",
	}

	// LexemeMap is used for holding the lexemes that will be used to identify tokens in the lexer
	LexemeMap = map[string]token.Token{}
)

// Initialize the lexeme map and ensure that all defined lexemes can be mapped to a token
func init() {
	var ok bool

	for _, lexeme := range Lexemes {
		LexemeMap[lexeme], ok = token.TokenMap[lexeme]
		if !ok {
			fmt.Println("ERROR: Lexeme not found in TokenMap: ", lexeme)
			os.Exit(7)
		}
	}
}
