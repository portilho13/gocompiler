package assembler

import (
	"fmt"
	"os"

	"github.com/portilho13/gocompiler/parser"
)

func CreateFile() (*os.File, error) {
	file, err := os.Create("output.asm")
	if err != nil {
		return nil, err
	}
	return file, nil
}

func AssembleFuncDeclaration(file *os.File, funcDecl *parser.Nt) error {
	_, err := file.WriteString(fmt.Sprintf("%s:\n\t", funcDecl.FuncDeclaration.FuncName))
	if err != nil {
		return err
	}
	// Add logic to handle parameters, local variables, and function body
	return nil
}

func Assemble(root *parser.Nt) error {
	file, err := CreateFile()
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("section .text\n")
	if err != nil {
		return err
	}

	for _, child := range root.Children {
		if child.Type == parser.TYPE_FUNC_DECLARATION {
			err = AssembleFuncDeclaration(file, child)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
