package type_checker

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/parser"
	"github.com/KazumaTakata/static-typed-language/type"
	"os"
)

var NumberType = map[basic_type.Type]bool{basic_type.INT: true, basic_type.DOUBLE: true}

func Type_Check_Arith_Factor(factor1_type basic_type.Type, factor2_type basic_type.Type, op parser.FactorOp) basic_type.Type {

	is_factor1_number := false
	is_factor2_number := false

	if _, ok := NumberType[factor1_type]; ok {
		is_factor1_number = true
	}

	if _, ok := NumberType[factor2_type]; ok {
		is_factor2_number = true
	}

	if is_factor1_number && is_factor2_number {
		if factor1_type < factor2_type {
			return factor2_type
		} else {
			return factor1_type
		}
	} else if is_factor1_number && !is_factor2_number {

		fmt.Printf("\ntype mismatch: %+v can not be %ved with %v\n", factor1_type, op, factor2_type)
		os.Exit(1)
	} else if !is_factor1_number && is_factor2_number {
		fmt.Printf("\ntype mismatch: %+v can not be %ved with %v\n", factor1_type, op, factor2_type)
		os.Exit(1)
	} else {
		fmt.Printf("\ntype mismatch: %v can not be %ved with %v\n", factor1_type, op, factor2_type)
		os.Exit(1)

	}

	return basic_type.INT

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

func get_Type_of_Factor(factor parser.Factor) basic_type.Type {
	return basic_type.LexerTypeToType[factor.Type]
}

func Type_Check_Arith(arith *parser.Arith_expr) basic_type.Type {

	arith.Type = Type_Check_Arith_Terms(arith.Terms)

	return arith.Type
}

func Type_Check_Arith_Terms(terms []parser.ArithElement) basic_type.Type {

	if len(terms) == 1 {
		return Type_Check_Arith_Factors(terms[0].Term.Factors)
	}

	var operand1_type basic_type.Type
	var operand2_type basic_type.Type

	for i, term := range terms {
		if i == 0 {
			operand1_type = Type_Check_Arith_Factors(term.Term.Factors)
			terms[i].Term.Type = operand1_type
			continue
		}

		operand2_type = operand1_type
		operand1_type = Type_Check_Arith_Factors(term.Term.Factors)
		terms[i].Term.Type = operand1_type

		operand1_type = Type_Check_Arith_Term(operand2_type, operand1_type, term.Op)
	}

	return operand1_type

}

func Type_Check_Arith_Factors(factors []parser.TermElement) basic_type.Type {

	if len(factors) == 1 {
		return get_Type_of_Factor(factors[0].Factor)
	}

	var operand1_type basic_type.Type
	var operand2_type basic_type.Type

	for i, factor := range factors {
		if i == 0 {
			operand1_type = get_Type_of_Factor(factor.Factor)
			continue
		}

		operand2_type = operand1_type
		operand1_type = Type_Check_Arith_Factor(operand2_type, get_Type_of_Factor(factor.Factor), factor.Op)
	}

	return operand1_type

}
