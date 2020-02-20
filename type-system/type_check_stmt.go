package type_system

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/parser"
	"github.com/KazumaTakata/static-typed-language/type"
	"os"
)

func Type_Check_Stmts(stmts []parser.Stmt, symbol_env *parser.Symbol_Env) {

	for _, stmt := range stmts {
		Type_Check_Stmt(stmt, symbol_env)
	}

}

func Type_Check_Assign(assign *parser.Assign, symbol_env *parser.Symbol_Env) basic_type.Variable_Type {

	if assign.Type == parser.INIT_ASSIGN {
		switch assign.Init.Type {
		case parser.ARRAY_INIT:
			{
				array_type := assign.Init.Array.ElementType
				for _, init_value := range assign.Init.Array.InitValue {
					cmp_type := Type_Check_Cmp(init_value, symbol_env)
					if array_type.DataStructType != cmp_type.DataStructType {
						fmt.Printf("array type %+v  mismatch to initialization type %+v\n", array_type, cmp_type)
						os.Exit(1)
					}
				}

				return basic_type.Variable_Type{DataStructType: basic_type.ARRAY, Array: &basic_type.ArrayType{ElementType: array_type}}
			}
		case parser.MAP_INIT:
			{
			}

		}
	} else if assign.Type == parser.EXPR_ASSIGN {

		cmp_type := Type_Check_Cmp(assign.Expr, symbol_env)

		return cmp_type
	}

	return basic_type.Variable_Type{}

}

func Type_Check_Stmt(stmt parser.Stmt, symbol_env *parser.Symbol_Env) {
	switch stmt.Type {
	case parser.EXPR:
		{
			_ = Type_Check_Arith(stmt.Expr, symbol_env)

		}

	case parser.ASSIGN_STMT:
		{
			variable_type := Type_Check_Assign(stmt.Assign.Assign, symbol_env)
			object := resolve_name(stmt.Assign.Id, symbol_env)

			switch variable_type.DataStructType {
			case basic_type.ARRAY:
				{
					if object.Type != parser.ArrayType {
						fmt.Printf("array value can not assigned to %+v  varieble\n", object.Type)
						os.Exit(1)
					}

					if object.Array.ElementType.Primitive.Type != variable_type.Array.ElementType.Primitive.Type {
						fmt.Printf("array of type %+v can not assigned to type of %+v\n", variable_type.DataStructType, object.Array.ElementType)
						os.Exit(1)

					}

				}

			case basic_type.PRIMITIVE:
				{
					if object.Type != parser.PrimitiveType {
						fmt.Printf("primitive value can not assigned to %+v  varieble\n", object.Type)
						os.Exit(1)
					}

					if object.Primitive.Type != variable_type.Primitive.Type {
						fmt.Printf("primitive of type %+v  can not assigned to type of %+v\n", variable_type.DataStructType, object.Primitive.Type)
						os.Exit(1)

					}
				}
			}
		}

	case parser.DECL_STMT:
		{
			var_type := stmt.Decl.Type
			assign_type := Type_Check_Assign(stmt.Decl.Assign, symbol_env)

			if var_type.DataStructType != assign_type.DataStructType {
				fmt.Printf("data structure mismatch %+v: %+v\n", var_type.DataStructType, assign_type.DataStructType)
				os.Exit(1)

			}

			switch var_type.DataStructType {
			case basic_type.PRIMITIVE:
				{
					primitive := parser.PrimitiveObj{Type: var_type.Primitive.Type}
					symbol_env.Table[stmt.Decl.Id] = parser.Object{Type: parser.PrimitiveType, Primitive: &primitive}
				}
			case basic_type.ARRAY:
				{
					array := parser.ArrayObj{ElementType: var_type.Array.ElementType}
					symbol_env.Table[stmt.Decl.Id] = parser.Object{Type: parser.ArrayType, Array: &array}
				}

			}

		}
	case parser.FOR_STMT:
		{
			_ = Type_Check_Cmp(&stmt.For.Cmp_expr, symbol_env)

			if stmt.For.Cmp_expr.Type.DataStructType != basic_type.PRIMITIVE {
				fmt.Printf("if conditional expression should be primitive: got %+v\n", stmt.For.Cmp_expr.Type)
				os.Exit(1)
			}
			if stmt.For.Cmp_expr.Type.Primitive.Type != basic_type.BOOL {
				fmt.Printf("if conditional expression should return bool type: return %+v\n", stmt.For.Cmp_expr.Type)
				os.Exit(1)
			}

			Child_env := &parser.Symbol_Env{Table: parser.Symbol_Table{}, Parent_Env: symbol_env}
			stmt.For.Symbol_Env = Child_env
			Type_Check_Stmts(stmt.For.Stmts, Child_env)

		}
	case parser.IF_STMT:
		{
			_ = Type_Check_Cmp(&stmt.If.Cmp_expr, symbol_env)

			if stmt.If.Cmp_expr.Type.DataStructType != basic_type.PRIMITIVE {
				fmt.Printf("if conditional expression should be primitive: got %+v\n", stmt.If.Cmp_expr.Type)
				os.Exit(1)
			}
			if stmt.If.Cmp_expr.Type.Primitive.Type != basic_type.BOOL {
				fmt.Printf("if conditional expression should return bool type: return %+v\n", stmt.If.Cmp_expr.Type)
				os.Exit(1)
			}

			Child_env := &parser.Symbol_Env{Table: parser.Symbol_Table{}, Parent_Env: symbol_env}
			stmt.If.Symbol_Env = Child_env
			Type_Check_Stmts(stmt.If.Stmts, Child_env)

		}

	case parser.RETURN_STMT:
		{
			//_ = Type_Check_Arith(&stmt.Return.Cmp_expr.Ariths[0].Arith, symbol_env)
			_ = Type_Check_Cmp(&stmt.Return.Cmp_expr, symbol_env)

			stmt.Return.Type = stmt.Return.Cmp_expr.Type

			if !basic_type.Variable_Equal(symbol_env.Return_Type, stmt.Return.Type) {
				fmt.Printf("func return type mismatch :expect %+v, got%+v\n", symbol_env.Return_Type.Primitive, stmt.Return.Type.Primitive)
				os.Exit(1)

			}

		}
	case parser.DEF_STMT:
		{

			function := parser.Object{Type: parser.FunctionType, Function: stmt.Def}
			symbol_env.Table[stmt.Def.Id] = function

			Child_env := &parser.Symbol_Env{Table: parser.Symbol_Table{}, Parent_Env: symbol_env}

			for _, arg := range stmt.Def.Args {

				if arg.Type.DataStructType != basic_type.PRIMITIVE {
					fmt.Printf("def argument should be primitive: got %+v\n", arg.Type.DataStructType)
					os.Exit(1)
				}

				Child_env.Table[arg.Ident] = parser.Object{Type: parser.PrimitiveType, Primitive: &parser.PrimitiveObj{Type: arg.Type.Primitive.Type}}
			}

			Child_env.Return_Type = function.Function.Return_type
			stmt.Def.Symbol_Env = Child_env
			Type_Check_Stmts(stmt.Def.Stmts, Child_env)

		}

	}
}
