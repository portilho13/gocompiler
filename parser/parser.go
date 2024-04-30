package parser

import (
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

type nt struct {
	Type string
	data *lexer.Token
	children []*nt	
}


func Parse() error {
	parser = &Parser{lexer.GetTokens(), 0}
	fmt.Println(parser.tokens)

    return nil
}
