package dialect

import (
	"unicode"

	"github.com/gernest/sequel/stream"
)

type Redshift struct {
	Base
}

func (Redshift) ID() string {
	return "redshift"
}

func (Redshift) IsDelimitedIdentStart(ch rune) bool {
	return ch == '"' || ch == '['
}

func (s Redshift) IsProperIdentInsideQuote(r stream.Runes) bool {
	r.Next()
	stream.Until(r, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	return s.IsIdentStart(r.Peek())
}

func (Redshift) IsIdentStart(ch rune) bool {
	return PostgreSql{}.IsIdentStart(ch) || ch == '#'
}

func (Redshift) IsIdentPart(ch rune) bool {
	return PostgreSql{}.IsIdentPart(ch) || ch == '#'
}

func (Redshift) ConvertTypeBeforeValues() bool {
	return true
}
