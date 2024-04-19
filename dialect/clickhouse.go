package dialect

import "github.com/gernest/sequel/ascii"

type ClickHouse struct {
	Base
}

func (ClickHouse) ID() string {
	return "clickhouse"
}

func (ClickHouse) IsIdentStart(ch rune) bool {
	return ascii.IsIdentPart(ch)
}

func (ClickHouse) IsIdentPart(ch rune) bool {
	return ascii.IsIdentStart(ch)
}
