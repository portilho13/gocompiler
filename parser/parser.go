package parser

import (
	"github.com/portilho13/gocompiler/lexer"
)

const (
	TYPE_PROGRAM = "PROGRAM"
	TYPE_INCLUDE_DIRECTIVE = "INCLUDE_DIRECTIVE"
	TYPE_STATEMENT = "STATEMENT"
	TYPE_EXPRESSION = "EXPRESSION"
	TYPE_TERM = "TERM"
)

var stack Stack

type Parser struct {
	tokens []lexer.Token
	index int
}

type nt struct {
	Type string
	data *lexer.Token
	children []*nt	
}

type BinExpr struct {
	leftOp string
	Op string
	rightOp string
}

func Parse() error {

    return nil
}
