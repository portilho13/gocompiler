package parser

import (
	"fmt"
	"strings"
)

func Display(pt *nt) {
	print_PT(pt, 0, true)
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