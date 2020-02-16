package parser

import (
	"github.com/KazumaTakata/static-typed-language/type"
)

type Symbol_Table map[string]Object

type ObjectType int

const (
	VariableObj ObjectType = iota + 1
	FunctionObj
)

type Object struct {
	Type     ObjectType
	Variable *Variable
	Function *Def_stmt
}

type Variable struct {
	Type   basic_type.Type
	Int    int
	Double float64
	String string
	Bool   bool
}

type Symbol_Env struct {
	Table      Symbol_Table
	Parent_Env *Symbol_Env
}
