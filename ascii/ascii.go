package ascii

func Lower(ch rune) bool {
	return 'a' <= ch && ch <= 'z'
}

func Upper(ch rune) bool {
	return 'A' <= ch && ch <= 'Z'
}

func Digit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func IsIdentStart(ch rune) bool {
	return Lower(ch) || Upper(ch)
}

func IsIdentPart(ch rune) bool {
	return Lower(ch) || Upper(ch) || Digit(ch) || ch == '_'
}

func IsHex(r rune) bool {
	return (r >= '0' && r <= '9' || r >= 'a' && r <= 'f' || r >= 'A' && r <= 'F')
}

func IsOctet(r rune) bool {
	return (r >= '0' && r <= '7')
}
