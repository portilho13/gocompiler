package lexer

import (
	"errors"
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
var file File

type Token struct {
	Type   int
	Value  string
}

type File struct {
	length int
	content string
	current int
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

func getc(content string) (rune, error){
	if file.length < 0 {
		return 0, errors.New("EOF")
	}
	file.current++
	c := rune(content[file.current - 1])
	return c, nil
}


func Lexer(content string) {
	content = removeComments(content)
	tokenList = append(tokenList, CreateToken(LITERAL, "1"))
	content = strings.TrimSpace(content)
	fmt.Println(content)

	file = File{length: len(content), content: content, current: 0}

	for i := 0; i < len(content); i++ {
		c, err := getc(content)
		switch content[i] {
			case ' ':
	}


}