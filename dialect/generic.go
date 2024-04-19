package dialect

import "unicode"

type Generic struct {
	Base
}

func (Generic) ID() string {
	return "generic"
}

func (Generic) IsIdentStart(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_' || ch == '#' || ch == '@'
}

func (Generic) IsIdentPart(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '@' || ch == '$' || ch == '#' || ch == '_'
}

func (Generic) SupportsGroupByExpr() bool {
	return true
}
func (Generic) SupportsStartTxnModifier() bool {
	return true
}
