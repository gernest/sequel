package ast

import (
	"fmt"
	"strings"

	"github.com/gernest/sequel/ast/common"
)

type Type interface {
	fmt.Formatter
	dt()
}

// Fixed-length character type e.g. CHARACTER(10)
type Character struct {
	Size CharacterLength
}

var _ Type = (*Character)(nil)

func (c Character) dt() {}

func (c Character) Format(f fmt.State, verb rune) {
	fmtStringType(f, "CHARACTER", c.Size)
}

// Fixed-length char type e.g. CHAR(10)
type Char struct {
	Size CharacterLength
}

var _ Type = (*Char)(nil)

func (d Char) dt() {}
func (d Char) Format(f fmt.State, verb rune) {
	fmtStringType(f, "CHAR", d.Size)
}

// Character varying type e.g. CHARACTER VARYING(10)
type CharacterVarying struct {
	Size CharacterLength
}

var _ Type = (*CharacterVarying)(nil)

func (d CharacterVarying) dt() {}
func (d CharacterVarying) Format(f fmt.State, verb rune) {
	fmtStringType(f, "CHARACTER VARYING", d.Size)
}

// Char varying type e.g. CHAR VARYING(10)
type CharVarying struct {
	Size CharacterLength
}

var _ Type = (*CharVarying)(nil)

func (d CharVarying) dt() {}
func (d CharVarying) Format(f fmt.State, verb rune) {
	fmtStringType(f, "CHAR VARYING", d.Size)
}

// Variable-length character type e.g. VARCHAR(10)
type Varchar struct {
	Size CharacterLength
}

var _ Type = (*Varchar)(nil)

func (d Varchar) dt() {}
func (d Varchar) Format(f fmt.State, verb rune) {
	fmtStringType(f, "VARCHAR", d.Size)
}

// Variable-length character type e.g. NVARCHAR(10)
type Nvarchar struct {
	Size *uint64
}

var _ Type = (*Nvarchar)(nil)

func (d Nvarchar) dt() {}

func (d Nvarchar) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "NVARCHAR", d.Size, false)
}

// Uuid type
type UUID struct{}

var _ Type = (*UUID)(nil)

func (d UUID) dt() {}
func (d UUID) Format(f fmt.State, verb rune) {
	f.Write([]byte("UUID"))
}

// Large character object with optional length e.g. CHARACTER LARGE OBJECT, CHARACTER LARGE OBJECT(1000), [standard]
//
// [standard]: https://jakewheat.github.io/sql-overview/sql-2016-foundation-grammar.html#character-large-object-type
type CharacterLargeObject struct {
	Size *uint64
}

var _ Type = (*CharacterLargeObject)(nil)

func (d CharacterLargeObject) dt() {}

func (d CharacterLargeObject) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "CHARACTER LARGE OBJECT", d.Size, false)
}

// Large character object with optional length e.g. CHAR LARGE OBJECT, CHAR LARGE OBJECT(1000), [standard]
//
// [standard]: https://jakewheat.github.io/sql-overview/sql-2016-foundation-grammar.html#character-large-object-type
type CharLargeObject struct {
	Size *uint64
}

var _ Type = (*CharLargeObject)(nil)

func (d CharLargeObject) dt() {}

func (d CharLargeObject) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "CHAR LARGE OBJECT", d.Size, false)
}

// Large character object with optional length e.g. CLOB, CLOB(1000), [standard]
//
// [standard]: https://jakewheat.github.io/sql-overview/sql-2016-foundation-grammar.html#character-large-object-type
//
// [Oracle]: https://docs.oracle.com/javadb/10.10.1.2/ref/rrefclob.html
type Clob struct {
	Size *uint64
}

var _ Type = (*Clob)(nil)

func (d Clob) dt() {}

func (d Clob) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "CLOB", d.Size, false)
}

// Fixed-length binary type with optional length e.g.  [standard], [MS SQL Server]
//
// [standard]: https://jakewheat.github.io/sql-overview/sql-2016-foundation-grammar.html#binary-string-type
// [MS SQL Server]: https://learn.microsoft.com/pt-br/sql/t-sql/data-types/binary-and-varbinary-transact-sql?view=sql-server-ver16
type Binary struct {
	Size *uint64
}

