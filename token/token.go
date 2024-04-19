package token

import (
	"fmt"

	keyword "github.com/gernest/sequel/ast/keywords"
)

type Token interface {
	fmt.Formatter
	t()
}

type EOF struct{}

var _ Token = (*EOF)(nil)

func (EOF) t() {}

func (t EOF) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Number struct{}

var _ Token = (*Number)(nil)

func (Number) t() {}

func (t Number) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Char struct{}

var _ Token = (*Char)(nil)

func (Char) t() {}

func (t Char) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type SingleQuotedString struct{}

var _ Token = (*SingleQuotedString)(nil)

func (SingleQuotedString) t() {}

func (t SingleQuotedString) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type DoubleQuotedString struct{}

var _ Token = (*DoubleQuotedString)(nil)

func (DoubleQuotedString) t() {}

func (t DoubleQuotedString) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type DollarQuotedString struct{}

var _ Token = (*DollarQuotedString)(nil)

func (DollarQuotedString) t() {}

func (t DollarQuotedString) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type SingleQuotedByteStringLiteral struct{}

var _ Token = (*SingleQuotedByteStringLiteral)(nil)

func (SingleQuotedByteStringLiteral) t() {}

func (t SingleQuotedByteStringLiteral) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type DoubleQuotedByteStringLiteral struct{}

var _ Token = (*DoubleQuotedByteStringLiteral)(nil)

func (DoubleQuotedByteStringLiteral) t() {}

func (t DoubleQuotedByteStringLiteral) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type RawStringLiteral struct{}

var _ Token = (*RawStringLiteral)(nil)

func (RawStringLiteral) t() {}

func (t RawStringLiteral) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type NationalStringLiteral struct{}

var _ Token = (*NationalStringLiteral)(nil)

func (NationalStringLiteral) t() {}

func (t NationalStringLiteral) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type EscapedStringLiteral struct{}

var _ Token = (*EscapedStringLiteral)(nil)

func (EscapedStringLiteral) t() {}

func (t EscapedStringLiteral) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type HexStringLiteral struct{}

var _ Token = (*HexStringLiteral)(nil)

func (HexStringLiteral) t() {}

func (t HexStringLiteral) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Comma struct{}

var _ Token = (*Comma)(nil)

func (Comma) t() {}

func (t Comma) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type DoubleEq struct{}

var _ Token = (*DoubleEq)(nil)

func (DoubleEq) t() {}

func (t DoubleEq) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Eq struct{}

var _ Token = (*Eq)(nil)

func (Eq) t() {}

func (t Eq) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Neq struct{}

var _ Token = (*Neq)(nil)

func (Neq) t() {}

func (t Neq) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Lt struct{}

var _ Token = (*Lt)(nil)

func (Lt) t() {}

func (t Lt) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Gt struct{}

var _ Token = (*Gt)(nil)

func (Gt) t() {}

func (t Gt) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type LtEq struct{}

var _ Token = (*LtEq)(nil)

func (LtEq) t() {}

func (t LtEq) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type GtEq struct{}

var _ Token = (*GtEq)(nil)

func (GtEq) t() {}

func (t GtEq) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Spaceship struct{}

var _ Token = (*Spaceship)(nil)

func (Spaceship) t() {}

func (t Spaceship) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Plus struct{}

var _ Token = (*Plus)(nil)

func (Plus) t() {}

func (t Plus) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Minus struct{}

var _ Token = (*Minus)(nil)

func (Minus) t() {}

func (t Minus) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Mul struct{}

var _ Token = (*Mul)(nil)

func (Mul) t() {}

func (t Mul) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Div struct{}

var _ Token = (*Div)(nil)

func (Div) t() {}

func (t Div) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type DuckIntDiv struct{}

var _ Token = (*DuckIntDiv)(nil)

func (DuckIntDiv) t() {}

func (t DuckIntDiv) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Mod struct{}

var _ Token = (*Mod)(nil)

func (Mod) t() {}

func (t Mod) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type StringConcat struct{}

var _ Token = (*StringConcat)(nil)

func (StringConcat) t() {}

func (t StringConcat) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type LParen struct{}

var _ Token = (*LParen)(nil)

func (LParen) t() {}

func (t LParen) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type RParen struct{}

var _ Token = (*RParen)(nil)

func (RParen) t() {}

func (t RParen) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Period struct{}

