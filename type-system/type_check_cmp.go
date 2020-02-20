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

	fmt.Printf("\ntype mismatch: %v can not be %ved with %v\n", term1_type, op, term2_type)
	os.Exit(1)

	return basic_type.INT

}

func Type_Check_Cmp(cmp_expr *parser.Cmp_expr, symbol_env *parser.Symbol_Env) basic_type.Variable_Type {

	left_type := Type_Check_Arith(cmp_expr.Left, symbol_env)

	if cmp_expr.Right != nil {
		right_type := Type_Check_Arith(cmp_expr.Right, symbol_env)
		if left_type != right_type {
			fmt.Printf("\ntype mismatch: %v can not be %ved with %v\n", left_type, cmp_expr.Op, right_type)
			os.Exit(1)
		} else {
			cmp_expr.Type = basic_type.BOOL
			return cmp_expr.Type
		}
	}

	cmp_expr.Type = left_type

	return cmp_expr.Type
}