var _ Type = (*Binary)(nil)

func (d Binary) dt() {}

func (d Binary) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "BINARY", d.Size, false)
}

type Varbinary struct {
	Size *uint64
}

var _ Type = (*Varbinary)(nil)

func (d Varbinary) dt() {}

func (d Varbinary) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "VARBINARY", d.Size, false)
}

type Blob struct {
	Size *uint64
}

var _ Type = (*Blob)(nil)

func (d Blob) dt() {}

func (d Blob) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "BLOB", d.Size, false)
}

type Bytes struct {
	Size *uint64
}

var _ Type = (*Bytes)(nil)

func (d Bytes) dt() {}

func (d Bytes) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "BYTES", d.Size, false)
}

type Numeric struct {
	Info ExactNumberInfo
}

var _ Type = (*Numeric)(nil)

func (d Numeric) dt() {}

func (d Numeric) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "NUMERIC%v", d.Info)
}

type Decimal struct {
	Info ExactNumberInfo
}

var _ Type = (*Decimal)(nil)

func (d Decimal) dt() {}

func (d Decimal) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "DECIMAL%v", d.Info)
}

type Dec struct {
	Info ExactNumberInfo
}

var _ Type = (*Dec)(nil)

func (d Dec) dt() {}

func (d Dec) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "DEC%v", d.Info)
}

type BigNumeric struct {
	Info ExactNumberInfo
}

var _ Type = (*BigNumeric)(nil)

func (d BigNumeric) dt() {}

func (d BigNumeric) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "BIGNUMERIC%v", d.Info)
}

type BigDecimal struct {
	Info ExactNumberInfo
}

var _ Type = (*BigDecimal)(nil)

func (d BigDecimal) dt() {}

func (d BigDecimal) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "BIGDECIMAL%v", d.Info)
}

type Float struct {
	Size *uint64
}

var _ Type = (*Float)(nil)

func (d Float) dt() {}

func (d Float) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "FLOAT", d.Size, false)
}

type TinyInt struct {
	ZeroFill *uint64
}

var _ Type = (*TinyInt)(nil)

func (d TinyInt) dt() {}

func (d TinyInt) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "TINYINT", d.ZeroFill, false)
}

type UnsignedTinyInt struct {
	ZeroFill *uint64
}

var _ Type = (*UnsignedTinyInt)(nil)

func (d UnsignedTinyInt) dt() {}

func (d UnsignedTinyInt) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "TINYINT", d.ZeroFill, false)
}

type Int2 struct {
	ZeroFill *uint64
}

var _ Type = (*Int2)(nil)

func (d Int2) dt() {}

func (d Int2) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "INT2", d.ZeroFill, true)
}

type SmallInt struct {
	ZeroFill *uint64
}

var _ Type = (*SmallInt)(nil)

func (d SmallInt) dt() {}

func (d SmallInt) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "SMALLINT", d.ZeroFill, false)
}

type UnsignedSmallInt struct {
	ZeroFill *uint64
}

var _ Type = (*UnsignedSmallInt)(nil)

func (d UnsignedSmallInt) dt() {}

func (d UnsignedSmallInt) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "SMALLINT", d.ZeroFill, true)
}

type MediumInt struct {
	ZeroFill *uint64
}

var _ Type = (*MediumInt)(nil)

func (d MediumInt) dt() {}

func (d MediumInt) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "MEDIUMINT", d.ZeroFill, false)
}

type UnsignedMediumInt struct {
	ZeroFill *uint64
}

var _ Type = (*UnsignedMediumInt)(nil)

func (d UnsignedMediumInt) dt() {}

func (d UnsignedMediumInt) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "MEDIUMINT", d.ZeroFill, true)
}

type Int struct {
	ZeroFill *uint64
}

var _ Type = (*Int)(nil)

func (d Int) dt() {}

