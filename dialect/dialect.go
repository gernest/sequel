package dialect

import (
	"github.com/gernest/sequel/stream"
)

type Dialect interface {
	ID() string
	IsDelimitedIdentStart(ch rune) bool
	IdentQuoteStyle(ident string) rune
	IsProperIdentInsideQuote(r stream.Runes) bool
	IsIdentStart(ch rune) bool
	IsIdentPart(ch rune) bool
	SupportsFilterDuringAggregation() bool
	SupportsWithinAfterArrayAggregation() bool
	SupportsGroupByExpr() bool
	SupportsInEmptyList() bool
	SupportsNamedFnArgsWithEqOp() bool
	SupportsStartTxnModifier() bool
	ConvertTypeBeforeValues() bool
}
type Base struct{}

var _ Dialect = (*Base)(nil)

func (Base) ID() string                                   { return "" }
func (Base) IsDelimitedIdentStart(ch rune) bool           { return ch == '"' || ch == '`' }
func (Base) IdentQuoteStyle(ident string) rune            { return 0 }
func (Base) IsProperIdentInsideQuote(r stream.Runes) bool { return true }
func (Base) IsIdentStart(ch rune) bool                    { return false }
func (Base) IsIdentPart(ch rune) bool                     { return false }
func (Base) SupportsFilterDuringAggregation() bool        { return false }
func (Base) SupportsWithinAfterArrayAggregation() bool    { return false }
func (Base) SupportsGroupByExpr() bool                    { return false }
func (Base) SupportsInEmptyList() bool                    { return false }
func (Base) SupportsNamedFnArgsWithEqOp() bool            { return false }
func (Base) SupportsStartTxnModifier() bool               { return false }
func (Base) ConvertTypeBeforeValues() bool                { return false }
