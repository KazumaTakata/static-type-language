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

func Type_Check_Stmt(stmt parser.Stmt, symbol_env *parser.Symbol_Env) {
	switch stmt.Type {
	case parser.EXPR:
		{
			_ = Type_Check_Arith(stmt.Expr, symbol_env)

		}
	case parser.DECL_STMT:
		{
			var_type := stmt.Decl.Type
			expr_type := Type_Check_Arith(stmt.Decl.Expr, symbol_env)

			variable := parser.Variable{Type: var_type}

			symbol_env.Table[stmt.Decl.Id] = parser.Object{Type: parser.VariableObj, Variable: &variable}

			if var_type != expr_type {
				fmt.Printf("%+v value can not assigned to %+v variable\n", expr_type, var_type)
				os.Exit(1)
			}

		}
	case parser.FOR_STMT:
		{
			_ = Type_Check_Cmp(&stmt.For.Cmp_expr, symbol_env)

			if stmt.For.Cmp_expr.Type != basic_type.BOOL {
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

			if stmt.If.Cmp_expr.Type != basic_type.BOOL {
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

			if symbol_env.Return_Type != stmt.Return.Type {
				fmt.Printf("func return type mismatch :expect %+v, got%+v\n", symbol_env.Return_Type, stmt.Return.Type)
				os.Exit(1)

			}

		}
	case parser.DEF_STMT:
		{

			function := parser.Object{Type: parser.FunctionObj, Function: stmt.Def}
			symbol_env.Table[stmt.Def.Id] = function

			Child_env := &parser.Symbol_Env{Table: parser.Symbol_Table{}, Parent_Env: symbol_env}

			for _, arg := range stmt.Def.Args {
				Child_env.Table[arg.Ident] = parser.Object{Type: parser.VariableObj, Variable: &parser.Variable{Type: arg.Type}}
			}

			Child_env.Return_Type = function.Function.Return_type
			stmt.Def.Symbol_Env = Child_env
			Type_Check_Stmts(stmt.Def.Stmts, Child_env)

		}

	}
}
