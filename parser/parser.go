package parser

import (
	"errors"
	"fmt"

	"github.com/portilho13/gocompiler/lexer"
)

const (
	TYPE_PROGRAM = "PROGRAM"
	TYPE_INCLUDE_DIRECTIVE = "INCLUDE_DIRECTIVE"
	TYPE_STATEMENT = "STATEMENT"
	TYPE_EXPRESSION = "EXPRESSION"
	TYPE_TERM = "TERM"
)

var parser *Parser

type Parser struct {
	tokens []lexer.Token
	index int
}

type FuncDeclaration struct {
	FuncName string
	Params []string
}

type nt struct {
	Type string
	data *lexer.Token
	children []*nt	
}

func get_t() (lexer.Token, error) {
	if len(parser.tokens) - 1 < parser.index {
		return lexer.Token{}, errors.New("no more tokens")
	}
	temp := parser.tokens[parser.index]
	parser.index++
	return temp, nil
}

func unget_t() (lexer.Token, error) {
	if len(parser.tokens) > parser.index {
		return lexer.Token{}, errors.New("no more tokens")
	}
	temp := parser.tokens[parser.index]
	parser.index--
	return temp, nil
}

func get_args() ([]string, error) {
	var args []string
	t, err := get_t()
	for t.Value != ")" || len(parser.tokens) - 1 < parser.index {
		if len(parser.tokens) - 1 < parser.index {
			return nil, errors.New("expected )")
		}
		args = append(args, t.Value)
		t, err = get_t()
		if err != nil {
			return nil, err
		}
	}
	return args, nil
}


func Parse() error {
	parser = &Parser{lexer.GetTokens(), 0}
	for len(parser.tokens) > parser.index {
		t, err := get_t()
		if err != nil {
			return err
		}
		var res []string
		switch {
			case t.Type == lexer.KEYWORD:
				t, err = get_t()
				if err != nil {
					return err
				}
				res = append(res, t.Value)
				if t.Type == lexer.IDENTIFIER || t.Type == lexer.LITERAL {
					t, err = get_t()
					if err != nil {
						return err
					}
					if (t.Type == lexer.DELIMITER && t.Value == "(") {
						args, err := get_args()
						if err != nil {
							return err
						}
						fmt.Println("Funcao: ", res[0])
						fmt.Println(args)
					}
				}

		}
	}
	fmt.Println(parser.tokens)

    return nil
}
