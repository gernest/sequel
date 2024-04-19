package ops

import (
	"fmt"
	"strings"
)

type Op interface {
	fmt.Formatter
	op()
}
type UnaryOperator interface {
	Op
	uop()
}

// +9
type UnaryPlus struct{}

var _ UnaryOperator = (*UnaryPlus)(nil)

func (UnaryPlus) uop()                          {}
func (UnaryPlus) op()                           {}
func (UnaryPlus) Format(f fmt.State, verb rune) { f.Write([]byte("+")) }

// -9
type UnaryMinus struct{}

var _ UnaryOperator = (*UnaryMinus)(nil)

func (UnaryMinus) op()                           {}
func (UnaryMinus) uop()                          {}
func (UnaryMinus) Format(f fmt.State, verb rune) { f.Write([]byte("-")) }

type Not struct{}

var _ UnaryOperator = (*Not)(nil)

func (Not) op()                           {}
func (Not) uop()                          {}
func (Not) Format(f fmt.State, verb rune) { f.Write([]byte("NOT")) }

type PGBitwiseNot struct{}

var _ UnaryOperator = (*PGBitwiseNot)(nil)

func (PGBitwiseNot) op()                           {}
func (PGBitwiseNot) uop()                          {}
func (PGBitwiseNot) Format(f fmt.State, verb rune) { f.Write([]byte("~")) }

type PGSquareRoot struct{}

var _ UnaryOperator = (*PGSquareRoot)(nil)

func (PGSquareRoot) op()                           {}
func (PGSquareRoot) uop()                          {}
func (PGSquareRoot) Format(f fmt.State, verb rune) { f.Write([]byte("|/")) }

type PGCubeRoot struct{}

var _ UnaryOperator = (*PGCubeRoot)(nil)

func (PGCubeRoot) op()                           {}
func (PGCubeRoot) uop()                          {}
func (PGCubeRoot) Format(f fmt.State, verb rune) { f.Write([]byte("||/")) }

type PGPostfixFactorial struct{}

var _ UnaryOperator = (*PGPostfixFactorial)(nil)

func (PGPostfixFactorial) op()                           {}
func (PGPostfixFactorial) uop()                          {}
func (PGPostfixFactorial) Format(f fmt.State, verb rune) { f.Write([]byte("!")) }

type PGPrefixFactorial struct{}

var _ UnaryOperator = (*PGPrefixFactorial)(nil)

func (PGPrefixFactorial) op()                           {}
func (PGPrefixFactorial) uop()                          {}
func (PGPrefixFactorial) Format(f fmt.State, verb rune) { f.Write([]byte("!!")) }

type PGAbs struct{}

var _ UnaryOperator = (*PGAbs)(nil)

func (PGAbs) op()                           {}
func (PGAbs) uop()                          {}
func (PGAbs) Format(f fmt.State, verb rune) { f.Write([]byte("@")) }

type BinaryOperator interface {
	Op
	bop()
}

// a + b
type Plus struct{}

var _ BinaryOperator = (*Plus)(nil)

func (Plus) bop()                          {}
func (Plus) op()                           {}
func (Plus) Format(f fmt.State, verb rune) { f.Write([]byte("+")) }

// a - b
type Minus struct{}

var _ BinaryOperator = (*Minus)(nil)

func (Minus) op()                           {}
func (Minus) bop()                          {}
func (Minus) Format(f fmt.State, verb rune) { f.Write([]byte("-")) }

// a * b
type Multiply struct{}

var _ BinaryOperator = (*Multiply)(nil)

func (Multiply) op()                           {}
func (Multiply) bop()                          {}
func (Multiply) Format(f fmt.State, verb rune) { f.Write([]byte("*")) }

// a * b
type Divide struct{}

var _ BinaryOperator = (*Divide)(nil)

func (Divide) op()                           {}
func (Divide) bop()                          {}
func (Divide) Format(f fmt.State, verb rune) { f.Write([]byte("/")) }

// a % b
type Modulo struct{}

var _ BinaryOperator = (*Modulo)(nil)

func (Modulo) op()                           {}
func (Modulo) bop()                          {}
func (Modulo) Format(f fmt.State, verb rune) { f.Write([]byte("%")) }

// a || b
type StringConcat struct{}

var _ BinaryOperator = (*StringConcat)(nil)

func (StringConcat) op()                           {}
func (StringConcat) bop()                          {}
func (StringConcat) Format(f fmt.State, verb rune) { f.Write([]byte("||")) }

// a > b
type Gt struct{}

var _ BinaryOperator = (*Gt)(nil)

func (Gt) op()                           {}
func (Gt) bop()                          {}
func (Gt) Format(f fmt.State, verb rune) { f.Write([]byte(">")) }

// a < b
type Lt struct{}

var _ BinaryOperator = (*Lt)(nil)

func (Lt) op()                           {}
func (Lt) bop()                          {}
func (Lt) Format(f fmt.State, verb rune) { f.Write([]byte("<")) }

// a >= b
type GtEq struct{}

var _ BinaryOperator = (*GtEq)(nil)

