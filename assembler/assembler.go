package assembler

import (
	"fmt"
	"os"

	"github.com/portilho13/gocompiler/parser"
)

var file *os.File

func CreateFile() (*os.File, error) {
	file, err := os.Create("output.asm")
	if err != nil {
		return nil, err
	}
	return file, nil
}

func Assemble(root *parser.Nt) error {
	_, err := CreateFile()
	if err != nil {
		return err
	}
	os.WriteFile("output.asm", []byte("section .text\n"), 0644)
	for _, child := range root.Children {
		if child.Type == parser.TYPE_FUNC_DECLARATION {
			fmt.Println("Function declaration")
		}
	}
	return nil
}