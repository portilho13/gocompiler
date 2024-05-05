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
	TYPE_FUNC_DECLARATION = "FUNC_DECLARATION"
	TYPE_VA = "VAR_DECLARATION"
)

var parser *Parser
var root *Nt

type Parser struct {
	tokens []lexer.Token
	index int
}

type VarDeclaration struct {
	VarName string
	VarType string
	Value string
}

type FuncDeclaration struct {
	FuncName string
	Params []string
}

type Nt struct {
	Type string
	FuncDeclaration *FuncDeclaration
	VarDeclaration *VarDeclaration
	Children []*Nt	

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
	if parser.index == 0 {
		return lexer.Token{}, errors.New("no more tokens")
	}
	parser.index--
	temp := parser.tokens[parser.index]
	return temp, nil
}

func get_var_type()	string {
	temp := parser.index
	t := lexer.Token{}
	var err error
	for i := 0; i < 3; i++ {
		t, err = unget_t()
		if err != nil {
			return ""
		}
	}
	parser.index = temp
	return t.Value
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


func Parse() (*Nt, error) {
	root = &Nt{TYPE_PROGRAM, nil, nil, nil}
	parser = &Parser{lexer.GetTokens(), 0}
	for len(parser.tokens) > parser.index {
		t, err := get_t()
		if err != nil {
			return nil, err
		}
		var res []string
		switch {
			case t.Type == lexer.KEYWORD:
				t, err = get_t()
				if err != nil {
					return nil, err
				}
				res = append(res, t.Value)
				if t.Type == lexer.IDENTIFIER || t.Type == lexer.LITERAL {
					t, err = get_t()
					if err != nil {
						return nil, err
					}
					if (t.Type == lexer.DELIMITER && t.Value == "(") {
						args, err := get_args()
						if err != nil {
							return nil, err
						}
						fd := FuncDeclaration{res[0], args}
						root.Children = append(root.Children, &Nt{TYPE_FUNC_DECLARATION, &fd, nil, nil})
						
					} else if t.Type == lexer.DELIMITER && t.Value == ";" {
						tp := get_var_type()
						fmt.Printf("Var type: %s\n", tp)
						vd := VarDeclaration{res[0], tp, ""}
						root.Children[0].Children = append(root.Children[0].Children, &Nt{TYPE_VA, nil, &vd, nil})
					}
				}

		}
	}

    return root, nil
}
