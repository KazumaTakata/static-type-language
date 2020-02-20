package parser

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/type"
)

type Symbol_Table map[string]Object

type ObjectType int

const (
	PrimitiveType ObjectType = iota + 1
	FunctionType
	ArrayType
)

func (e ObjectType) String() string {

	switch e {
	case PrimitiveType:
		return "Primitive"
	case FunctionType:
		return "Function"
	case ArrayType:
		return "Array"

	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type Object struct {
	Type      ObjectType
	Primitive *PrimitiveObj
	Function  *Def_stmt
	Array     *ArrayObj
}

type ArrayObj struct {
	Type  basic_type.Type
	Value []*PrimitiveObj
}

type PrimitiveObj struct {
	Type   basic_type.Type
	Int    int
	Double float64
	String string
	Bool   bool
}

type Symbol_Env struct {
	Table        Symbol_Table
	Parent_Env   *Symbol_Env
	Return_Value *Object
	Return_Type  basic_type.Variable_Type
}