func (d Int) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "INT", d.ZeroFill, false)
}

type UnsignedInt struct {
	ZeroFill *uint64
}

var _ Type = (*Int)(nil)

func (d UnsignedInt) dt() {}

func (d UnsignedInt) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "INT", d.ZeroFill, true)
}

type Int4 struct {
	ZeroFill *uint64
}

var _ Type = (*Int4)(nil)

func (d Int4) dt() {}

func (d Int4) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "INT4", d.ZeroFill, false)
}

type Int64 struct {
}

var _ Type = (*Int64)(nil)

func (d Int64) dt() {}

func (d Int64) Format(f fmt.State, verb rune) {
	f.Write([]byte("INT64"))
}

type UnsignedInt4 struct {
	ZeroFill *uint64
}

var _ Type = (*UnsignedInt4)(nil)

func (d UnsignedInt4) dt() {}

func (d UnsignedInt4) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "INT4", d.ZeroFill, true)
}

type Integer struct {
	ZeroFill *uint64
}

var _ Type = (*Integer)(nil)

func (d Integer) dt() {}

func (d Integer) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "INTEGER", d.ZeroFill, false)
}

type UnsignedInteger struct {
	ZeroFill *uint64
}

var _ Type = (*UnsignedInteger)(nil)

func (d UnsignedInteger) dt() {}

func (d UnsignedInteger) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "INTEGER", d.ZeroFill, true)
}

type BigInt struct {
	ZeroFill *uint64
}

var _ Type = (*BigInt)(nil)

func (d BigInt) dt() {}

func (d BigInt) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "BIGINT", d.ZeroFill, false)
}

type UnsignedBigInt struct {
	ZeroFill *uint64
}

var _ Type = (*UnsignedBigInt)(nil)

func (d UnsignedBigInt) dt() {}

func (d UnsignedBigInt) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "BIGINT", d.ZeroFill, true)
}

type Int8 struct {
	ZeroFill *uint64
}

var _ Type = (*Int8)(nil)

func (d Int8) dt() {}

func (d Int8) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "INT8", d.ZeroFill, false)
}

type UnsignedInt8 struct {
	ZeroFill *uint64
}

var _ Type = (*UnsignedInt8)(nil)

func (d UnsignedInt8) dt() {}

func (d UnsignedInt8) Format(f fmt.State, verb rune) {
	fmtOptionalLength(f, "INT8", d.ZeroFill, true)
}

type Real struct {
}

var _ Type = (*Real)(nil)

func (d Real) dt() {}

func (d Real) Format(f fmt.State, verb rune) {
	f.Write([]byte("REAL"))
}

type Float4 struct {
}

var _ Type = (*Float4)(nil)

func (d Float4) dt() {}

func (d Float4) Format(f fmt.State, verb rune) {
	f.Write([]byte("FLOAT4"))
}

type Float64 struct {
}

var _ Type = (*Float64)(nil)

func (d Float64) dt() {}

func (d Float64) Format(f fmt.State, verb rune) {
	f.Write([]byte("FLOAT64"))
}

type Double struct {
}

var _ Type = (*Double)(nil)

func (d Double) dt() {}

func (d Double) Format(f fmt.State, verb rune) {
	f.Write([]byte("DOUBLE"))
}

type Float8 struct {
}

var _ Type = (*Float8)(nil)

func (d Float8) dt() {}

func (d Float8) Format(f fmt.State, verb rune) {
	f.Write([]byte("FLOAT8"))
}

type DoublePrecision struct {
}

var _ Type = (*DoublePrecision)(nil)

func (d DoublePrecision) dt() {}

func (d DoublePrecision) Format(f fmt.State, verb rune) {
	f.Write([]byte("DOUBLE PRECISION"))
}

type Bool struct {
}

var _ Type = (*Bool)(nil)

func (d Bool) dt() {}

func (d Bool) Format(f fmt.State, verb rune) {
	f.Write([]byte("BOOL"))
}

type Boolean struct {
}

var _ Type = (*Boolean)(nil)

func (d Boolean) dt() {}

