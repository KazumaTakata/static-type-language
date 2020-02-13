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

func Type_Check_Stmt(stmt parser.Stmt) (basic_type.Type, Variable_Table) {

	variable_map := Variable_Table{}

	if stmt.Type == parser.EXPR {
		expr_type := Type_Check_Arith(stmt.Expr, variable_map)
		return expr_type, variable_map
	} else if stmt.Type == parser.DECL_STMT {
		var_type := stmt.Decl.Type
		expr_type := Type_Check_Arith(stmt.Decl.Expr, variable_map)

		variable_map[stmt.Decl.Id] = Variable{Type: var_type}

		if var_type != expr_type {
			fmt.Printf("%+v value can not assigned to %+v variable\n", expr_type, var_type)
			os.Exit(1)
		}

		return var_type, variable_map
	}

	return basic_type.INT, variable_map
}
