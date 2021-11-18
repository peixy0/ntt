package runtime

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"unicode"

	"github.com/nokia/ntt/internal/ttcn3/ast"
)

type Object interface {
	Inspect() string
	Type() ObjectType
}

type ObjectType string

const (
	UNKNOWN       ObjectType = "unknown"
	UNDEFINED     ObjectType = "undefined"
	RUNTIME_ERROR ObjectType = "runtime error"
	RETURN_VALUE  ObjectType = "return value"
	INTEGER       ObjectType = "integer"
	FLOAT         ObjectType = "float"
	BOOL          ObjectType = "boolean"
	STRING        ObjectType = "string"
	BITSTRING     ObjectType = "bitstring"
	FUNCTION      ObjectType = "function"
	LIST          ObjectType = "list"
	BUILTIN_OBJ   ObjectType = "builtin function"
	VERDICT       ObjectType = "verdict"

	Bit    Unit = 1
	Hex    Unit = 4
	Octett Unit = 8

	NoneVerdict   Verdict = "none"
	PassVerdict   Verdict = "pass"
	InconcVerdict Verdict = "inconc"
	FailVerdict   Verdict = "fail"
	ErrorVerdict  Verdict = "error"
)

type Unit int

func (u Unit) Base() int {
	switch u {
	case Bit:
		return 2
	case Hex, Octett:
		return 16
	default:
		return -1
	}
}

var (
	ErrSyntax = errors.New("invalid syntax")
	Undefined = &undefined{}
)

type undefined struct{}

func (u *undefined) Inspect() string  { return "undefined" }
func (u *undefined) Type() ObjectType { return UNDEFINED }

type Error struct {
	Message string
}

func (e *Error) Error() string    { return e.Message }
func (e *Error) Type() ObjectType { return RUNTIME_ERROR }
func (e *Error) Inspect() string  { return fmt.Sprintf("Error: %s", e.Error()) }

func Errorf(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func IsError(v interface{}) bool {
	_, ok := v.(*Error)
	return ok
}

type Bool bool

func (b Bool) Type() ObjectType { return BOOL }
func (b Bool) Inspect() string  { return fmt.Sprintf("%t", b) }
func (b Bool) Bool() bool       { return bool(b) }

func NewBool(b bool) Bool {
	return Bool(b)
}

type Float struct{ *big.Float }

func (f Float) Type() ObjectType { return FLOAT }
func (f Float) Inspect() string  { return f.String() }

func NewFloat(s string) Float {
	f, _ := new(big.Float).SetString(s)
	return Float{Float: f}
}

type Int struct{ *big.Int }

func (i Int) Type() ObjectType { return INTEGER }
func (i Int) Inspect() string  { return i.String() }
func (i Int) Value() *big.Int  { return i.Int }

func NewInt(s string) Int {
	i := &big.Int{}
	i.SetString(s, 10)
	return Int{i}
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING }
func (s *String) Inspect() string  { return s.Value }

type Bitstring struct {
	Value *big.Int
	Unit  Unit
}

func (b *Bitstring) Type() ObjectType { return BITSTRING }
func (b *Bitstring) Inspect() string {
	switch b.Unit {
	case Bit:
		return fmt.Sprintf("'%b'B", b.Value)
	case Octett:
		return fmt.Sprintf("'%h'O", b.Value)
	default:
		return fmt.Sprintf("'%h'H", b.Value)
	}
}

func NewBitstring(s string) (*Bitstring, error) {
	if len(s) < 3 || s[0] != '\'' || s[len(s)-2] != '\'' {
		return nil, ErrSyntax
	}

	var unit Unit
	switch strings.ToUpper(string(s[len(s)-1])) {
	case "B":
		unit = Bit
	case "H":
		unit = Hex
	case "O":
		unit = Octett
	default:
		return nil, ErrSyntax
	}

	removeWhitespaces := func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}
	s = strings.Map(removeWhitespaces, s[1:len(s)-2])

	if i, ok := new(big.Int).SetString(s, unit.Base()); ok {
		return &Bitstring{Value: i, Unit: unit}, nil
	}

	// TODO(5nord) parse and return Bitstring templates (e.g. '01*1'B)
	return nil, ErrSyntax
}

type List struct {
	Elements []Object
}

func (l *List) Type() ObjectType { return LIST }
func (l *List) Inspect() string {
	var ss []string
	for _, obj := range l.Elements {
		if obj != nil {
			ss = append(ss, obj.Inspect())
		} else {
			ss = append(ss, "null")
		}
	}
	return "{" + strings.Join(ss, ", ") + "}"
}

type Function struct {
	Params *ast.FormalPars
	Body   *ast.BlockStmt
	Env    *Env
}

func (f *Function) Type() ObjectType { return FUNCTION }
func (f *Function) Inspect() string {
	var buf bytes.Buffer
	buf.WriteString("function(\"")
	for i, p := range f.Params.List {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(p.Name.String())
	}
	buf.WriteString(")")
	return buf.String()
}

type ReturnValue struct {
	Value Object
}

func (r *ReturnValue) Type() ObjectType { return RETURN_VALUE }
func (r *ReturnValue) Inspect() string  { return r.Value.Inspect() }

type Verdict string

func (v Verdict) Type() ObjectType { return VERDICT }
func (v Verdict) Inspect() string  { return string(v) }