var _ Token = (*Period)(nil)

func (Period) t() {}

func (t Period) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Colon struct{}

var _ Token = (*Colon)(nil)

func (Colon) t() {}

func (t Colon) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type DoubleColon struct{}

var _ Token = (*DoubleColon)(nil)

func (DoubleColon) t() {}

func (t DoubleColon) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Assignment struct{}

var _ Token = (*Assignment)(nil)

func (Assignment) t() {}

func (t Assignment) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type SemiColon struct{}

var _ Token = (*SemiColon)(nil)

func (SemiColon) t() {}

func (t SemiColon) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Backslash struct{}

var _ Token = (*Backslash)(nil)

func (Backslash) t() {}

func (t Backslash) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type LBracket struct{}

var _ Token = (*LBracket)(nil)

func (LBracket) t() {}

func (t LBracket) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type RBracket struct{}

var _ Token = (*RBracket)(nil)

func (RBracket) t() {}

func (t RBracket) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Ampersand struct{}

var _ Token = (*Ampersand)(nil)

func (Ampersand) t() {}

func (t Ampersand) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Pipe struct{}

var _ Token = (*Pipe)(nil)

func (Pipe) t() {}

func (t Pipe) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Caret struct{}

var _ Token = (*Caret)(nil)

func (Caret) t() {}

func (t Caret) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type LBrace struct{}

var _ Token = (*LBrace)(nil)

func (LBrace) t() {}

func (t LBrace) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type RBrace struct{}

var _ Token = (*RBrace)(nil)

func (RBrace) t() {}

func (t RBrace) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type RArrow struct{}

var _ Token = (*RArrow)(nil)

func (RArrow) t() {}

func (t RArrow) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Sharp struct{}

var _ Token = (*Sharp)(nil)

func (Sharp) t() {}

func (t Sharp) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Tilde struct{}

var _ Token = (*Tilde)(nil)

func (Tilde) t() {}

func (t Tilde) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type TildeAsterisk struct{}

var _ Token = (*TildeAsterisk)(nil)

func (TildeAsterisk) t() {}

func (t TildeAsterisk) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type ExclamationMarkTilde struct{}

var _ Token = (*ExclamationMarkTilde)(nil)

func (ExclamationMarkTilde) t() {}

func (t ExclamationMarkTilde) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type ExclamationMarkTildeAsterisk struct{}

var _ Token = (*ExclamationMarkTildeAsterisk)(nil)

func (ExclamationMarkTildeAsterisk) t() {}

func (t ExclamationMarkTildeAsterisk) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type DoubleTilde struct{}

var _ Token = (*DoubleTilde)(nil)

func (DoubleTilde) t() {}

func (t DoubleTilde) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type DoubleTildeAsterisk struct{}

var _ Token = (*DoubleTildeAsterisk)(nil)

func (DoubleTildeAsterisk) t() {}

func (t DoubleTildeAsterisk) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type ExclamationMarkDoubleTilde struct{}

var _ Token = (*ExclamationMarkDoubleTilde)(nil)

func (ExclamationMarkDoubleTilde) t() {}

func (t ExclamationMarkDoubleTilde) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type ExclamationMarkDoubleTildeAsterisk struct{}

var _ Token = (*ExclamationMarkDoubleTildeAsterisk)(nil)

func (ExclamationMarkDoubleTildeAsterisk) t() {}

func (t ExclamationMarkDoubleTildeAsterisk) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type ShiftLeft struct{}

var _ Token = (*ShiftLeft)(nil)

func (ShiftLeft) t() {}

func (t ShiftLeft) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type ShiftRight struct{}

var _ Token = (*ShiftRight)(nil)

func (ShiftRight) t() {}

func (t ShiftRight) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Overlap struct{}

var _ Token = (*Overlap)(nil)

func (Overlap) t() {}

func (t Overlap) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type ExclamationMark struct{}

var _ Token = (*ExclamationMark)(nil)

func (ExclamationMark) t() {}

func (t ExclamationMark) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type DoubleExclamationMark struct{}

var _ Token = (*DoubleExclamationMark)(nil)

func (DoubleExclamationMark) t() {}

func (t DoubleExclamationMark) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type AtSign struct{}

var _ Token = (*AtSign)(nil)

func (AtSign) t() {}

func (t AtSign) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type CaretAt struct{}

var _ Token = (*CaretAt)(nil)

