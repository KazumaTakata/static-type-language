package type_checker

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/parser"
	"github.com/KazumaTakata/static-typed-language/type"
	"os"
)

type Variable_Table map[string]Variable

type Variable struct {
	Type   basic_type.Type
	Int    int
	Double float64
	String string
}

type Type_Env struct {
	Table     Variable_Table
	Child_Env *Type_Env
}

func Type_Check_Stmts(stmts []parser.Stmt, variable_map Variable_Table) {

	for _, stmt := range stmts {
		Type_Check_Stmt(stmt, variable_map)
	}

}
func Type_Check_Stmt(stmt parser.Stmt, variable_map Variable_Table) {
	switch stmt.Type {
	case parser.EXPR:
		{
			_ = Type_Check_Arith(stmt.Expr, variable_map)

		}
	case parser.DECL_STMT:
		{
			var_type := stmt.Decl.Type
			expr_type := Type_Check_Arith(stmt.Decl.Expr, variable_map)

			variable_map[stmt.Decl.Id] = Variable{Type: var_type}

			if var_type != expr_type {
				fmt.Printf("%+v value can not assigned to %+v variable\n", expr_type, var_type)
				os.Exit(1)
			}

		}
	case parser.FOR_STMT:
		{
			_ = Type_Check_Cmp(&stmt.For.Cmp_expr, variable_map)

			if stmt.For.Cmp_expr.Type != basic_type.BOOL {
				fmt.Printf("if conditional expression should return bool type: return %+v\n", stmt.For.Cmp_expr.Type)
				os.Exit(1)
			}

			Type_Check_Stmts(stmt.For.Stmts, variable_map)

		}
	case parser.IF_STMT:
		{
			_ = Type_Check_Cmp(&stmt.For.Cmp_expr, variable_map)

			if stmt.For.Cmp_expr.Type != basic_type.BOOL {
				fmt.Printf("if conditional expression should return bool type: return %+v\n", stmt.For.Cmp_expr.Type)
				os.Exit(1)
			}

			Type_Check_Stmts(stmt.For.Stmts, variable_map)

		}

	}
}
