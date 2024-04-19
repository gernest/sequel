package dialect

import "github.com/gernest/sequel/ascii"

type BigQuery struct {
	Base
}

func (BigQuery) IsDelimitedIdentStart(ch rune) bool {
	return ch == '`'
}

func (BigQuery) ID() string {
	return "bigquery"
}

func (BigQuery) IsIdentStart(ch rune) bool {
	return ascii.IsIdentStart(ch)
}

func (BigQuery) IsIdentPart(ch rune) bool {
	return ascii.IsIdentPart(ch)
}
