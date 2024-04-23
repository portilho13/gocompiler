package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/portilho13/gocompiler/lexer"
)

func openFile() (string, error) {
	if len(os.Args) != 2 {
		return "", errors.New("usage: go run main.go <file>")
	}

	return os.Args[1], nil

}

func main() {
	fileName, err := openFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	
	lexer.Lexer(string(content))
}