package stream

import (
	"bytes"
	"strconv"

	"github.com/gernest/sequel/ascii"
)

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
	buf []byte
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
	b.buf = p
}

func (b *Bytes) CollectUntil(f func(rune) bool) []byte {
	start := b.pos
	for b.HasNext() && f(b.Peek()) {
		b.Next()
	}
	return b.buf[start:b.pos]
}

func (b *Bytes) Int32(n int, base int) rune {
	v, err := strconv.ParseInt(string(b.read(n)), base, 32)
	if err != nil {
		return -1
	}
	return rune(v)
}

func (b *Bytes) Hex() rune {
	start := b.pos
	for range 2 {
		if !ascii.IsHex(b.Peek()) {
			break
		}
		b.Next()
	}
	o := b.buf[start:b.pos]
	if len(o) == 0 {
		return 'x'
	}
	v, err := strconv.ParseInt(string(o), 16, 32)
	if err != nil {
		return -1
	}
	return rune(v)
}

func (b *Bytes) Octet() rune {
	start := b.pos
	for range 2 {
		if !ascii.IsOctet(b.Peek()) {
			break
		}
		b.Next()
	}
	o := b.buf[start:b.pos]
	v, err := strconv.ParseInt(string(o), 8, 32)
	if err != nil {
		return -1
	}
	return rune(v)
}

func (b *Bytes) read(n int) []byte {
	start := b.pos
	for range n {
		b.Next()
	}
	return b.buf[start:b.pos]
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
