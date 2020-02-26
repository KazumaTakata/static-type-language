package ir_gen

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/parser"
	"strconv"
)

type Operator int

var factorOpToOp = map[parser.FactorOp]Operator{parser.MUL: MUL, parser.DIV: DIV}
var termOpToOp = map[parser.TermOp]Operator{parser.ADD: ADD, parser.SUB: SUB}

const (
	ADD Operator = iota + 1
	SUB
	MUL
	DIV
	EQUAL
	LT
	NONE
)

func (e Operator) String() string {

	switch e {
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	case EQUAL:
		return "=="

	case NONE:
		return ""
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type OperandType int

const (
	Int OperandType = iota + 1
	Float
	String
	Ident
)

func (e OperandType) String() string {

	switch e {
	case Int:
		return "Int"
	case Float:
		return "Float"
	case String:
		return "String"
	case Ident:
		return "Ident"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type Operand struct {
	Int    int
	Float  float64
	String string
	Id     string
	Bool   bool
	Type   OperandType
	IfTmp  bool
}

func (Op *Operand) Str() string {

	switch Op.Type {
	case Ident:
		{
			return Op.Id
		}
	case String:
		{
			return Op.String
		}
	case Int:
		{
			return strconv.Itoa(Op.Int)
		}
	case Float:
		{
			return fmt.Sprintf("%f", Op.Float)
		}

	}

	return ""
}
