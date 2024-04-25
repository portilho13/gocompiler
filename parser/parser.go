package parser

import (
	"fmt"
	"strings"

	"github.com/portilho13/gocompiler/lexer"
)

const (
	TYPE_PROGRAM = "PROGRAM"
	TYPE_STATEMENT = "STATEMENT"
	TYPE_EXPRESSION = "EXPRESSION"
	TYPE_TERM = "TERM"
)

type nt struct {
	Type string
	data *lexer.Token
	children []*nt
}

var pt *nt

func createNodeTree(Type string, data *lexer.Token) *nt {
	return &nt{Type, data, make([]*nt, 0)}

}

func PrintToken(token *nt) {
	fmt.Printf("Type: %s, Value: %s\n", token.Type, token.data.Value)
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
	PrintToken(pt)

	// Traverse child nodes
	for i, child := range pt.children {
		print_PT(child, indent+2, i == len(pt.children)-1)
	}
}

func Display(pt *nt) {
	print_PT(pt, 0, true)
}


func Parse() {
	tokenList := lexer.GetTokens()
	if len(tokenList) == 0 {
		panic("No tokens to parse")
	}
	root := createNodeTree(TYPE_PROGRAM, &tokenList[0])
	Display(root)

}