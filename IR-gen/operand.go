package main

import (
	"fmt"
)

type Operator int

const (
	ADD Operator = iota + 1
	SUB
	MUL
	DIV
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
	case NONE:
		return "NONE"
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
