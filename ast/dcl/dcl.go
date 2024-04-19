package dcl

import (
	"fmt"

	"github.com/gernest/sequel/ast/common"
	"github.com/gernest/sequel/ast/expr"
	"github.com/gernest/sequel/ast/password"
)

type RoleOption interface {
	fmt.Formatter
	ro()
}

type BypassRLS struct {
	Value bool
}

var _ RoleOption = (*BypassRLS)(nil)

func (BypassRLS) ro() {}
func (v BypassRLS) Format(f fmt.State, verb rune) {
	if v.Value {
		f.Write([]byte("BYPASSRLS"))
	} else {
		f.Write([]byte("NOBYPASSRLS"))
	}
}

type ConnectionLimit struct {
	Value expr.Expr
}

var _ RoleOption = (*ConnectionLimit)(nil)

func (ConnectionLimit) ro() {}
func (v ConnectionLimit) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "CONNECTION LIMIT %v", v.Value)
}

type CreateDB struct {
	Value bool
}

var _ RoleOption = (*CreateDB)(nil)

func (CreateDB) ro() {}
func (v CreateDB) Format(f fmt.State, verb rune) {
	if v.Value {
		f.Write([]byte("CREATEDB"))
	} else {
		f.Write([]byte("NOCREATEDB"))
	}
}

type CreateRole struct {
	Value bool
}

var _ RoleOption = (*CreateRole)(nil)

func (CreateRole) ro() {}
func (v CreateRole) Format(f fmt.State, verb rune) {
	if v.Value {
		f.Write([]byte("CREATEROLE"))
	} else {
		f.Write([]byte("NOCREATEROLE"))
	}
}

type Inherit struct {
	Value bool
}

var _ RoleOption = (*Inherit)(nil)

func (Inherit) ro() {}
func (v Inherit) Format(f fmt.State, verb rune) {
	if v.Value {
		f.Write([]byte("INHERIT"))
	} else {
		f.Write([]byte("NOINHERIT"))
	}
}

type Login struct {
	Value bool
}

var _ RoleOption = (*Login)(nil)

func (Login) ro() {}
func (v Login) Format(f fmt.State, verb rune) {
	if v.Value {
		f.Write([]byte("LOGIN"))
	} else {
		f.Write([]byte("NOLOGIN"))
	}
}

type Password struct {
	Value password.Password
}

var _ RoleOption = (*Password)(nil)

func (Password) ro() {}
func (v Password) Format(f fmt.State, verb rune) {
	switch e := v.Value.(type) {
	case password.Pass:
		fmt.Fprintf(f, "PASSWORD %v", e.Expr)
	case password.Null:
		f.Write([]byte("PASSWORD NULL"))
	}
}

type Replication struct {
	Value bool
}

var _ RoleOption = (*Replication)(nil)

func (Replication) ro() {}
func (v Replication) Format(f fmt.State, verb rune) {
	if v.Value {
		f.Write([]byte("REPLICATION"))
	} else {
		f.Write([]byte("NOREPLICATION"))
	}
}

type SuperUser struct {
	Value bool
}

var _ RoleOption = (*SuperUser)(nil)

func (SuperUser) ro() {}
func (v SuperUser) Format(f fmt.State, verb rune) {
	if v.Value {
		f.Write([]byte("SUPERUSER"))
	} else {
		f.Write([]byte("NOSUPERUSER"))
	}
}

type ValidUntil struct {
	Value expr.Expr
}

var _ RoleOption = (*ValidUntil)(nil)

func (ValidUntil) ro() {}
func (v ValidUntil) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "VALID UNTIL %v", v.Value)
}

type ResetConfig interface {
	reset()
}

type ResetAll struct{}

func (ResetAll) reset() {}

type ResetConfigName struct {
	Value common.ObjectName
}

func (ResetConfigName) reset() {}

type SetConfigValue interface {
	scv()
}

type Default struct{}

var _ SetConfigValue = (*Default)(nil)

func (Default) scv() {}

type FromCurrent struct{}

var _ SetConfigValue = (*FromCurrent)(nil)

func (FromCurrent) scv() {}

type Value struct {
	Value expr.Expr
}

var _ SetConfigValue = (*Value)(nil)

func (Value) scv() {}

type AlterRoleOperation interface {
	fmt.Formatter
	ar()
}

type RenameRole struct {
	Name common.Ident
}

var _ AlterRoleOperation = (*RenameRole)(nil)

func (RenameRole) ar() {}

func (r RenameRole) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "RENAME TO %v", r.Name)
}

type AddMember struct {
	Name common.Ident
}

var _ AlterRoleOperation = (*AddMember)(nil)

func (AddMember) ar() {}

func (r AddMember) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "ADD MEMBER %v", r.Name)
}

type DropMember struct {
	Name common.Ident
}

var _ AlterRoleOperation = (*DropMember)(nil)

func (DropMember) ar() {}

func (r DropMember) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "DROP MEMBER %v", r.Name)
}

type WithOptions struct {
	Options []RoleOption
}

var _ AlterRoleOperation = (*WithOptions)(nil)

func (WithOptions) ar() {}

func (r WithOptions) Format(f fmt.State, verb rune) {
	f.Write([]byte("WITH "))
	common.Separated(f, []byte("."), r.Options...)
}

type Set struct {
	Name       common.ObjectName
	Value      SetConfigValue
	InDatabase common.ObjectName
}

var _ AlterRoleOperation = (*Set)(nil)

func (Set) ar() {}

func (r Set) Format(f fmt.State, verb rune) {
	if len(r.InDatabase) > 0 {
		fmt.Fprintf(f, "IN DATABASE %v", r.InDatabase)
	}
	switch e := r.Value.(type) {
	case Default:
		fmt.Fprintf(f, "SET %v TO DEFAULT", r.Name)
	case FromCurrent:
		fmt.Fprintf(f, "SET %v FROM CURRENT", r.Name)
	case Value:
		fmt.Fprintf(f, "SET %v TO %v", r.Name, e.Value)
	}
}

type Reset struct {
	Name       ResetConfig
	InDatabase common.ObjectName
}

var _ AlterRoleOperation = (*Reset)(nil)

func (Reset) ar() {}

func (r Reset) Format(f fmt.State, verb rune) {
	if len(r.InDatabase) > 0 {
		fmt.Fprintf(f, "IN DATABASE %v", r.InDatabase)
	}
	switch e := r.Name.(type) {
	case ResetAll:
		f.Write([]byte("RESET ALL"))
	case ResetConfigName:
		fmt.Fprintf(f, "RESET %v", e.Value)
	}
}
