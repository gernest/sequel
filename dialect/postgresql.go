package dialect

import (
	"unicode"

	"github.com/gernest/sequel/ascii"
)

type PostgreSql struct {
	Base
}

func (PostgreSql) ID() string {
	return "postgresql"
}

func (PostgreSql) IdentQuoteStyle(_ string) *rune {
	var q = '"'
	return &q
}

func (PostgreSql) IsIdentStart(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func (PostgreSql) IsIdentPart(ch rune) bool {
	return unicode.IsLetter(ch) ||
		ascii.Digit(ch) || ch == '$' || ch == '_'
}

func (PostgreSql) SupportsFilterDuringAggregation() bool {
	return true
}

func (PostgreSql) SupportsGroupByExpr() bool {
	return true
}
