package lex_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	lex "github.com/scottshotgg/express-lex"
	token "github.com/scottshotgg/express-token"
)

var (
	l *lex.Lexer
	// TODO: one thing the old lexer architecture fixed was the space at the end
	simpleTest = `
	package something
	i++
	
	c {
		int i = 0;
	}

	import "me"
	include "me"

	// 🔥 comments r kewl 🔥
	var 👌 = "hey, it's \"me\" 😏" +	5 + 10.2/* WOAH YEAH */ ;

	string pokemans = "Woah! That's super effective!"

	floatyMcFloatFace := -66.67383824732894

	object 宇宙カウボーイ = {
		космос: "ковбой"
	}

	char[] bae_toe_ven = "i got luv 4 tha street"

	宇宙カウボーイ["космос"] = bae_toe_ven + 666
	char[] me

	interface i = {}`
)

func TestNew(t *testing.T) {
	l = lex.New(simpleTest)
	fmt.Printf("Lexer: %+v\n", l)
}

func TestLex(t *testing.T) {
	TestNew(t)

	lexemes, err := l.Lex()
	if err != nil {
		panic(err)
	}

	// fmt.Printf("lexemes: %+v\n", lexemes)
	for i, lexeme := range lexemes {
		fmt.Println("i", i, lexeme)
	}

	lexemeJSON, err := json.Marshal(lexemes)
	if err != nil {
		fmt.Println("jsonErr", err)
	}

	fmt.Println(string(lexemeJSON))
}

func TestNewFromFile(t *testing.T) {
	lexer, err := lex.NewFromFile("../test/programs/struct.expr")
	if err != nil {
		fmt.Println("NewFromFile", err)
		os.Exit(9)
	}

	lexTokens, err := lexer.Lex()
	if err != nil {
		fmt.Println("LexErr", err)
		os.Exit(9)
	}

	token.PrintTokens(lexTokens, "\t")
}
