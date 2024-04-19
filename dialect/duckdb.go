package dialect

import (
	"unicode"

	"github.com/gernest/sequel/ascii"
)

type DuckDB struct {
	Base
}

func (DuckDB) ID() string {
	return "duckdb"
}

func (DuckDB) IsIdentStart(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func (DuckDB) IsIdentPart(ch rune) bool {
	return unicode.IsLetter(ch) || ascii.Digit(ch) || ch == '$' || ch == '_'
}

func (DuckDB) SupportsFilterDuringAggregation() bool {
	return true
}
func (DuckDB) SupportsGroupByExpr() bool {
	return true
}
func (DuckDB) SupportsNamedFnArgsWithEqOp() bool {
	return true
}
