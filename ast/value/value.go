package value

import (
	"bytes"
	"fmt"

	"github.com/gernest/sequel/ast/common"
)

type Value interface {
	fmt.Formatter
	v()
}

type Number struct {
	Value string
	Long  bool
}

var _ Value = (*Number)(nil)

func (Number) v() {}

func (v Number) Format(f fmt.State, verb rune) {
	f.Write([]byte(v.Value))
	if v.Long {
		f.Write([]byte("L"))
	}
}

type SingleQuotedString struct {
	Value string
}

var _ Value = (*SingleQuotedString)(nil)

func (SingleQuotedString) v() {}

func (v SingleQuotedString) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "\"%v\"", common.EscapeSingleQuote(v.Value))
}

type DoubleQuotedString struct {
	Value string
}

var _ Value = (*DoubleQuotedString)(nil)

func (DoubleQuotedString) v() {}

func (v DoubleQuotedString) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "\"%v\"", common.EscapeDoubleQuote(v.Value))
}

type DollarQuotedString struct {
	Value string
	Tag   string
}

var _ Value = (*DollarQuotedString)(nil)

func (DollarQuotedString) v() {}

func (v DollarQuotedString) Format(f fmt.State, verb rune) {
	if v.Tag != "" {
		fmt.Fprintf(f, "$%v$%v$%v", v.Tag, v.Value, v.Tag)
	} else {
		fmt.Fprintf(f, "$$%v$$", v.Value)
	}
}

type EscapedStringLiteral struct {
	Value string
}

var _ Value = (*EscapedStringLiteral)(nil)

func (EscapedStringLiteral) v() {}

func (v EscapedStringLiteral) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "E'%v'", EscapeEscapedStringLiteral(v.Value))
}

type NationalStringLiteral struct {
	Value string
}

var _ Value = (*NationalStringLiteral)(nil)

func (NationalStringLiteral) v() {}

func (v NationalStringLiteral) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "N'%v'", v.Value)
}

type HexStringLiteral struct {
	Value string
}

var _ Value = (*HexStringLiteral)(nil)

func (HexStringLiteral) v() {}

func (v HexStringLiteral) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "X'%v'", v.Value)
}

type Boolean struct {
	Value bool
}

var _ Value = (*Boolean)(nil)

func (Boolean) v() {}

func (v Boolean) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "%v", v.Value)
}

type SingleQuotedByteStringLiteral struct {
	Value string
}

var _ Value = (*SingleQuotedByteStringLiteral)(nil)

func (SingleQuotedByteStringLiteral) v() {}

func (v SingleQuotedByteStringLiteral) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "B'%v'", v.Value)
}

type DoubleQuotedByteStringLiteral struct {
	Value string
}

var _ Value = (*DoubleQuotedByteStringLiteral)(nil)

func (DoubleQuotedByteStringLiteral) v() {}

func (v DoubleQuotedByteStringLiteral) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "B\"%v\"", v.Value)
}

type RawStringLiteral struct {
	Value string
}

var _ Value = (*RawStringLiteral)(nil)

func (RawStringLiteral) v() {}

func (v RawStringLiteral) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "R'%v'", v.Value)
}

type Null struct {
}

var _ Value = (*Null)(nil)

func (Null) v() {}

func (Null) Format(f fmt.State, verb rune) {
	f.Write([]byte("NULL"))
}

type Placeholder struct {
	Value string
}

var _ Value = (*Placeholder)(nil)

func (Placeholder) v() {}

func (v Placeholder) Format(f fmt.State, verb rune) {
	f.Write([]byte(v.Value))
}

type UnQuotedString struct {
	Value string
}

var _ Value = (*UnQuotedString)(nil)

func (UnQuotedString) v() {}

func (v UnQuotedString) Format(f fmt.State, verb rune) {
	f.Write([]byte(v.Value))
}

type EscapeEscapedStringLiteral string

