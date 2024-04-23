package lexer

import (
	"errors"
	"fmt"
	"strings"
)

const (
	KEYWORD = "KEYWORD"
	IDENTIFIER = "IDENTIFIER"
	LITERAL = "LITERAL"
	OPERATOR = "OPERATOR"
	DELIMITER = "DELIMITER"
)

var tokenList []Token
var file File

type Token struct {
	Type   string
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

func CreateToken(tokenType string, value string) Token {
	return Token{tokenType, value}
}

func Display() {
	for _, token := range tokenList {
		fmt.Printf("Type: %s, Value: %s\n", token.Type, token.Value)
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

func get_string() string {
	content := file.content
	start := file.current
	for file.current < file.length {
		c := rune(content[file.current])
		if c == '"' {
			break
		}
		file.current++
	}
	return content[start:file.current]

}

func get_keyword(identifier string) (bool, string) {
	keywords := []string{"if", "else", "for", "while", "do", "break", "continue", "int", "float", "char", "double", "short", "long", "unsigned", "signed", "void", "struct", "union", "enum", "typedef", "sizeof", "auto", "register", "static", "extern", "const", "volatile", "return", "switch", "case", "default", "goto", "asm", "inline", "restrict", "_Bool", "_Complex", "_Imaginary"}
	for _, keyword := range keywords {
		if keyword == identifier {
			return true, keyword
		}
	}
	return false, ""

}

func isNumeric(c rune) bool {
	return '0' <= c && c <= '9'
}


func Lexer(content string) {
	content = removeComments(content)
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
				identifier := get_identifier()
				isKeyword, keyword := get_keyword(identifier)
				if isKeyword {
					tokenList = append(tokenList, CreateToken(KEYWORD, keyword))
				} else {
					tokenList = append(tokenList, CreateToken(IDENTIFIER, identifier))
				}
			case c == '"':
				tokenList = append(tokenList, CreateToken(LITERAL, get_string()))
			case c == ';':
				tokenList = append(tokenList, CreateToken(DELIMITER, string(c)))
			case isNumeric(c):
				tokenList = append(tokenList, CreateToken(LITERAL, string(c)))
			default:
				fmt.Printf("Invalid character %c\n", c)
		}
	}


}