func (d Boolean) Format(f fmt.State, verb rune) {
	f.Write([]byte("BOOLEAN"))
}

type Date struct {
}

var _ Type = (*Date)(nil)

func (d Date) dt() {}

func (d Date) Format(f fmt.State, verb rune) {
	f.Write([]byte("DATE"))
}

type Time struct {
	Precision    uint64
	TimezoneInfo TimezoneInfo
}

var _ Type = (*Time)(nil)

func (d Time) dt() {}

func (d Time) Format(f fmt.State, verb rune) {
	fmtDateTimePrecisionTz(f, "TIME", d.Precision, d.TimezoneInfo)
}

type Timestamp struct {
	Precision    uint64
	TimezoneInfo TimezoneInfo
}

var _ Type = (*Timestamp)(nil)

func (d Timestamp) dt() {}

func (d Timestamp) Format(f fmt.State, verb rune) {
	fmtDateTimePrecisionTz(f, "TIMESTAMP", d.Precision, d.TimezoneInfo)
}

type Interval struct{}

var _ Type = (*Interval)(nil)

func (d Interval) dt() {}

func (d Interval) Format(f fmt.State, verb rune) {
	f.Write([]byte("INTERVAL"))
}

type JSON struct{}

var _ Type = (*JSON)(nil)

func (d JSON) dt() {}

func (d JSON) Format(f fmt.State, verb rune) {
	f.Write([]byte("JSON"))
}

type JSONB struct{}

var _ Type = (*JSONB)(nil)

func (d JSONB) dt() {}

func (d JSONB) Format(f fmt.State, verb rune) {
	f.Write([]byte("JSONB"))
}

type Regclass struct{}

var _ Type = (*Regclass)(nil)

func (d Regclass) dt() {}

func (d Regclass) Format(f fmt.State, verb rune) {
	f.Write([]byte("REGCLASS"))
}

type Text struct{}

var _ Type = (*Text)(nil)

func (d Text) dt() {}

func (d Text) Format(f fmt.State, verb rune) {
	f.Write([]byte("TEXT"))
}

type Bytea struct{}

var _ Type = (*Bytea)(nil)

func (d Bytea) dt() {}

func (d Bytea) Format(f fmt.State, verb rune) {
	f.Write([]byte("BYTEA"))
}

type Custom struct {
	Type      common.ObjectName
	Modifiers []string
}

var _ Type = (*Custom)(nil)

func (d Custom) dt() {}

func (d Custom) Format(f fmt.State, verb rune) {
	if len(d.Modifiers) == 0 {
		fmt.Fprintf(f, "%v", d.Type)
	} else {
		fmt.Fprintf(f, "%v(%v)", d.Type, strings.Join(d.Modifiers, ", "))
	}
}

type Enum struct {
	Elem []string
}

var _ Type = (*Enum)(nil)

func (d Enum) dt() {}

func (d Enum) Format(f fmt.State, verb rune) {
	f.Write([]byte("ENUM("))
	for i, v := range d.Elem {
		if i != 0 {
			f.Write([]byte(", "))
		}
		fmt.Fprintf(f, "'%v')", common.EscapeSingleQuote(v))
	}
	f.Write([]byte(")"))
}

type Set struct {
	Elem []string
}

var _ Type = (*Set)(nil)

func (d Set) dt() {}

func (d Set) Format(f fmt.State, verb rune) {
	f.Write([]byte("SET("))
	for i, v := range d.Elem {
		if i != 0 {
			f.Write([]byte(", "))
		}
		fmt.Fprintf(f, "'%v')", common.EscapeSingleQuote(v))
	}
	f.Write([]byte(")"))
}

type Struct struct {
	Fields []StructField
}

var _ Type = (*Struct)(nil)

func (d Struct) dt() {}

func (d Struct) Format(f fmt.State, verb rune) {
	if len(d.Fields) > 0 {
		f.Write([]byte("STRUCT<"))
		for i, v := range d.Fields {
			if i != 0 {
				f.Write([]byte(", "))
			}
			fmt.Fprintf(f, "'%v')", v)
		}
		f.Write([]byte(">"))
	} else {
		f.Write([]byte("STRUCT"))

	}

}

