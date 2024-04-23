package parser

import (
	"fmt"
	"strings"

	"github.com/portilho13/gocompiler/lexer"
)

var pt *ParseTree

type ParseTree struct {
	data *lexer.Token
	child *ParseTree
	sibling *ParseTree
}

func CreateParseTree(data *lexer.Token) *ParseTree {
	return &ParseTree{data, nil, nil}
}

func PrintToken(token *lexer.Token) {
	fmt.Printf("Type: %s, Value: %s\n", token.Type, token.Value)
}

func print_PT(pt *ParseTree, indent int) {
	if pt == nil {
		return
	}

	// Depth-First print
	if indent > 0 {
		space := strings.Repeat(" ", indent)
		fmt.Printf("%s", space)
	}

	// Print token
	PrintToken(pt.data)

	sibling := pt.sibling
	child := pt.child
	// Traverse child and sibling nodes
	print_PT(child, indent+2)
	print_PT(sibling, indent)

}


func Display() {
	print_PT(pt, 0)
}

func Parse() {
	fmt.Println("Parse")
}