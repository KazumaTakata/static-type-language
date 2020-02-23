package type_system

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/parser"
	"github.com/KazumaTakata/static-typed-language/type"
	"os"
)

var NumberType = map[basic_type.Type]bool{basic_type.INT: true, basic_type.DOUBLE: true}

func Type_Check_Cmp_Arith(term1_type basic_type.Type, term2_type basic_type.Type, op parser.ArithOp) basic_type.Type {

	if term1_type == term2_type {
		return basic_type.BOOL
	}

	fmt.Printf("\ntype    mismatch: %v can not be %ved with %v\n", term1_type, op, term2_type)
	os.Exit(1)

	return basic_type.INT

}

func Type_Check_Cmp(cmp_expr *parser.Cmp_expr, symbol_env *parser.Symbol_Env) basic_type.Variable_Type {

	left_type := Type_Check_Arith(cmp_expr.Left, symbol_env)

	if cmp_expr.Right != nil {
		right_type := Type_Check_Arith(cmp_expr.Right, symbol_env)

		if !basic_type.Variable_Primitive_Equal(left_type, right_type) {
			fmt.Printf("\ncmp operand should be primitive and have same type: got: (%+v : %+v)\n", left_type, right_type)
			os.Exit(1)
		}

		cmp_expr.Type = basic_type.Variable_Type{DataStructType: basic_type.PRIMITIVE, Primitive: &basic_type.PrimitiveType{Type: basic_type.BOOL}}
		return cmp_expr.Type

	} else {

		cmp_expr.Type = left_type
		return cmp_expr.Type

	}
}

func Type_Check_Logic(logic_expr *parser.Logic_expr, symbol_env *parser.Symbol_Env) basic_type.Variable_Type {

	if len(logic_expr.Cmps) == 1 {
		cmp_type := Type_Check_Cmp(&logic_expr.Cmps[0], symbol_env)
		logic_expr.Cmps[0].Type = cmp_type
	}

	var operand_variable_type basic_type.Variable_Type

	for i, cmp := range logic_expr.Cmps {
		operand_variable_type = Type_Check_Cmp(&cmp, symbol_env)

		if !basic_type.Variable_Equal(operand_variable_type, basic_type.BoolPrimitiveType) {
			fmt.Printf("\nlogic operand should be boolean type: got %+v.\n", operand_variable_type)
			os.Exit(1)
		}

		logic_expr.Cmps[i].Type = operand_variable_type

	}

	logic_expr.Type = basic_type.BoolPrimitiveType

	return basic_type.BoolPrimitiveType

}
