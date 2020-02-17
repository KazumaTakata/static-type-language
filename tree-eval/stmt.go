package eval

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/parser"
	basic_type "github.com/KazumaTakata/static-typed-language/type"
)

func Eval_Stmts(stmts []parser.Stmt, symbol_env *parser.Symbol_Env) {

	for _, stmt := range stmts {
		if_return := Eval_Stmt(stmt, symbol_env)

		if if_return {
			break
		}
	}
}

func Calc_Arith(expr *parser.Arith_expr, symbol_env *parser.Symbol_Env) parser.Object {

	switch expr.Type {
	case basic_type.INT:
		{
			result := Arith_Terms_INT(expr.Terms, symbol_env)
			return parser.Object{Type: parser.VariableObj, Variable: &parser.Variable{Int: result, Type: basic_type.INT}}

		}
	case basic_type.DOUBLE:
		{
			result := Arith_Terms_DOUBLE(expr.Terms, symbol_env)
			return parser.Object{Type: parser.VariableObj, Variable: &parser.Variable{Double: result, Type: basic_type.DOUBLE}}

		}
	case basic_type.STRING:
		{
			result := Arith_Terms_STRING(expr.Terms, symbol_env)
			return parser.Object{Type: parser.VariableObj, Variable: &parser.Variable{String: result, Type: basic_type.STRING}}

		}

	case basic_type.BOOL:
		{
			result := Arith_Terms_BOOL(expr.Terms, symbol_env)
			return parser.Object{Type: parser.VariableObj, Variable: &parser.Variable{Bool: result, Type: basic_type.BOOL}}

		}

	}

	return parser.Object{}
}

func Eval_Stmt(stmt parser.Stmt, symbol_env *parser.Symbol_Env) bool {

	switch stmt.Type {

	case parser.DEF_STMT:
		{

		}
	case parser.RETURN_STMT:
		{
			return_value := Calc_Arith(&stmt.Return.Cmp_expr.Ariths[0].Arith, symbol_env)
			symbol_env.Return_Value = &return_value

			return true

		}

	case parser.FOR_STMT:
		{
			for Eval_Cmp(stmt.For.Cmp_expr, symbol_env) {
				Eval_Stmts(stmt.For.Stmts, stmt.For.Symbol_Env)
			}

		}
	case parser.IF_STMT:
		{
			if Eval_Cmp(stmt.If.Cmp_expr, symbol_env) {
				Eval_Stmts(stmt.If.Stmts, stmt.If.Symbol_Env)
			}
		}

	case parser.DECL_STMT:
		{

			result := Calc_Arith(stmt.Decl.Expr, symbol_env)
			symbol_env.Table[stmt.Decl.Id] = result
			//fmt.Printf("%+v", result)

		}

	case parser.EXPR:
		{

			switch stmt.Expr.Type {
			case basic_type.INT:
				{
					result := Arith_Terms_INT(stmt.Expr.Terms, symbol_env)
					fmt.Printf("%+v\n", result)

				}
			case basic_type.DOUBLE:
				{
					result := Arith_Terms_DOUBLE(stmt.Expr.Terms, symbol_env)

					fmt.Printf("%+v\n", result)

				}
			case basic_type.STRING:
				{
					result := Arith_Terms_STRING(stmt.Expr.Terms, symbol_env)

					fmt.Printf("%+v\n", result)

				}
			case basic_type.BOOL:
				{
					result := Arith_Terms_BOOL(stmt.Expr.Terms, symbol_env)

					fmt.Printf("%+v\n", result)

				}

			}
		}

	}

	return false
}
