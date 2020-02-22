package basic_type

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/lexer"
)

type Type int

const (
	INT Type = iota + 1
	DOUBLE
	STRING
	BOOL
)

var LexerTypeToType = map[lexer.TokenType]Type{lexer.INT: INT, lexer.DOUBLE: DOUBLE, lexer.STRING: STRING}
var BoolPrimitiveType Variable_Type = Variable_Type{DataStructType: PRIMITIVE, Primitive: &PrimitiveType{Type: BOOL}}
var IntPrimitiveType Variable_Type = Variable_Type{DataStructType: PRIMITIVE, Primitive: &PrimitiveType{Type: INT}}

func (e Type) String() string {

	switch e {
	case INT:
		return "INT"
	case DOUBLE:
		return "DOUBLE"
	case STRING:
		return "STRING"
	case BOOL:
		return "BOOL"

	default:
		return fmt.Sprintf("NULL")
	}
}

type DataStructType int

const (
	MAP DataStructType = iota + 1
	ARRAY
	PRIMITIVE
)

func (e DataStructType) String() string {

	switch e {
	case MAP:
		return "MAP"
	case ARRAY:
		return "ARRAY"
	case PRIMITIVE:
		return "PRIMITIVE"
	default:
		return fmt.Sprintf("%d", e)
	}
}

func Variable_Equal(var1 Variable_Type, var2 Variable_Type) bool {

	if var1.DataStructType != var2.DataStructType {
		return false
	}

	switch var1.DataStructType {
	case ARRAY:
		{

			if var1.DataStructType != var2.DataStructType {
				return false
			}

			return Variable_Equal(var1.Array.ElementType, var2.Array.ElementType)

		}
	case PRIMITIVE:
		{
			if var1.Primitive.Type == var2.Primitive.Type {
				return true
			} else {
				return false
			}
		}
	}

	return false

}

func WrapWithArrayType(element_type Variable_Type) Variable_Type {
	return Variable_Type{DataStructType: ARRAY, Array: &ArrayType{ElementType: element_type}}
}

type Variable_Type struct {
	DataStructType DataStructType
	Array          *ArrayType
	Map            *MapType
	Primitive      *PrimitiveType
}

type ArrayType struct {
	ElementType Variable_Type
}

type MapType struct {
	KeyType   Variable_Type
	ValueType Variable_Type
}

type PrimitiveType struct {
	Type Type
}

func Builtin_func(name string) bool {
	builtins := map[string]bool{"len": true}
	if _, ok := builtins[name]; ok {
		return true
	}

	return false
}
