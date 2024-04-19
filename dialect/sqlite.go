package dialect

import "github.com/gernest/sequel/ascii"

type Sqlite struct {
	Base
}

func (Sqlite) ID() string {
	return "sqlite"
}

func (Sqlite) IsDelimitedIdentStart(ch rune) bool {
	return ch == '`' || ch == '"' || ch == '['
}

func (Sqlite) IdentQuoteStyle(_ string) *rune {
	q := '`'
	return &q
}

func (Sqlite) IsIdentStart(ch rune) bool {
	return ascii.IsIdentStart(ch) ||
		ch == '_' ||
		ch == '$' || in2(ch)
}

func in2(ch rune) bool {
	return '\u007f' <= ch && ch <= '\uffff'
}

func (Sqlite) SupportsFilterDuringAggregation() bool { return true }
func (Sqlite) SupportsStartTxnModifier() bool        { return true }
func (s Sqlite) IsIdentPart(ch rune) bool {
	return s.IsIdentStart(ch) || ascii.Digit(ch)
}
