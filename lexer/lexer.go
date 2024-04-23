package lexer

import (
	"fmt"
	"strings"
)

type Token struct {
	Type    string
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

func CreateToken(tokenType string, value string) Token {
	return Token{tokenType, value}
}

func Lexer(content string) {
	content = removeComments(content)
	fmt.Println(content)


}