func (GtEq) op()                           {}
func (GtEq) bop()                          {}
func (GtEq) Format(f fmt.State, verb rune) { f.Write([]byte(">=")) }

// a <= b
type LtEq struct{}

var _ BinaryOperator = (*LtEq)(nil)

func (LtEq) op()                           {}
func (LtEq) bop()                          {}
func (LtEq) Format(f fmt.State, verb rune) { f.Write([]byte("<=")) }

// a <=> b
type Spaceship struct{}

var _ BinaryOperator = (*Spaceship)(nil)

func (Spaceship) op()                           {}
func (Spaceship) bop()                          {}
func (Spaceship) Format(f fmt.State, verb rune) { f.Write([]byte("<=>")) }

// a = b
type Eq struct{}

var _ BinaryOperator = (*Eq)(nil)

func (Eq) op()                           {}
func (Eq) bop()                          {}
func (Eq) Format(f fmt.State, verb rune) { f.Write([]byte("=")) }

// a <> b
type NotEq struct{}

var _ BinaryOperator = (*NotEq)(nil)

func (NotEq) op()                           {}
func (NotEq) bop()                          {}
func (NotEq) Format(f fmt.State, verb rune) { f.Write([]byte("<>")) }

// a AND b
type And struct{}

var _ BinaryOperator = (*And)(nil)

func (And) op()                           {}
func (And) bop()                          {}
func (And) Format(f fmt.State, verb rune) { f.Write([]byte("AND")) }

// a OR b
type Or struct{}

var _ BinaryOperator = (*Or)(nil)

func (Or) op()                           {}
func (Or) bop()                          {}
func (Or) Format(f fmt.State, verb rune) { f.Write([]byte("OR")) }

// a XOR b
type Xor struct{}

var _ BinaryOperator = (*Xor)(nil)

func (Xor) op()                           {}
func (Xor) bop()                          {}
func (Xor) Format(f fmt.State, verb rune) { f.Write([]byte("XOR")) }

// a | b
type BitwiseOr struct{}

var _ BinaryOperator = (*BitwiseOr)(nil)

func (BitwiseOr) op()                           {}
func (BitwiseOr) bop()                          {}
func (BitwiseOr) Format(f fmt.State, verb rune) { f.Write([]byte("|")) }

// a & b
type BitwiseAnd struct{}

var _ BinaryOperator = (*BitwiseAnd)(nil)

func (BitwiseAnd) op()                           {}
func (BitwiseAnd) bop()                          {}
func (BitwiseAnd) Format(f fmt.State, verb rune) { f.Write([]byte("&")) }

// a ^ b
type BitwiseXor struct{}

var _ BinaryOperator = (*BitwiseXor)(nil)

func (BitwiseXor) op()                           {}
func (BitwiseXor) bop()                          {}
func (BitwiseXor) Format(f fmt.State, verb rune) { f.Write([]byte("^")) }

// Integer division operator `//` in DuckDB
type DuckIntegerDivide struct{}

var _ BinaryOperator = (*DuckIntegerDivide)(nil)

func (DuckIntegerDivide) op()                           {}
func (DuckIntegerDivide) bop()                          {}
func (DuckIntegerDivide) Format(f fmt.State, verb rune) { f.Write([]byte("//")) }

// MySQL [`DIV`](https://dev.mysql.com/doc/refman/8.0/en/arithmetic-functions.html) integer division
type MyIntegerDivide struct{}

var _ BinaryOperator = (*MyIntegerDivide)(nil)

func (MyIntegerDivide) op()                           {}
func (MyIntegerDivide) bop()                          {}
func (MyIntegerDivide) Format(f fmt.State, verb rune) { f.Write([]byte("DIV")) }

type Custom struct {
	OP string
}

var _ BinaryOperator = (*Custom)(nil)

func (Custom) op()                             {}
func (Custom) bop()                            {}
func (o Custom) Format(f fmt.State, verb rune) { f.Write([]byte(o.OP)) }

// Bitwise XOR, e.g. `a # b` (PostgreSQL-specific)
type PGBitwiseXor struct{}

var _ BinaryOperator = (*PGBitwiseXor)(nil)

func (PGBitwiseXor) op()                           {}
func (PGBitwiseXor) bop()                          {}
func (PGBitwiseXor) Format(f fmt.State, verb rune) { f.Write([]byte("#")) }

// Bitwise shift left, e.g. `a << b` (PostgreSQL-specific)
type PGBitwiseShiftLeft struct{}

var _ BinaryOperator = (*PGBitwiseShiftLeft)(nil)

func (PGBitwiseShiftLeft) op()                           {}
func (PGBitwiseShiftLeft) bop()                          {}
func (PGBitwiseShiftLeft) Format(f fmt.State, verb rune) { f.Write([]byte("<<")) }

// Bitwise shift right, e.g. `a >> b` (PostgreSQL-specific)
type PGBitwiseShiftRight struct{}

var _ BinaryOperator = (*PGBitwiseShiftRight)(nil)

