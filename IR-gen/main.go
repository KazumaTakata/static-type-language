package main

import (
	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"
	"github.com/KazumaTakata/static-typed-language/tree-eval"
	"github.com/KazumaTakata/static-typed-language/type-system"
	"io/ioutil"
	"os"
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

func Gen_IR_Expr(cmp parser.Cmp_expr, symbol_env *parser.Symbol_Env) {
	Gen_IR_Arith(cmp.Left, symbol_env)
}

func Gen_IR_Arith(arith parser.Arith_expr, symbol_env *parser.Symbol_Env) {

	Gen_IR_Term(arith.Terms[0].Term, symbol_env)

}

func Gen_IR_Term(term parser.Term, symbol_env *parser.Symbol_Env) {
	term.Factors[0].Factor
}

func run_program(input []byte) {
	regex_string := lexer.Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	symbol_env := parser.Symbol_Env{Table: parser.Symbol_Table{}}

	string_input := string(input)
	tokens := lexer.GetTokens(regex, string_input)
	parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}
	stmt := parser.Parse_Stmts(&parser_input)

	type_system.Type_Check_Stmts(stmt, &symbol_env)

	eval.Eval_Stmts(stmt, &symbol_env)

}

func main() {

	argsWithoutProg := os.Args[1:]

	dat, _ := ioutil.ReadFile(argsWithoutProg[0])

	run_program(dat)

}
