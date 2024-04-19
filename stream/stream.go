package stream

import "bytes"

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

type Bytes struct {
	r   bytes.Reader
	pos int
	err error
}

func New(p []byte) *Bytes {
	var b Bytes
	b.Reset(p)
	return &b
}

var _ Runes = (*Bytes)(nil)

func (b *Bytes) HasNext() bool {
	return b.err == nil && b.pos < b.r.Len()
}

func (b *Bytes) Reset(p []byte) {
	b.r.Reset(p)
	b.pos = 0
	b.err = nil
}
func (b *Bytes) Next() rune {
	r, size, err := b.r.ReadRune()
	if err != nil {
		b.err = err
		return 0
	}
	b.pos += size
	return r
}

func (b *Bytes) Peek() (n rune) {
	n, _, _ = b.r.ReadRune()
	b.r.UnreadRune()
	return
}
