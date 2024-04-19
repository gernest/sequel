package password

import "github.com/gernest/sequel/ast/expr"

type Password interface {
	passwd()
}

type Pass struct {
	Expr expr.Expr
}

var _ Password = (*Pass)(nil)

func (Pass) passwd() {}

type Null struct{}

var _ Password = (*Null)(nil)

func (Null) passwd() {}