func (PGBitwiseShiftRight) op()                           {}
func (PGBitwiseShiftRight) bop()                          {}
func (PGBitwiseShiftRight) Format(f fmt.State, verb rune) { f.Write([]byte(">>")) }

// Exponent, e.g. `a ^ b` (PostgreSQL-specific)
type PGExp struct{}

var _ BinaryOperator = (*PGExp)(nil)

func (PGExp) op()                           {}
func (PGExp) bop()                          {}
func (PGExp) Format(f fmt.State, verb rune) { f.Write([]byte("^")) }

// Overlap operator, e.g. `a && b` (PostgreSQL-specific)
type PGOverlap struct{}

var _ BinaryOperator = (*PGOverlap)(nil)

func (PGOverlap) op()                           {}
func (PGOverlap) bop()                          {}
func (PGOverlap) Format(f fmt.State, verb rune) { f.Write([]byte("&&")) }

// String matches regular expression (case sensitively), e.g. `a ~ b` (PostgreSQL-specific)
type PGRegexMatch struct{}

var _ BinaryOperator = (*PGRegexMatch)(nil)

func (PGRegexMatch) op()                           {}
func (PGRegexMatch) bop()                          {}
func (PGRegexMatch) Format(f fmt.State, verb rune) { f.Write([]byte("~")) }

// String matches regular expression (case insensitively), e.g. `a ~* b` (PostgreSQL-specific)
type PGRegexIMatch struct{}

var _ BinaryOperator = (*PGRegexIMatch)(nil)

func (PGRegexIMatch) op()                           {}
func (PGRegexIMatch) bop()                          {}
func (PGRegexIMatch) Format(f fmt.State, verb rune) { f.Write([]byte("~*")) }

// String does not match regular expression (case sensitively), e.g. `a !~ b` (PostgreSQL-specific)
type PGRegexNotMatch struct{}

var _ BinaryOperator = (*PGRegexNotMatch)(nil)

func (PGRegexNotMatch) op()                           {}
func (PGRegexNotMatch) bop()                          {}
func (PGRegexNotMatch) Format(f fmt.State, verb rune) { f.Write([]byte("!~")) }

// String does not match regular expression (case insensitively), e.g. `a !~* b` (PostgreSQL-specific)
type PGRegexNotIMatch struct{}

var _ BinaryOperator = (*PGRegexNotIMatch)(nil)

func (PGRegexNotIMatch) op()                           {}
func (PGRegexNotIMatch) bop()                          {}
func (PGRegexNotIMatch) Format(f fmt.State, verb rune) { f.Write([]byte("!~*")) }

// String matches pattern (case sensitively), e.g. `a ~~ b` (PostgreSQL-specific)
type PGLikeMatch struct{}

var _ BinaryOperator = (*PGLikeMatch)(nil)

func (PGLikeMatch) op()                           {}
func (PGLikeMatch) bop()                          {}
func (PGLikeMatch) Format(f fmt.State, verb rune) { f.Write([]byte("~~")) }

// String matches pattern (case insensitively), e.g. `a ~~* b` (PostgreSQL-specific)
type PGILikeMatch struct{}

var _ BinaryOperator = (*PGILikeMatch)(nil)

func (PGILikeMatch) op()                           {}
func (PGILikeMatch) bop()                          {}
func (PGILikeMatch) Format(f fmt.State, verb rune) { f.Write([]byte("~~*")) }

// String does not match pattern (case sensitively), e.g. `a !~~ b` (PostgreSQL-specific)
type PGNotLikeMatch struct{}

var _ BinaryOperator = (*PGNotLikeMatch)(nil)

func (PGNotLikeMatch) op()                           {}
func (PGNotLikeMatch) bop()                          {}
func (PGNotLikeMatch) Format(f fmt.State, verb rune) { f.Write([]byte("!~~")) }

// String does not match pattern (case insensitively), e.g. `a !~~* b` (PostgreSQL-specific)
type PGNotILikeMatch struct{}

var _ BinaryOperator = (*PGNotILikeMatch)(nil)

func (PGNotILikeMatch) op()                           {}
func (PGNotILikeMatch) bop()                          {}
func (PGNotILikeMatch) Format(f fmt.State, verb rune) { f.Write([]byte("!~~*")) }

// String "starts with", eg: `a ^@ b` (PostgreSQL-specific)
type PGStartsWith struct{}

var _ BinaryOperator = (*PGStartsWith)(nil)

func (PGStartsWith) op()                           {}
func (PGStartsWith) bop()                          {}
func (PGStartsWith) Format(f fmt.State, verb rune) { f.Write([]byte("^@")) }

// PostgreSQL-specific custom operator.
//
// See [CREATE OPERATOR](https://www.postgresql.org/docs/current/sql-createoperator.html)
// for more information.
type PGCustomBinaryOperator struct {
	Ident []string
}

var _ BinaryOperator = (*PGCustomBinaryOperator)(nil)

func (PGCustomBinaryOperator) op()  {}
func (PGCustomBinaryOperator) bop() {}
func (o PGCustomBinaryOperator) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "OPERATOR(%s)", strings.Join(o.Ident, "."))
}