func (CaretAt) t() {}

func (t CaretAt) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type PGSquareRoot struct{}

var _ Token = (*PGSquareRoot)(nil)

func (PGSquareRoot) t() {}

func (t PGSquareRoot) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type PGCubeRoot struct{}

var _ Token = (*PGCubeRoot)(nil)

func (PGCubeRoot) t() {}

func (t PGCubeRoot) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Placeholder struct{}

var _ Token = (*Placeholder)(nil)

func (Placeholder) t() {}

func (t Placeholder) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Arrow struct{}

var _ Token = (*Arrow)(nil)

func (Arrow) t() {}

func (t Arrow) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type LongArrow struct{}

var _ Token = (*LongArrow)(nil)

func (LongArrow) t() {}

func (t LongArrow) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type HashArrow struct{}

var _ Token = (*HashArrow)(nil)

func (HashArrow) t() {}

func (t HashArrow) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type HashLongArrow struct{}

var _ Token = (*HashLongArrow)(nil)

func (HashLongArrow) t() {}

func (t HashLongArrow) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type AtArrow struct{}

var _ Token = (*AtArrow)(nil)

func (AtArrow) t() {}

func (t AtArrow) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type ArrowAt struct{}

var _ Token = (*ArrowAt)(nil)

func (ArrowAt) t() {}

func (t ArrowAt) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type HashMinus struct{}

var _ Token = (*HashMinus)(nil)

func (HashMinus) t() {}

func (t HashMinus) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type AtQuestion struct{}

var _ Token = (*AtQuestion)(nil)

func (AtQuestion) t() {}

func (t AtQuestion) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type AtAt struct{}

var _ Token = (*AtAt)(nil)

func (AtAt) t() {}

func (t AtAt) Format(f fmt.State, verb rune) {
	f.Write([]byte(""))
}

type Word struct {
	// The value of the token, without the enclosing quotes, and with the
	// escape sequences (if any) processed (TODO: escapes are not handled)
	Value string
	// An identifier can be "quoted" (&lt;delimited identifier> in ANSI parlance).
	// The standard and most implementations allow using double quotes for this,
	// but some implementations support other quoting styles as well (e.g. \[MS SQL])
	QUoteStyle *rune
	// If the word was not quoted and it matched one of the known keywords,
	// this will have one of the values from dialect::keywords, otherwise empty
	Keyword keyword.KeyWord
}

var _ Token = (*Word)(nil)

func (Word) t() {}

func (w Word) Format(f fmt.State, verb rune) {
	if w.QUoteStyle != nil {
		fmt.Fprintf(f, "%c%s%c", *w.QUoteStyle, w.Value, *w.QUoteStyle)
	} else {
		f.Write([]byte(w.Value))
	}
}

type Whitespace interface {
	Token
	ws()
}

type Space struct{}

var _ Whitespace = (*Space)(nil)

func (Space) ws() {}
func (Space) t()  {}

func (Space) Format(f fmt.State, ver rune) {
	f.Write([]byte(" "))
}

type NewLine struct{}

var _ Whitespace = (*NewLine)(nil)

func (NewLine) ws() {}
func (NewLine) t()  {}

func (NewLine) Format(f fmt.State, ver rune) {
	f.Write([]byte("\n"))
}

type Tab struct{}

var _ Whitespace = (*Tab)(nil)

func (Tab) ws() {}
func (Tab) t()  {}

func (Tab) Format(f fmt.State, ver rune) {
	f.Write([]byte("\t"))
}

type SingleLineComment struct {
	Comment string
	Prefix  string
}

var _ Whitespace = (*SingleLineComment)(nil)

func (SingleLineComment) ws() {}
func (SingleLineComment) t()  {}

func (s SingleLineComment) Format(f fmt.State, ver rune) {
	fmt.Fprintf(f, "%s%s", s.Prefix, s.Comment)
}

type MultiLineComment struct {
	Comment string
}

var _ Whitespace = (*MultiLineComment)(nil)

func (MultiLineComment) ws() {}
func (MultiLineComment) t()  {}

func (m MultiLineComment) Format(f fmt.State, ver rune) {
	fmt.Fprintf(f, "/*%s*/", m.Comment)
}

type Location struct {
	Line   uint64
	Column uint64
}

type TokenWithLocation struct {
	Token    Token
	Location *Location
}

type Error struct {
	Message  string
	Location *Location
}