func (s EscapeEscapedStringLiteral) Format(f fmt.State, verb rune) {
	var o bytes.Buffer
	o.Grow(len(s))
	for _, b := range s {
		switch b {
		case '\'':
			o.WriteString(`\'`)
		case '\\':
			o.WriteString(`\`)
		case '\n':
			o.WriteString(`\n`)
		case '\r':
			o.WriteString(`\r`)
		case '\t':
			o.WriteString(`\t`)
		default:
			o.WriteRune(b)
		}
	}
	f.Write(o.Bytes())
}

type DateTimeField interface {
	fmt.Formatter
	dtf()
}

type Year struct{}

var _ DateTimeField = (*Year)(nil)

func (Year) dtf() {}

func (Year) Format(f fmt.State, verb rune) {
	f.Write([]byte("YEAR"))
}

type Month struct{}

var _ DateTimeField = (*Month)(nil)

func (Month) dtf() {}

func (Month) Format(f fmt.State, verb rune) {
	f.Write([]byte("MONTH"))
}

type Week struct {
	WeekDay *common.Ident
}

var _ DateTimeField = (*Week)(nil)

func (Week) dtf() {}

func (w Week) Format(f fmt.State, verb rune) {
	f.Write([]byte("WEEK"))
	if w.WeekDay != nil {
		fmt.Fprintf(f, "(%v)", w.WeekDay)
	}
}

type Day struct{}

var _ DateTimeField = (*Day)(nil)

func (Day) dtf() {}

func (Day) Format(f fmt.State, verb rune) {
	f.Write([]byte("DAY"))
}

type DayOfWeek struct{}

var _ DateTimeField = (*DayOfWeek)(nil)

func (DayOfWeek) dtf() {}

func (DayOfWeek) Format(f fmt.State, verb rune) {
	f.Write([]byte("DAYOFWEEK"))
}

type DayOfYear struct{}

var _ DateTimeField = (*DayOfYear)(nil)

func (DayOfYear) dtf() {}

func (DayOfYear) Format(f fmt.State, verb rune) {
	f.Write([]byte("DAYOFYEAR"))
}

type Date struct{}

var _ DateTimeField = (*Date)(nil)

func (Date) dtf() {}

func (Date) Format(f fmt.State, verb rune) {
	f.Write([]byte("DATE"))
}

type Datetime struct{}

var _ DateTimeField = (*Datetime)(nil)

func (Datetime) dtf() {}

func (Datetime) Format(f fmt.State, verb rune) {
	f.Write([]byte("DATETIME"))
}

type Hour struct{}

var _ DateTimeField = (*Hour)(nil)

func (Hour) dtf() {}

func (Hour) Format(f fmt.State, verb rune) {
	f.Write([]byte("HOUR"))
}

type Minute struct{}

var _ DateTimeField = (*Minute)(nil)

func (Minute) dtf() {}

func (Minute) Format(f fmt.State, verb rune) {
	f.Write([]byte("MINUTE"))
}

type Second struct{}

var _ DateTimeField = (*Second)(nil)

func (Second) dtf() {}

func (Second) Format(f fmt.State, verb rune) {
	f.Write([]byte("SECOND"))
}

type Century struct{}

var _ DateTimeField = (*Century)(nil)

func (Century) dtf() {}

func (Century) Format(f fmt.State, verb rune) {
	f.Write([]byte("CENTURY"))
}

type Decade struct{}

var _ DateTimeField = (*Decade)(nil)

func (Decade) dtf() {}

func (Decade) Format(f fmt.State, verb rune) {
	f.Write([]byte("DECADE"))
}

type Dow struct{}

var _ DateTimeField = (*Dow)(nil)

func (Dow) dtf() {}

func (Dow) Format(f fmt.State, verb rune) {
	f.Write([]byte("DOW"))
}

type Doy struct{}

var _ DateTimeField = (*Doy)(nil)

func (Doy) dtf() {}

func (Doy) Format(f fmt.State, verb rune) {
	f.Write([]byte("DOY"))
}

type Epoch struct{}

var _ DateTimeField = (*Epoch)(nil)

func (Epoch) dtf() {}

func (Epoch) Format(f fmt.State, verb rune) {
	f.Write([]byte("EPOCH"))
}

type Isodow struct{}

var _ DateTimeField = (*Isodow)(nil)

func (Isodow) dtf() {}

func (Isodow) Format(f fmt.State, verb rune) {
	f.Write([]byte("ISODOW"))
}

type IsoWeek struct{}

var _ DateTimeField = (*IsoWeek)(nil)

func (IsoWeek) dtf() {}

func (IsoWeek) Format(f fmt.State, verb rune) {
	f.Write([]byte("ISOWEEK"))
}

type Isoyear struct{}

var _ DateTimeField = (*Isoyear)(nil)

func (Isoyear) dtf() {}

func (Isoyear) Format(f fmt.State, verb rune) {
	f.Write([]byte("ISOYEAR"))
}

type Julian struct{}

var _ DateTimeField = (*Julian)(nil)

func (Julian) dtf() {}

func (Julian) Format(f fmt.State, verb rune) {
	f.Write([]byte("JULIAN"))
}

type Microsecond struct{}

var _ DateTimeField = (*Microsecond)(nil)

func (Microsecond) dtf() {}

func (Microsecond) Format(f fmt.State, verb rune) {
	f.Write([]byte("MICROSECOND"))
}

type Microseconds struct{}

var _ DateTimeField = (*Microseconds)(nil)

func (Microseconds) dtf() {}

func (Microseconds) Format(f fmt.State, verb rune) {
	f.Write([]byte("MICROSECONDS"))
}

type Millenium struct{}

var _ DateTimeField = (*Millenium)(nil)

func (Millenium) dtf() {}

func (Millenium) Format(f fmt.State, verb rune) {
	f.Write([]byte("MILLENIUM"))
}

type Millennium struct{}

var _ DateTimeField = (*Millennium)(nil)

func (Millennium) dtf() {}

func (Millennium) Format(f fmt.State, verb rune) {
	f.Write([]byte("MILLENNIUM"))
}

type Millisecond struct{}

var _ DateTimeField = (*Millisecond)(nil)

func (Millisecond) dtf() {}

func (Millisecond) Format(f fmt.State, verb rune) {
	f.Write([]byte("MILLISECOND"))
}

type Milliseconds struct{}

var _ DateTimeField = (*Milliseconds)(nil)

func (Milliseconds) dtf() {}

func (Milliseconds) Format(f fmt.State, verb rune) {
	f.Write([]byte("MILLISECONDS"))
}

type Nanosecond struct{}

var _ DateTimeField = (*Nanosecond)(nil)

func (Nanosecond) dtf() {}

func (Nanosecond) Format(f fmt.State, verb rune) {
	f.Write([]byte("NANOSECOND"))
}

type Nanoseconds struct{}

var _ DateTimeField = (*Nanoseconds)(nil)

func (Nanoseconds) dtf() {}

func (Nanoseconds) Format(f fmt.State, verb rune) {
	f.Write([]byte("NANOSECONDS"))
}

type Quarter struct{}

var _ DateTimeField = (*Quarter)(nil)

func (Quarter) dtf() {}

func (Quarter) Format(f fmt.State, verb rune) {
	f.Write([]byte("QUARTER"))
}

type Time struct{}

var _ DateTimeField = (*Time)(nil)

func (Time) dtf() {}

func (Time) Format(f fmt.State, verb rune) {
	f.Write([]byte("TIME"))
}

type Timezone struct{}

var _ DateTimeField = (*Timezone)(nil)

func (Timezone) dtf() {}

func (Timezone) Format(f fmt.State, verb rune) {
	f.Write([]byte("TIMEZONE"))
}

type TimezoneAbbr struct{}

var _ DateTimeField = (*TimezoneAbbr)(nil)

func (TimezoneAbbr) dtf() {}

func (TimezoneAbbr) Format(f fmt.State, verb rune) {
	f.Write([]byte("TIMEZONE_ABBR"))
}

type TimezoneHour struct{}

var _ DateTimeField = (*TimezoneHour)(nil)

func (TimezoneHour) dtf() {}

func (TimezoneHour) Format(f fmt.State, verb rune) {
	f.Write([]byte("TIMEZONE_HOUR"))
}

type TimezoneMinute struct{}

var _ DateTimeField = (*TimezoneMinute)(nil)

func (TimezoneMinute) dtf() {}

func (TimezoneMinute) Format(f fmt.State, verb rune) {
	f.Write([]byte("TIMEZONE_MINUTE"))
}

type TimezoneRegion struct{}

var _ DateTimeField = (*TimezoneRegion)(nil)

func (TimezoneRegion) dtf() {}

func (TimezoneRegion) Format(f fmt.State, verb rune) {
	f.Write([]byte("TIMEZONE_REGION"))
}

type NoDateTime struct{}

var _ DateTimeField = (*NoDateTime)(nil)

func (NoDateTime) dtf() {}

func (NoDateTime) Format(f fmt.State, verb rune) {
	f.Write([]byte("NODATETIME"))
}

type Custom struct {
	Value string
}

var _ DateTimeField = (*Custom)(nil)

func (Custom) dtf() {}

func (c Custom) Format(f fmt.State, verb rune) {
	f.Write([]byte(c.Value))
}
