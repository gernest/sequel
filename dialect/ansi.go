package dialect

import "github.com/gernest/sequel/ascii"

type ANSi struct {
	Base
}

func (ANSi) ID() string {
	return "ansi"
}

func (ANSi) IsIdentStart(ch rune) bool {
	return ascii.IsIdentStart(ch)
}

func (ANSi) IsIdentPart(ch rune) bool {
	return ascii.IsIdentPart(ch)
}
