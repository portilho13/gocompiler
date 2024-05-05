package parser

import (
	"fmt"
	"strings"
)

func Display(pt *Nt) {
	print_PT(pt, 0, true)
}

func PrintToken(token *Nt) {
	if token.FuncDeclaration != nil {
		fmt.Printf("Type: %s, Value: %s\n", token.Type, token.FuncDeclaration)
	} else if token.VarDeclaration != nil {
		fmt.Printf("Type: %s, Value: %s\n", token.Type, token.VarDeclaration)
	} else {
		fmt.Printf("Type: %s\n", token.Type)
	}
}

func print_PT(pt *Nt, indent int, isLast bool) {
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
	for i, child := range pt.Children {
		print_PT(child, indent+2, i == len(pt.Children)-1)
	}
}