type StructField struct {
	Name *common.Ident
	Type Type
}

func (s StructField) Format(f fmt.State, verb rune) {
	if s.Name != nil {
		fmt.Fprintf(f, "%v %v", s.Name, s.Type)
	} else {
		fmt.Fprintf(f, "%v", s.Type)
	}
}

func fmtOptionalLength(f fmt.State, sql_type string, length *uint64, unsigned bool) {
	f.Write([]byte(sql_type))
	if length != nil {
		fmt.Fprintf(f, "(%d)", *length)
	}
	if unsigned {
		f.Write([]byte(" UNSIGNED"))
	}
}

func fmtStringType(f fmt.State, sql_type string, size CharacterLength) {
	fmt.Fprint(f, sql_type)
	if size != nil {
		fmt.Fprintf(f, "%v", size)
	}
}

func fmtDateTimePrecisionTz(f fmt.State, sql_type string, length uint64, zone TimezoneInfo) {
	_, ok := zone.(TimezoneInfoTz)
	if ok {
		fmt.Fprintf(f, "%s%v%d", sql_type, zone, length)
	} else {
		fmt.Fprintf(f, "%s%d%v", sql_type, length, zone)
	}
}

type CharacterLength interface {
	fmt.Formatter
	cl()
}

type CharacterLengthMAX struct{}

var _ CharacterLength = (*CharacterLengthMAX)(nil)

func (CharacterLengthMAX) Format(f fmt.State, verb rune) {
	f.Write([]byte("MAX"))
}

func (CharacterLengthMAX) cl() {}

type CharacterLengthIntegerLength struct {
	Length uint64
	Unit   CharLengthUnits
}

var _ CharacterLength = (*CharacterLengthIntegerLength)(nil)

func (i CharacterLengthIntegerLength) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "%d (%s)", i.Length, i.Unit)
}

func (CharacterLengthIntegerLength) cl() {}

type CharLengthUnits byte

const (
	Characters CharLengthUnits = iota
	Octets
)

func (u CharLengthUnits) String() string {
	switch u {
	case Characters:
		return "CHARACTERS"
	case Octets:
		return "OCTETS"
	default:
		return ""
	}
}

type TimezoneInfo interface {
	fmt.Formatter
	tzi()
}

// No information about time zone. E.g., TIMESTAMP
type TimezoneInfoNone struct{}

var _ TimezoneInfo = (*TimezoneInfoNone)(nil)

func (TimezoneInfoNone) tzi()                          {}
func (TimezoneInfoNone) Format(f fmt.State, verb rune) {}

// Temporal type 'WITH TIME ZONE'. E.g., TIMESTAMP WITH TIME ZONE, [standard], [Oracle]
//
// [standard]: https://jakewheat.github.io/sql-overview/sql-2016-foundation-grammar.html#datetime-type
// [Oracle]: https://docs.oracle.com/en/database/oracle/oracle-database/12.2/nlspg/datetime-data-types-and-time-zone-support.html#GUID-3F1C388E-C651-43D5-ADBC-1A49E5C2CA05
type TimezoneInfoWithTimeZone struct{}

var _ TimezoneInfo = (*TimezoneInfoWithTimeZone)(nil)

func (TimezoneInfoWithTimeZone) tzi() {}
func (TimezoneInfoWithTimeZone) Format(f fmt.State, verb rune) {
	f.Write([]byte(" WITH TIME ZONE"))
}

// Temporal type 'WITHOUT TIME ZONE'. E.g., TIME WITHOUT TIME ZONE, [standard], [Postgresql]
//
// [standard]: https://jakewheat.github.io/sql-overview/sql-2016-foundation-grammar.html#datetime-type
// [Postgresql]: https://www.postgresql.org/docs/current/datatype-datetime.html
type TimezoneInfoWithoutTimeZone struct{}

var _ TimezoneInfo = (*TimezoneInfoWithoutTimeZone)(nil)

