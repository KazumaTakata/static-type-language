package type_system

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/lexer"
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

func Type_Check_Cmp_Arith(term1_type basic_type.Type, term2_type basic_type.Type, op parser.ArithOp) basic_type.Type {

	if term1_type == term2_type {
		return basic_type.BOOL
	}

	fmt.Printf("\ntype mismatch: %v can not be %ved with %v\n", term1_type, op, term2_type)
	os.Exit(1)

	return basic_type.INT

}

func resolve_name(id string, symbol_env *parser.Symbol_Env) parser.Object {
	if object, ok := symbol_env.Table[id]; ok {
		return object
	} else {
		if symbol_env.Parent_Env != nil {
			return resolve_name(id, symbol_env.Parent_Env)
		}

		fmt.Printf("\nnot defined variable %v\n", id)
		os.Exit(1)

	}

	return parser.Object{}
}

func get_Type_of_Factor(factor parser.Factor, symbol_env *parser.Symbol_Env) basic_type.Type {

	if factor.Type == lexer.IDENT {

		object := resolve_name(factor.Id, symbol_env)

		if factor.IsCall {
			if object.Type != parser.FunctionObj {
				fmt.Printf("\nvariable %s is not function\n", factor.Id)
				os.Exit(1)
			}
			params := factor.Args
			function := object.Function

			for i, arg := range function.Args {
				param := params[i]
				if param.Type == lexer.IDENT {
					param_object := resolve_name(param.Value, symbol_env)
					switch param_object.Type {
					case parser.VariableObj:
						{
							if param_object.Variable.Type != arg.Type {
								fmt.Printf("\nparam type mismatch:variable of type %v can not be passed as param of type %v\n", param_object.Variable.Type, arg.Type)
								os.Exit(1)
							}
						}
					}

				} else {
					if arg.Type != basic_type.LexerTypeToType[params[i].Type] {
						fmt.Printf("\nparam type mismatch: %v can not be passed as type %v\n", param.Type, arg.Type)
						os.Exit(1)
					}
				}

			}

		} else {
			return object.Variable.Type
		}
	}

	return basic_type.LexerTypeToType[factor.Type]
}

func Type_Check_Cmp(cmp_expr *parser.Cmp_expr, symbol_env *parser.Symbol_Env) basic_type.Type {

	cmp_expr.Type = Type_Check_Cmp_Ariths(cmp_expr.Ariths, symbol_env)

	return cmp_expr.Type
}

func Type_Check_Cmp_Ariths(ariths []parser.CmpElement, symbol_env *parser.Symbol_Env) basic_type.Type {

	if len(ariths) == 1 {
		return Type_Check_Arith_Terms(ariths[0].Arith.Terms, symbol_env)
	}

	var operand1_type basic_type.Type
	var operand2_type basic_type.Type

	for i, arith := range ariths {
		if i == 0 {
			operand1_type = Type_Check_Arith_Terms(arith.Arith.Terms, symbol_env)
			ariths[i].Arith.Type = operand1_type
			continue
		}

		operand2_type = operand1_type

		operand1_type = Type_Check_Arith_Terms(arith.Arith.Terms, symbol_env)
		ariths[i].Arith.Type = operand1_type

		if operand1_type != operand2_type {

			fmt.Printf("\ntype mismatch: %v can not be %ved with %v\n", operand2_type, arith.Op, operand1_type)
			os.Exit(1)

		} else {
			operand1_type = basic_type.BOOL
		}
	}

	return operand1_type

}

func Type_Check_Arith(arith *parser.Arith_expr, symbol_env *parser.Symbol_Env) basic_type.Type {

	arith.Type = Type_Check_Arith_Terms(arith.Terms, symbol_env)

	return arith.Type
}

func Type_Check_Arith_Terms(terms []parser.ArithElement, symbol_env *parser.Symbol_Env) basic_type.Type {

	if len(terms) == 1 {
		return Type_Check_Arith_Factors(terms[0].Term.Factors, symbol_env)
	}

	var operand1_type basic_type.Type
	var operand2_type basic_type.Type

	for i, term := range terms {
		if i == 0 {
			operand1_type = Type_Check_Arith_Factors(term.Term.Factors, symbol_env)
			terms[i].Term.Type = operand1_type
			continue
		}

		operand2_type = operand1_type
		operand1_type = Type_Check_Arith_Factors(term.Term.Factors, symbol_env)
		terms[i].Term.Type = operand1_type

		operand1_type = Type_Check_Arith_Term(operand2_type, operand1_type, term.Op)
	}

	return operand1_type

}

func Type_Check_Arith_Factors(factors []parser.TermElement, symbol_env *parser.Symbol_Env) basic_type.Type {

	if len(factors) == 1 {
		return get_Type_of_Factor(factors[0].Factor, symbol_env)
	}

	var operand1_type basic_type.Type
	var operand2_type basic_type.Type

	for i, factor := range factors {
		if i == 0 {
			operand1_type = get_Type_of_Factor(factor.Factor, symbol_env)
			continue
		}

		operand2_type = operand1_type
		operand1_type = Type_Check_Arith_Factor(operand2_type, get_Type_of_Factor(factor.Factor, symbol_env), factor.Op)
	}

	return operand1_type

}
