package type_system

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/object"
	"github.com/KazumaTakata/static-typed-language/parser"

	"github.com/KazumaTakata/static-typed-language/type"
	"os"
)

func Type_Check_Stmts(stmts []parser.Stmt, symbol_env *object.Symbol_Env) {

	for _, stmt := range stmts {
		Type_Check_Stmt(stmt, symbol_env)
	}

}

func Type_Check_Stmt(stmt parser.Stmt, symbol_env *object.Symbol_Env) {
	switch stmt.Type {
	case parser.EXPR:
		{
			_ = Type_Check_Arith(stmt.Expr, symbol_env)

		}
	case parser.DECL_STMT:
		{
			var_type := stmt.Decl.Type
			expr_type := Type_Check_Arith(stmt.Decl.Expr, symbol_env)

			symbol_env.Table[stmt.Decl.Id] = object.Variable{Type: var_type}

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

			Child_env := &object.Symbol_Env{Table: object.Variable_Table{}, Parent_Env: symbol_env}
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

			Child_env := &object.Symbol_Env{Table: object.Variable_Table{}, Parent_Env: symbol_env}
			stmt.If.Symbol_Env = Child_env
			Type_Check_Stmts(stmt.If.Stmts, Child_env)

		}

	}
}
