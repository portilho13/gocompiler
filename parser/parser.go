package parser

import (
	"errors"
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

type nt struct {
	Type string
	data *lexer.Token
	children []*nt
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

func Parse() error {
    tokenList := lexer.GetTokens()
    if len(tokenList) == 0 {
        panic("No tokens to parse")
    }
    root := createNodeTree(TYPE_PROGRAM, nil)

    fmt.Println("----------------- PARSER -----------------")
	for i, token := range tokenList {
		if token.Type == lexer.DIRECTIVE {
			c := lexer.Peek(i + 1) // Use i+1 to peek at the next token
			if c != nil && c.Type == lexer.HEADER {
				root.AddChild(createNodeTree(TYPE_INCLUDE_DIRECTIVE, c))
			} else {
				return errors.New("missing word #include")
			}
		} else if token.Type == lexer.HEADER {
			if lexer.Peek(i - 1).Type != lexer.DIRECTIVE {
				return errors.New("missing word #include")
			}
		} else {
			fmt.Printf("Token not chosen: %s, Value: %s", token.Type, token.Value)
			fmt.Println()
		}
	}
	

    Display(root)
    return nil
}
