package ir_gen

import (
	"github.com/KazumaTakata/static-typed-language/parser"
)

func Gen_IR_Stmts(stmts []parser.Stmt, symbol_env *parser.Symbol_Env) {

	for _, stmt := range stmts {
		Gen_IR_Stmt(stmt, symbol_env)
	}
}

func Gen_IR_Stmt(stmt parser.Stmt, symbol_env *parser.Symbol_Env) {

	switch stmt.Type {

	case parser.DECL_STMT:
		{

		}
	case parser.ASSIGN_STMT:
		{

		}

	}

}

func Gen_IR_Assign(assign parser.Assign, symbol_env *parser.Symbol_Env) {

	switch assign.Type {
	case parser.EXPR_ASSIGN:
		{

		}
	}

}
