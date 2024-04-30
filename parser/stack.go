package parser

import "github.com/portilho13/gocompiler/lexer"

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