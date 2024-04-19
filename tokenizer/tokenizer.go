package tokenizer

import (
	"fmt"
	"strings"

	keyword "github.com/gernest/sequel/ast/keywords"
)

type Token interface {
	fmt.Formatter
	t()
}

type Word struct {
	// The value of the token, without the enclosing quotes, and with the
	// escape sequences (if any) processed (TODO: escapes are not handled)
	Value string
	// An identifier can be "quoted" (&lt;delimited identifier> in ANSI parlance).
	// The standard and most implementations allow using double quotes for this,
	// but some implementations support other quoting styles as well (e.g. \[MS SQL])
	QUoteStyle *rune
	// If the word was not quoted and it matched one of the known keywords,
	// this will have one of the values from dialect::keywords, otherwise empty
	Keyword keyword.KeyWord
}

func (w Word) Format(f fmt.State, verb rune) {
	if w.QUoteStyle != nil {
		fmt.Fprintf(f, "%c%s%c", *w.QUoteStyle, w.Value, *w.QUoteStyle)
	} else {
		f.Write([]byte(w.Value))
	}
}

type WhiteSpace interface {
	fmt.Formatter
	ws()
}

type Space struct{}

var _ WhiteSpace = (*Space)(nil)

func (Space) ws() {}

func (Space) Format(f fmt.State, ver rune) {
	f.Write([]byte(" "))
}

type NewLine struct{}

var _ WhiteSpace = (*NewLine)(nil)

func (NewLine) ws() {}

func (NewLine) Format(f fmt.State, ver rune) {
	f.Write([]byte("\n"))
}

type Tab struct{}

var _ WhiteSpace = (*Tab)(nil)

func (Tab) ws() {}

func (Tab) Format(f fmt.State, ver rune) {
	f.Write([]byte("\t"))
}

type SingleLineComment struct {
	Comment string
	Prefix  string
}

var _ WhiteSpace = (*SingleLineComment)(nil)

func (SingleLineComment) ws() {}

func (s SingleLineComment) Format(f fmt.State, ver rune) {
	fmt.Fprintf(f, "%s%s", s.Prefix, s.Comment)
}

type MultiLineComment struct {
	Comment string
}

var _ WhiteSpace = (*MultiLineComment)(nil)

func (MultiLineComment) ws() {}

func (m MultiLineComment) Format(f fmt.State, ver rune) {
	fmt.Fprintf(f, "/*%s*/", m.Comment)
}

type Location struct {
	Line   uint64
	Column uint64
}

type TokenWithLocation struct {
	Token    Token
	Location *Location
}

type Error struct {
	Message  string
	Location *Location
}

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

func (s *State) Location() *Location {
	return &Location{
		Line:   s.line,
		Column: s.column,
	}
}
