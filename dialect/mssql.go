package dialect

import (
	"unicode"

	"github.com/gernest/sequel/ascii"
)

type MSSQL struct {
	Base
}

func (MSSQL) ID() string {
	return "mssql"
}

func (MSSQL) IsDelimitedIdentStart(ch rune) bool {
	return ch == '"' || ch == '['
}

func (MSSQL) IsIdentStart(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_' || ch == '#' || ch == '@'
}

func (MSSQL) IsIdentPart(ch rune) bool {
	return unicode.IsLetter(ch) || ascii.Digit(ch) ||
		ch == '$' || ch == '_' || ch == '#' || ch == '@'
}

func (MSSQL) ConvertTypeBeforeValues() bool {
	return true
}
