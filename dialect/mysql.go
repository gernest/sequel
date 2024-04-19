package dialect

import (
	"unicode"

	"github.com/gernest/sequel/ascii"
)

type MySql struct {
	Base
}

func (MySql) ID() string {
	return "mysql"
}

func (MySql) IsIdentStart(ch rune) bool {
	return unicode.IsLetter(ch) ||
		ch == '_' ||
		ch == '$' ||
		ch == '@' || in(ch)

}

func in(ch rune) bool {
	return '\u0080' <= ch && ch <= '\uffff'
}

func (MySql) IsIdentPart(ch rune) bool {
	return MySql{}.IsIdentStart(ch) || ascii.Digit(ch)
}

func (MySql) IsDelimitedIdentStart(ch rune) bool {
	return ch == '`'
}

func (MySql) IdentQuoteStyle(ch rune) *rune {
	var backTick = '`'
	return &backTick
}
