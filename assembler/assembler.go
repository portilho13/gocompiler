package assembler

import (
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

func Assemble(root *parser.Nt) error {
	_, err := CreateFile()
	if err != nil {
		return err
	}
	for _, child := range root.Children {
		
		Assemble(child)
	}
	return nil
}