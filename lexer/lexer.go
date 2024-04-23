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

func getc() (rune, error){
	content := file.content
	if file.length < 0 {
		return 0, errors.New("EOF")
	}
	c := rune(content[file.current])
	file.current++
	return c, nil
}

func get_identifier() string {
	content := file.content
	start := file.current - 1
	for file.current < file.length {
		c := rune(content[file.current])
		if !('a' <= c && c <= 'z' || 'A' <= c && c <= 'Z') {
			break
		}
		file.current++
	}
	return content[start:file.current]
}


func Lexer(content string) {
	content = removeComments(content)
	tokenList = append(tokenList, CreateToken(LITERAL, "1"))
	content = strings.TrimSpace(content)
	fmt.Println(content)

	file = File{length: len(content), content: content, current: 0}

	for file.current < file.length {
		c, err := getc()
		if err != nil {
			break
		}
		switch {
			case c == ' ' || c == '\n' || c == '\t':
				continue
			case c == '(' || c == ')' || c == '{' || c == '}' || c == '[' || c == ']':
				tokenList = append(tokenList, CreateToken(DELIMITER, string(c)))
			case c == '+' || c == '-' || c == '*' || c == '/' || c == '=':
				tokenList = append(tokenList, CreateToken(OPERATOR, string(c)))
			case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z':
				tokenList = append(tokenList, CreateToken(IDENTIFIER, get_identifier()))
			default:
				fmt.Println("Invalid character")
		}
	}


}