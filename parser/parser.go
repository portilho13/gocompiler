package parser

import (
	"fmt"
	"strings"

	"github.com/portilho13/gocompiler/lexer"
)

const (
	TYPE_PROGRAM = iota
	TYPE_STATEMENT
	TYPE_EXPRESSION
	TYPE_TERM
)

type nt struct {
	Type interface{}
	data *lexer.Token
	left_child *nt
	right_child *nt
}

var pt *nt

func createNodeTree(Type interface{}, data *lexer.Token) *nt {
	return &nt{Type, data, nil, nil}
}

func PrintToken(token *lexer.Token) {
	fmt.Printf("Type: %s, Value: %s\n", token.Type, token.Value)
}

func print_PT(pt *nt, indent int, isLast bool) {
	if pt == nil {
		return
	}

	// Print indentation
	if indent > 0 {
		space := strings.Repeat(" ", indent)
		if isLast {
			fmt.Printf("%s└─", space)
		} else {
			fmt.Printf("%s├─", space)
		}
	}

	// Print token
	PrintToken(pt.data)

	// Traverse left child
	print_PT(pt.left_child, indent+2, false)

	// Traverse right sibling
	if pt.right_child != nil {
		print_PT(pt.right_child, indent, true)
	}
}

func Display() {
	print_PT(pt, 0, true)
}


func Parse() {
	tokenList := lexer.GetTokens()
	if len(tokenList) == 0 {
		panic("No tokens to parse")
	}
	root := createNodeTree(TYPE_PROGRAM, nil)

}