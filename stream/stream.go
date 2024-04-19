package stream

type Runes interface {
	HasNext() bool
	Next() rune
	Peek() rune
}

func Until(stream Runes, f func(r rune) bool) {
	for stream.HasNext() && f(stream.Peek()) {
		stream.Next()
	}
}
