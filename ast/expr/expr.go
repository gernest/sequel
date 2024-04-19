package expr

import "fmt"

type Expr interface {
	fmt.Formatter
	Expr()
}
