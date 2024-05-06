package assembler

import (
	"fmt"
	"os"

	"github.com/portilho13/gocompiler/parser"
)

type Assembler struct {
	start []string
	funcs []*parser.Nt
}

func checkFileExists() bool {
	if _, err := os.Stat("output.asm"); os.IsNotExist(err) {
		return false
	}
	return true
}

func DeleteFile() error {
	err := os.Remove("output.asm")
	if err != nil {
		return err
	}
	return nil
}

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
	assemble := Assembler{}
	if checkFileExists() {
		err := DeleteFile()
		if err != nil {
			return err
		}
	}
	file, err := CreateFile()
	if err != nil {
		return err
	}
	defer file.Close()


	for _, child := range root.Children {
		if child.Type == parser.TYPE_FUNC_DECLARATION {
			assemble.start = append(assemble.start, child.FuncDeclaration.FuncName)
			assemble.funcs = append(assemble.funcs, child)
		}
	}

	return nil
}
