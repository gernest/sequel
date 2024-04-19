package tokenizer

import (
	"strings"

	"github.com/gernest/sequel/token"
)

type State struct {
	r      *strings.Reader
	line   uint64
	column uint64
}

func (s *State) Next() (ch rune, err error) {
	ch, _, err = s.r.ReadRune()
	if err != nil {
		return
	}
	if ch == '\n' {
		s.line++
		s.column = 1
	} else {
		s.column++
	}
	return
}

func (s *State) Peek() (ch rune) {
	ch, _, _ = s.r.ReadRune()
	s.r.UnreadRune()
	return
}

func (s *State) Location() *token.Location {
	return &token.Location{
		Line:   s.line,
		Column: s.column,
	}
}
