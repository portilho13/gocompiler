package lexer

import "fmt"

type Token struct {
	Type    string
	Value  string
}

func CreateToken(tokenType string, value string) Token {
	return Token{tokenType, value}
}

func Lexer() {
	fmt.Println("Hello, World!")
}