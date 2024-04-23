package main

import (
	"fmt"

	"github.com/portilho13/gocompiler/lexer"
)

func main() {
	l := lexer.CreateToken("INT", "42")
	fmt.Println(l)
}