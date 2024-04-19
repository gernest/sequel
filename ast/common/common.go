package common

import (
	"bytes"
	"fmt"
	"strings"
)

type Ident struct {
	Value      string
	QuoteStyle *rune
}

func (o Ident) Format(f fmt.State, verb rune) {
	if o.QuoteStyle != nil {
		switch *o.QuoteStyle {
		case '\'', '"':
			fmt.Fprintf(f, "%c%v%c", *o.QuoteStyle,
				EscapeQuotedString{Value: o.Value, Quote: *o.QuoteStyle}, *o.QuoteStyle)
		default:
			panic("unexpected quote style")
		}
	} else {
		f.Write([]byte(o.Value))
	}
}

type ObjectName []Ident

func (o ObjectName) Format(f fmt.State, verb rune) {
	Separated(f, []byte("."), []Ident(o)...)
}

func Separated[T fmt.Formatter](f fmt.State, sep []byte, elem ...T) {
	for i, v := range elem {
		if i != 0 {
			f.Write(sep)
		}
		v.Format(f, 0)
	}
}

func EscapeSingleQuote(s string) EscapeQuotedString {
	return EscapeQuotedString{
		Value: s, Quote: '\'',
	}
}

func EscapeDoubleQuote(s string) EscapeQuotedString {
	return EscapeQuotedString{
		Value: s, Quote: '"',
	}
}

type EscapeQuotedString struct {
	Value string
	Quote rune
}

func (e EscapeQuotedString) Format(f fmt.State, verb rune) {
	r := strings.NewReader(e.Value)
	o := new(bytes.Buffer)
	qoute := e.Quote
	var prev rune
	for {
		ch, _, err := r.ReadRune()
		if err != nil {
			f.Write(o.Bytes())
			return
		}
		if ch == qoute {
			if prev == '\\' {
				o.WriteRune(ch)
				continue
			}
			if peek(r) == qoute {
				o.WriteRune(ch)
				o.WriteRune(ch)
				r.ReadRune()
			} else {
				o.WriteRune(ch)
				o.WriteRune(ch)
			}
		} else {
			o.WriteRune(ch)
		}
		prev = ch

	}
}

func peek(r *strings.Reader) (ch rune) {
	defer r.UnreadByte()
	ch, _, _ = r.ReadRune()
	return
}
