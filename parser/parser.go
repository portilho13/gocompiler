package parser

import (
	"fmt"
	"strings"

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

type nt struct {
	Type string
	data *lexer.Token
	children []*nt
}

type Stack struct {
	items []*lexer.Token
}

func (s *Stack) Push(item *lexer.Token) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() *lexer.Token {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() *lexer.Token {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}


func createNodeTree(Type string, data *lexer.Token) *nt {
	return &nt{Type, data, make([]*nt, 0)}

}

func PrintToken(token *nt) {
	if token.data == nil {
		fmt.Printf("Type: %s\n", token.Type)
		return
	}
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

func (n *nt) AddChild(child *nt) {
	n.children = append(n.children, child)
}

func Display(pt *nt) {
	print_PT(pt, 0, true)
}

func isHeader(token *lexer.Token, i int) (bool, *lexer.Token) {
	if token.Type == lexer.DIRECTIVE {
		if lexer.Peek(i + 1).Type == lexer.HEADER {
			return true, lexer.Peek(i + 1)
		}
	}
	return false, nil
}

func Parse() error {
    tokenList := lexer.GetTokens()
    if len(tokenList) == 0 {
        panic("No tokens to parse")
    }
    root := createNodeTree(TYPE_PROGRAM, nil)

    fmt.Println("----------------- PARSER -----------------")
	for _, token := range tokenList {
		stack.Push(&token)
	}
	

    Display(root)
    return nil
}
