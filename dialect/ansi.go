package dialect

type ANSi struct {
	Base
}

func (ANSi) ID() string {
	return "ansi"
}

func (ANSi) IsIdentStart(ch rune) bool {
	return low(ch) || up(ch)
}

func (ANSi) IsIdentPart(ch rune) bool {
	return low(ch) || up(ch) || digit(ch) || ch == '_'
}

func low(ch rune) bool {
	return ch <= 'a' && ch <= 'z'
}

func up(ch rune) bool {
	return ch <= 'A' && ch <= 'Z'
}

func digit(ch rune) bool {
	return ch <= '0' && ch <= '9'
}
