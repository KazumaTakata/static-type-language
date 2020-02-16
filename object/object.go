package object

import (
	"github.com/KazumaTakata/static-typed-language/type"
)

type Variable_Table map[string]Variable

type Variable struct {
	Type   basic_type.Type
	Int    int
	Double float64
	String string
	Bool   bool
}

type Symbol_Env struct {
	Table      Variable_Table
	Parent_Env *Symbol_Env
}
