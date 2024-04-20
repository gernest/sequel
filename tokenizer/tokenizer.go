package tokenizer

import (
	"strings"

	"github.com/gernest/sequel/ascii"
	"github.com/gernest/sequel/dialect"
	"github.com/gernest/sequel/stream"
	"github.com/gernest/sequel/token"
)

type State struct {
	r      stream.Bytes
	line   uint64
	column uint64
}

func (s *State) Reset(query string) {
	s.r.Reset([]byte(query))
	s.line = 1
	s.column = 1
}

func (s *State) HasNext() bool {
	return s.r.HasNext()
}

func (s *State) Next() (ch rune) {
	ch = s.r.Next()
	if ch == '\n' {
		s.line++
		s.column = 1
	} else {
		s.column++
	}
	return
}

func (s *State) Peek() rune {
	return s.r.Next()
}

func (s *State) Location() *token.Location {
	return &token.Location{
		Line:   s.line,
		Column: s.column,
	}
}

type Tokenizer struct {
	tokens  []token.Token
	state   State
	query   string
	dialect dialect.Dialect
}

func New(query string) *Tokenizer {
	var t Tokenizer
	t.Reset(query)
	return t.WithDialect(dialect.Generic{})
}

func (t *Tokenizer) WithDialect(d dialect.Dialect) *Tokenizer {
	t.dialect = d
	return t
}

func (t *Tokenizer) Reset(query string) {
	t.state.Reset(query)
	t.tokens = t.tokens[:0]
	t.query = query
}

func (t *Tokenizer) NextToken() token.Token {
	if !t.state.HasNext() {
		return token.EOF{}
	}
	ch := t.state.Peek()
	switch ch {
	case ' ':
		return t.consumeAndReturn(token.Space{})
	case '\t':
		return t.consumeAndReturn(token.Tab{})
	case '\n':
		return t.consumeAndReturn(token.NewLine{})
	case '\r':
		t.state.Next()
		if t.state.Peek() == '\n' {
			t.state.Next()
		}
		return token.NewLine{}
	default:
		return nil
	}
}

func (t *Tokenizer) consumeAndReturn(tok token.Token) token.Token {
	t.state.Next()
	return tok
}

func (t *Tokenizer) collectUntil(f func(rune) bool) []byte {
	return t.state.r.CollectUntil(f)
}

func unescape(s string) string {
	o := new(strings.Builder)
	var b stream.Bytes
	b.Reset([]byte(s))
	b.Next()
	for b.HasNext() {
		c := b.Next()
		if c == '\'' {
			if b.Peek() == '\'' {
				b.Next()
				o.WriteRune('\'')
				continue
			}
			return o.String()
		}
		if c != '\\' {
			o.WriteRune(c)
			continue
		}
		c = b.Next()
		switch c {
		case 'b':
			o.WriteRune('\u0008')
		case 'f':
			o.WriteRune('\u000C')
		case 'n':
			o.WriteRune('\n')
		case 'r':
			o.WriteRune('\r')
		case 't':
			o.WriteRune('\t')
		case 'u':
			n := b.Int32(4, 16)
			if n != -1 {
				o.WriteRune(n)
			}
		case 'U':
			n := b.Int32(8, 16)
			if n != -1 {
				o.WriteRune(n)
			}
		default:
			if ascii.IsOctet(c) {
				n := b.Octet()
				if n != -1 {
					o.WriteRune(n)
				}
			} else {
				o.WriteRune(c)
			}

		}
	}
	return o.String()
}
