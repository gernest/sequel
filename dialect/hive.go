package dialect

import "github.com/gernest/sequel/ascii"

type Hive struct {
	Base
}

func (Hive) ID() string {
	return "hive"
}

func (Hive) IsDelimitedIdentStart(ch rune) bool {
	return (ch == '"') || (ch == '`')
}

func (Hive) IsIdentStart(ch rune) bool {
	return ascii.Lower(ch) || ascii.Upper(ch) || ascii.Digit(ch) || ch == '$'
}

func (Hive) IsIdentPart(ch rune) bool {
	return ascii.Lower(ch) ||
		ascii.Upper(ch) ||
		ascii.Digit(ch) ||
		ch == '_' ||
		ch == '$' ||
		ch == '{' ||
		ch == '}'
}

func (Hive) SupportsFilterDuringAggregation(ch rune) bool {
	return true
}
