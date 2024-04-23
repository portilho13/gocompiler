package lexer

import (
	"fmt"
	"strings"
)

const (
	KEYWORD = iota
	IDENTIFIER
	LITERAL
	OPERATOR
	DELIMITER
)

var tokenList []Token

type Token struct {
	Type   int
	Value  string
}

func removeComments(content string) string {
	lines := strings.Split(content, "\n")
	var result []string

	for _, line := range lines {
		if !strings.HasPrefix(line, "//") {
			result = append(result, line)
		}
	}
	return strings.Join(result, "\n")
}

func CreateToken(tokenType int, value string) Token {
	return Token{KEYWORD, value}
}

func Display() {
	for _, token := range tokenList {
		fmt.Printf("Type: %d, Value: %s\n", token.Type, token.Value)
	}
}

func Lexer(content string) {
	content = removeComments(content)
	tokenList = append(tokenList, CreateToken(LITERAL, "1"))
	fmt.Println(content)


}