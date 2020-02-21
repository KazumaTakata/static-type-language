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

func Eval_Init(init parser.Init, symbol_env *parser.Symbol_Env) parser.Object {

	switch init.Type {
	case parser.ARRAY_INIT:
		{
			arrayobj := parser.ArrayObj{ElementType: init.Array.ElementType}

			for _, init_value := range init.Array.InitValue {
				assign := Eval_Assign(*init_value, symbol_env)
				arrayobj.Value = append(arrayobj.Value, &assign)

			}

			return parser.Object{Type: parser.ArrayType, Array: &arrayobj}
		}
	case parser.MAP_INIT:
		{
		}
	}

	return parser.Object{}
}

func Eval_Assign(assign parser.Assign, symbol_env *parser.Symbol_Env) parser.Object {
	switch assign.Type {
	case parser.INIT_ASSIGN:
		{
			return Eval_Init(*assign.Init, symbol_env)
		}
	case parser.EXPR_ASSIGN:
		{
			return Calc_Arith(assign.Expr.Left, symbol_env)
		}
	}

	return parser.Object{}
}

func Calc_Arith(expr *parser.Arith_expr, symbol_env *parser.Symbol_Env) parser.Object {

	switch expr.Type.DataStructType {
	case basic_type.PRIMITIVE:
		{
			switch expr.Type.Primitive.Type {
			case basic_type.INT:
				{
					result := Arith_Terms_INT(expr.Terms, symbol_env)
					return parser.Object{Type: parser.PrimitiveType, Primitive: &parser.PrimitiveObj{Int: result, Type: basic_type.INT}}

				}
			case basic_type.DOUBLE:
				{
					result := Arith_Terms_DOUBLE(expr.Terms, symbol_env)
					return parser.Object{Type: parser.PrimitiveType, Primitive: &parser.PrimitiveObj{Double: result, Type: basic_type.DOUBLE}}

				}
			case basic_type.STRING:
				{
					result := Arith_Terms_STRING(expr.Terms, symbol_env)
					return parser.Object{Type: parser.PrimitiveType, Primitive: &parser.PrimitiveObj{String: result, Type: basic_type.STRING}}

				}

			case basic_type.BOOL:
				{
					result := Arith_Terms_BOOL(expr.Terms, symbol_env)
					return parser.Object{Type: parser.PrimitiveType, Primitive: &parser.PrimitiveObj{Bool: result, Type: basic_type.BOOL}}

				}

			}
		}

	case basic_type.ARRAY:
		{
			factor := expr.Terms[0].Term.Factors[0].Factor
			object := resolve_variable(factor.Id, symbol_env)

			if object.Type == parser.FunctionType {
				returned_value := handle_func_call(object, factor, symbol_env)
				return *returned_value

			}
			return object
		}
	}

	return parser.Object{}
}

func assign_Table(id string, symbol_env *parser.Symbol_Env, object parser.Object) {
	if _, ok := symbol_env.Table[id]; ok {
		symbol_env.Table[id] = object
	} else {
		assign_Table(id, symbol_env.Parent_Env, object)
	}

}

func Eval_Stmt(stmt parser.Stmt, symbol_env *parser.Symbol_Env) bool {

	switch stmt.Type {

	case parser.DEF_STMT:
		{

		}
	case parser.RETURN_STMT:
		{
			return_value := Calc_Arith(stmt.Return.Cmp_expr.Left, symbol_env)
			symbol_env.Return_Value = &return_value

			return true

		}

	case parser.FOR_STMT:
		{
			for Eval_Cmp_Bool(stmt.For.Cmp_expr, symbol_env) {
				Eval_Stmts(stmt.For.Stmts, stmt.For.Symbol_Env)
			}

		}
	case parser.IF_STMT:
		{
			if Eval_Cmp_Bool(stmt.If.Cmp_expr, symbol_env) {
				Eval_Stmts(stmt.If.Stmts, stmt.If.Symbol_Env)
			}
		}

	case parser.ASSIGN_STMT:
		{

			result := Eval_Assign(*stmt.Assign.Assign, symbol_env)

			if len(stmt.Assign.Indexs) > 0 {
				index := Calc_Arith(&stmt.Assign.Indexs[0], symbol_env)
				if object, ok := symbol_env.Table[stmt.Assign.Id]; ok {
					object.Array.Value[index.Primitive.Int] = &result
					symbol_env.Table[stmt.Assign.Id] = object
				}
			} else {
				assign_Table(stmt.Assign.Id, symbol_env, result)
			}
		}

	case parser.DECL_STMT:
		{

			result := Eval_Assign(*stmt.Decl.Assign, symbol_env)
			symbol_env.Table[stmt.Decl.Id] = result
			//fmt.Printf("%+v", result)

		}

	case parser.EXPR:
		{

			switch stmt.Expr.Type.DataStructType {
			case basic_type.PRIMITIVE:
				{

					switch stmt.Expr.Type.Primitive.Type {
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
			case basic_type.ARRAY:
				{
					object := resolve_variable(stmt.Expr.Terms[0].Term.Factors[0].Factor.Id, symbol_env)
					fmt.Printf("[")

					for i, ele := range object.Array.Value {
						fmt.Printf("%+v", ele.Array.Value[0].Primitive)

						if i != len(object.Array.Value)-1 {
							fmt.Printf(", ")
						}
					}
					fmt.Printf("]")
				}
			}

		}
	}

	return false
}
