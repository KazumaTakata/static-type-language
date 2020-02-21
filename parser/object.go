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
	ElementType basic_type.Variable_Type
	Value       []*Object
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

func PrintObject(object Object) {
	switch object.Type {
	case PrimitiveType:
		{
			printPrimitive(*object.Primitive)
		}
	case ArrayType:
		{
			fmt.Printf("[")
			for i, element := range object.Array.Value {
				PrintObject(*element)

				if len(object.Array.Value)-1 != i {
					fmt.Printf(", ")
				}
			}
			fmt.Printf("]")

		}
	}
}

func printPrimitive(primitive PrimitiveObj) {

	switch primitive.Type {
	case basic_type.INT:
		{
			fmt.Printf("%d", primitive.Int)
		}
	case basic_type.DOUBLE:
		{
			fmt.Printf("%v", primitive.Double)
		}
	case basic_type.STRING:
		{
			fmt.Printf("%s", primitive.String)
		}
	}

}
