package type_system

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/parser"
	"github.com/KazumaTakata/static-typed-language/type"
	"os"
)

func Type_Check_Arith(arith *parser.Arith_expr, symbol_env *parser.Symbol_Env) basic_type.Variable_Type {

	arith.Type = Type_Check_Arith_Terms(arith.Terms, symbol_env)

	return arith.Type
}

func Type_Check_Arith_Terms(terms []parser.ArithElement, symbol_env *parser.Symbol_Env) basic_type.Variable_Type {

	if len(terms) == 1 {
		return Type_Check_Arith_Factors(terms[0].Term.Factors, symbol_env)
	}

	var operand1_type basic_type.Type
	var operand2_type basic_type.Type

	var operand1_variable_type basic_type.Variable_Type
	var operand2_variable_type basic_type.Variable_Type

	for i, term := range terms {
		if i == 0 {
			operand1_variable_type = Type_Check_Arith_Factors(term.Term.Factors, symbol_env)

			if operand1_variable_type.DataStructType != basic_type.PRIMITIVE {
				fmt.Printf("\ntype %+v cannot be added or subed\n", operand1_variable_type.DataStructType)
				os.Exit(1)
			}

			terms[i].Term.Type = operand1_variable_type
			continue
		}

		operand2_type = operand1_type
		operand2_variable_type = operand1_variable_type
		operand2_type = operand2_variable_type.Primitive.Type

		operand1_variable_type = Type_Check_Arith_Factors(term.Term.Factors, symbol_env)

		if operand1_variable_type.DataStructType != basic_type.PRIMITIVE {

			fmt.Printf("\ntype %+v cannot be added or subed\n", operand1_variable_type.DataStructType)
			os.Exit(1)

		}

		operand1_type = operand1_variable_type.Primitive.Type
		terms[i].Term.Type = operand1_variable_type

		operand1_type = Type_Check_Arith_Term(operand2_type, operand1_type, term.Op)
	}

	return basic_type.Variable_Type{DataStructType: basic_type.PRIMITIVE, Primitive: &basic_type.PrimitiveType{Type: operand1_type}}

}
func Type_Check_Arith_Term(term1_type basic_type.Type, term2_type basic_type.Type, op parser.TermOp) basic_type.Type {

	is_factor1_number := false
	is_factor2_number := false

	if _, ok := NumberType[term1_type]; ok {
		is_factor1_number = true
	}

	if _, ok := NumberType[term2_type]; ok {
		is_factor2_number = true
	}

	if is_factor1_number && is_factor2_number {
		if term1_type < term2_type {
			return term2_type
		} else {
			return term1_type
		}
	} else if is_factor1_number && !is_factor2_number {

		fmt.Printf("\ntype mismatch: %+v can not be %ved with %v\n", term1_type, op, term2_type)
		os.Exit(1)
	} else if !is_factor1_number && is_factor2_number {
		fmt.Printf("\ntype mismatch: %+v can not be %ved with %v\n", term1_type, op, term2_type)
		os.Exit(1)
	} else {

		if term1_type == term2_type && term1_type == basic_type.STRING && op == parser.ADD {
			return basic_type.STRING
		} else {
			fmt.Printf("\ntype mismatch: %v can not be %ved with %v\n", term1_type, op, term2_type)
			os.Exit(1)
		}

	}

	return basic_type.INT

}