func (TimezoneInfoWithoutTimeZone) tzi() {}
func (TimezoneInfoWithoutTimeZone) Format(f fmt.State, verb rune) {
	f.Write([]byte(" WITHOUT TIME ZONE"))
}

// / Postgresql specific `WITH TIME ZONE` formatting, for both TIME and TIMESTAMP. E.g., TIMETZ, [Postgresql]
// /
// / [Postgresql]: https://www.postgresql.org/docs/current/datatype-datetime.html
type TimezoneInfoTz struct{}

var _ TimezoneInfo = (*TimezoneInfoTz)(nil)

func (TimezoneInfoTz) tzi() {}
func (TimezoneInfoTz) Format(f fmt.State, verb rune) {
	// TZ is the only one that is displayed BEFORE the precision, so the datatype display
	// must be aware of that. Check <https://www.postgresql.org/docs/14/datatype-datetime.html>
	// for more information
	f.Write([]byte("TZ"))
}

// Additional information for `NUMERIC`, `DECIMAL`, and `DEC` data types
// following the 2016 [standard].
//
// [standard]: https://jakewheat.github.io/sql-overview/sql-2016-foundation-grammar.html#exact-numeric-type
type ExactNumberInfo interface {
	fmt.Formatter
	ni()
}

// DECIMAL
type ExactNumberInfoNone struct{}

var _ ExactNumberInfo = (*ExactNumberInfoNone)(nil)

func (ExactNumberInfoNone) ni()                           {}
func (ExactNumberInfoNone) Format(f fmt.State, verb rune) {}

// DECIMAL(10)
type ExactNumberInfoPrecision struct {
	Precision uint64
}

var _ ExactNumberInfo = (*ExactNumberInfoPrecision)(nil)

func (ExactNumberInfoPrecision) ni() {}
func (i ExactNumberInfoPrecision) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "(%d)", i.Precision)
}

type ExactNumberInfoPrecisionScale struct {
	Precision, Scale uint64
}

var _ ExactNumberInfo = (*ExactNumberInfoPrecisionScale)(nil)

func (ExactNumberInfoPrecisionScale) ni() {}
func (i ExactNumberInfoPrecisionScale) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "(%d,%d)", i.Precision, i.Scale)
}

// ArrayElemTypeDef represents the data type of the elements in an array (if any) as well as
// the syntax used to declare the array.
//
// For example: Bigquery/Hive use `ARRAY<INT>` whereas snowflake uses ARRAY.
type ArrayElemTypeDef interface{ aetd() }

// ARRAY
type ArrayElemTypeNone struct{}

var _ Type = (*ArrayElemTypeNone)(nil)
var _ ArrayElemTypeDef = (*ArrayElemTypeNone)(nil)

func (d ArrayElemTypeNone) aetd() {}
func (d ArrayElemTypeNone) dt()   {}
func (d ArrayElemTypeNone) Format(f fmt.State, verb rune) {
	f.Write([]byte("ARRAY"))
}

// ARRAY<INT>
type ArrayElemTypeAngleBracket struct {
	Elem Type
}

var _ Type = (*ArrayElemTypeAngleBracket)(nil)
var _ ArrayElemTypeDef = (*ArrayElemTypeAngleBracket)(nil)

func (d ArrayElemTypeAngleBracket) aetd() {}
func (d ArrayElemTypeAngleBracket) dt()   {}
func (d ArrayElemTypeAngleBracket) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "%v[]", d.Elem)
}

type ArrayElemTypeSquareBracket struct {
	Elem Type
}

var _ Type = (*ArrayElemTypeSquareBracket)(nil)
var _ ArrayElemTypeDef = (*ArrayElemTypeSquareBracket)(nil)

func (d ArrayElemTypeSquareBracket) aetd() {}
func (d ArrayElemTypeSquareBracket) dt()   {}
func (d ArrayElemTypeSquareBracket) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "ARRAY<%v>", d.Elem)
}

type ArrayElemTypeT[T any] struct{ Elem T }

func (a ArrayElemTypeT[T]) aetd() {}
