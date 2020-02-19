package eval

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"
	"strconv"

	basic_type "github.com/KazumaTakata/static-typed-language/type"
	"os"
)

func resolve_variable(id string, symbol_env *parser.Symbol_Env) parser.Object {
	if object, ok := symbol_env.Table[id]; ok {
		return object
	} else {
		if symbol_env.Parent_Env != nil {
			return resolve_variable(id, symbol_env.Parent_Env)
		}
		fmt.Printf("\nnot defined variable %v\n", id)
		os.Exit(1)
	}
	return parser.Object{}

}

func argToObject(arg lexer.Token) parser.Object {
	switch arg.Type {
	case lexer.INT:
		{
			int_value, _ := strconv.Atoi(arg.Value)
			return parser.Object{Type: parser.PrimitiveType, Primitive: &parser.PrimitiveObj{Type: basic_type.INT, Int: int_value}}
		}

	case lexer.DOUBLE:
		{
			double_value, _ := strconv.ParseFloat(arg.Value, 64)
			return parser.Object{Type: parser.PrimitiveType, Primitive: &parser.PrimitiveObj{Type: basic_type.DOUBLE, Double: double_value}}
		}

	case lexer.STRING:
		{
			return parser.Object{Type: parser.PrimitiveType, Primitive: &parser.PrimitiveObj{Type: basic_type.STRING, String: arg.Value}}
		}

	}

	return parser.Object{}
}

func handle_func_call(object parser.Object, factor parser.Factor, symbol_env *parser.Symbol_Env) *parser.Object {
	if object.Type != parser.FunctionType {
		fmt.Printf("\nvariable %s is not function\n", factor.Id)
		os.Exit(1)
	}

	env := object.Function.Symbol_Env
	for i, arg := range object.Function.Args {
		if factor.Args[i].Type == lexer.IDENT {
			param_object := resolve_variable(factor.Args[i].Value, symbol_env)
			switch param_object.Type {
			case parser.PrimitiveType:
				{
					env.Table[arg.Ident] = param_object
				}
			}
		} else {
			env.Table[arg.Ident] = argToObject(factor.Args[i])
		}
	}

	Eval_Stmts(object.Function.Stmts, env)
	returned_value := env.Return_Value

	return returned_value

}

func Arith_Factors_INT(factors []parser.TermElement, symbol_env *parser.Symbol_Env) int {

	if len(factors) == 1 {

		if factors[0].Factor.Type == lexer.IDENT {

			variable := resolve_ident(factors[0].Factor, symbol_env)
			return variable.Int

		} else {
			return factors[0].Factor.Int
		}
	}

	var result int

	for i, factor := range factors {
		if i == 0 {

			if factor.Factor.Type == lexer.IDENT {

				variable := resolve_ident(factor.Factor, symbol_env)
				result = variable.Int

			} else {
				result = factor.Factor.Int
			}

			continue
		}
		switch factor.Op {
		case parser.MUL:
			{
				if factor.Factor.Type == lexer.IDENT {

					variable := resolve_ident(factor.Factor, symbol_env)
					result = result * variable.Int

				} else {
					result = result * factor.Factor.Int
				}
			}

		case parser.DIV:
			{
				if factor.Factor.Type == lexer.IDENT {

					variable := resolve_ident(factor.Factor, symbol_env)
					result = result / variable.Int
				} else {
					result = result / factor.Factor.Int
				}

			}
		}

	}
	return result
}

func resolve_ident(factor parser.Factor, symbol_env *parser.Symbol_Env) *parser.PrimitiveObj {
	object := resolve_variable(factor.Id, symbol_env)

	if factor.IsCall {
		returned_value := handle_func_call(object, factor, symbol_env)
		return returned_value.Primitive

	} else {
		return object.Primitive
	}

}
func Arith_Factors_STRING(factors []parser.TermElement, symbol_env *parser.Symbol_Env) string {

	if len(factors) == 1 {

		if factors[0].Factor.Type == lexer.IDENT {
			variable := resolve_ident(factors[0].Factor, symbol_env)
			return variable.String
		}

		return factors[0].Factor.String
	}
	os.Exit(1)
	return ""
}

func Arith_Factors_BOOL(factors []parser.TermElement, symbol_env *parser.Symbol_Env) bool {

	if len(factors) == 1 {

		if factors[0].Factor.Type == lexer.IDENT {
			variable := resolve_ident(factors[0].Factor, symbol_env)
			return variable.Bool
		}

		return factors[0].Factor.Bool
	}
	os.Exit(1)
	return true
}

func Arith_Factors_DOUBLE(factors []parser.TermElement, symbol_env *parser.Symbol_Env) float64 {

	if len(factors) == 1 {

		if factors[0].Factor.Type == lexer.IDENT {
			variable := resolve_ident(factors[0].Factor, symbol_env)

			return variable.Double
		}

		return factors[0].Factor.Float
	}

	var result float64

	for i, factor := range factors {
		if i == 0 {

			switch factor.Factor.Type {

			case lexer.IDENT:
				{
					variable := resolve_ident(factor.Factor, symbol_env)

					if variable.Type == basic_type.DOUBLE {
						result = variable.Double
					} else if variable.Type == basic_type.INT {
						result = float64(variable.Int)
					}

				}

			case lexer.INT:
				{
					result = float64(factor.Factor.Int)
				}
			case lexer.DOUBLE:
				{
					result = factor.Factor.Float
				}
			}
			continue
		}
		switch factor.Op {
		case parser.MUL:
			{
				switch factor.Factor.Type {
				case lexer.IDENT:
					{
						variable := resolve_ident(factor.Factor, symbol_env)

						if variable.Type == basic_type.DOUBLE {
							result = result * variable.Double
						} else if variable.Type == basic_type.INT {
							result = result * float64(variable.Int)
						}

					}

				case lexer.INT:
					{
						result = result * float64(factor.Factor.Int)
					}
				case lexer.DOUBLE:
					{
						result = result * factor.Factor.Float
					}
				}
			}

		case parser.DIV:
			{

				switch factor.Factor.Type {
				case lexer.IDENT:
					{
						variable := resolve_ident(factor.Factor, symbol_env)
						if variable.Type == basic_type.DOUBLE {
							result = result / variable.Double
						} else if variable.Type == basic_type.INT {
							result = result / float64(variable.Int)
						}
					}

				case lexer.INT:
					{
						result = result / float64(factor.Factor.Int)
					}
				case lexer.DOUBLE:
					{
						result = result / factor.Factor.Float
					}
				}
			}
		}

	}
	return result
}
