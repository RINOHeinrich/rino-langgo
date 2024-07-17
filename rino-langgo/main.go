package main

import (
	"fmt"

	"github.com/RINOHeinrich/olang/utils"
)

func main() {
	var Lexer utils.Lexer
	text := "Hello , world ! This is a sample text."
	Lexer.Analyse(text)
	for _, token := range *Lexer.GetTokenSlice() {
		switch token.GetTokenType() {
		case utils.WORD:
			fmt.Println("WORD: ", token.GetValue())
		case utils.NUMBER:
			fmt.Println("NUMBER: ", token.GetValue())
		case utils.PUNCTUATION:
			fmt.Println("PUNCTUATION: ", token.GetValue())
		case utils.SPACE:
			fmt.Println("SPACE: ", token.GetValue())
		case utils.NEWLINE:
			fmt.Println("NEWLINE: ", token.GetValue())
		}
	}

}
