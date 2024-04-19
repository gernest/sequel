package dialect

import "github.com/gernest/sequel/ascii"

type Snowflake struct {
	Base
}

func (Snowflake) ID() string {
	return "snowflake"
}

func (Snowflake) IsIdentStart(ch rune) bool {
	return ascii.IsIdentStart(ch) || ch == '_'
}

func (Snowflake) IsIdentPart(ch rune) bool {
	return ascii.IsIdentPart(ch) || ch == '$'
}

func (Snowflake) SupportsWithinAfterArrayAggregation() bool {
	return true
}
