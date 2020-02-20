package type_system

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"
	"github.com/KazumaTakata/static-typed-language/type"
	"os"
)

func resolve_name_pointer(id string, symbol_env *parser.Symbol_Env) *parser.Object {
	if object, ok := symbol_env.Table[id]; ok {
		return &object
	} else {
		if symbol_env.Parent_Env != nil {
			return resolve_name_pointer(id, symbol_env.Parent_Env)
		}

		fmt.Printf("\nnot defined variable %v\n", id)
		os.Exit(1)

	}

	return nil
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

func get_Type_of_Factor(factor parser.Factor, symbol_env *parser.Symbol_Env) basic_type.Variable_Type {

	if factor.Type == lexer.IDENT {

		object := resolve_name(factor.Id, symbol_env)

		switch factor.FactorType {

		case parser.FuncCall:
			{
				if object.Type != parser.FunctionType {
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
						case parser.PrimitiveType:
							{
								if basic_type.PRIMITIVE != arg.Type.Type {
									fmt.Printf("\nparam type mismatch:argument type  %+v is not primitive\n", arg.Type.Type)
									os.Exit(1)
								}

								if arg.Type.Primitive.Type != param_object.Primitive.Type {
									fmt.Printf("\nparam type mismatch:argument type  %+v is not %+v\n", arg.Type.Primitive.Type, param_object.Primitive.Type)
									os.Exit(1)

								}

							}

						case parser.ArrayType:
							{

							}
						}

					} else {
						if arg.Type.Type != basic_type.PRIMITIVE {
							fmt.Printf("\nparam type mismatch: %v is not primitive type \n", arg.Type)
							os.Exit(1)
						}

						if arg.Type.Primitive.Type != basic_type.LexerTypeToType[params[i].Type] {
							fmt.Printf("\nparam type mismatch: %v can not be passed as type %v\n", param.Type, arg.Type)
							os.Exit(1)
						}
					}

				}
				return function.Return_type
			}
		case parser.ArrayMapAccess:
			{

				if object.Type != parser.ArrayType {
					fmt.Printf("\nvariable %s is not function\n", factor.Id)
					os.Exit(1)
				}

				index_type := Type_Check_Arith(factor.AccessIndex, symbol_env)
				if index_type.Type != basic_type.PRIMITIVE {
					fmt.Printf("\nparam type mismatch: %v is not primitive type \n", index_type.Type)
					os.Exit(1)
				}

				if index_type.Primitive.Type != basic_type.INT {
					fmt.Printf("\narray index type should be int type: got %+v\n", index_type)
					os.Exit(1)
				}
				return basic_type.Variable_Type{Type: basic_type.PRIMITIVE, Primitive: &basic_type.PrimitiveType{Type: object.Array.Type}}
			}
		default:
			{
				switch object.Type {
				case parser.ArrayType:
					{
						return basic_type.Variable_Type{Type: basic_type.ARRAY, Array: &basic_type.ArrayType{Type: basic_type.Variable_Type{Type: basic_type.PRIMITIVE, Primitive: &basic_type.PrimitiveType{Type: object.Array.Type}}}}
					}
				case parser.PrimitiveType:
					{
						return basic_type.Variable_Type{Type: basic_type.PRIMITIVE, Primitive: &basic_type.PrimitiveType{Type: object.Primitive.Type}}

					}

				}

			}
		}
	}

	return basic_type.Variable_Type{Type: basic_type.PRIMITIVE, Primitive: &basic_type.PrimitiveType{Type: basic_type.LexerTypeToType[factor.Type]}}
}

func Type_Check_Arith_Factors(factors []parser.TermElement, symbol_env *parser.Symbol_Env) basic_type.Variable_Type {

	if len(factors) == 1 {
		return get_Type_of_Factor(factors[0].Factor, symbol_env)
	}

	var operand1_type basic_type.Type
	var operand2_type basic_type.Type

	var operand1_variable_type basic_type.Variable_Type
	var operand2_variable_type basic_type.Variable_Type

	for i, factor := range factors {
		if i == 0 {
			operand1_variable_type = get_Type_of_Factor(factor.Factor, symbol_env)
			if operand1_variable_type.Type != basic_type.PRIMITIVE {

				fmt.Printf("\ntype %+v cannot be added or subed\n", operand1_variable_type.Type)
				os.Exit(1)

			}

			operand1_type = operand1_variable_type.Primitive.Type
			continue
		}

		operand2_type = operand1_type
		operand2_variable_type = operand1_variable_type

		operand1_variable_type = get_Type_of_Factor(factor.Factor, symbol_env)
		operand1_type = operand1_variable_type.Primitive.Type

		if operand1_variable_type.Type != basic_type.PRIMITIVE {

			fmt.Printf("\ntype %+v cannot be added or subed\n", operand1_variable_type.Type)
			os.Exit(1)

		}

		operand1_type = Type_Check_Arith_Factor(operand2_type, operand1_type, factor.Op)
	}

	return basic_type.Variable_Type{Type: basic_type.PRIMITIVE, Primitive: &basic_type.PrimitiveType{Type: operand1_type}}

}

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
