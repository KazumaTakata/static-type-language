package main

import (
	"fmt"
	"github.com/KazumaTakata/readline"
	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/object"
	"github.com/KazumaTakata/static-typed-language/parser"
	"github.com/KazumaTakata/static-typed-language/tree-eval"
	"github.com/KazumaTakata/static-typed-language/type-system"

	"io/ioutil"
	"os"
)

func getClosure() func([]byte) {

	regex_string := lexer.Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	symbol_env := object.Symbol_Env{Table: object.Variable_Table{}}

	return func(input []byte) {
		string_input := string(input)
		tokens := lexer.GetTokens(regex, string_input)
		parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}
		stmt := parser.Parse_Stmt(&parser_input)

		type_system.Type_Check_Stmt(stmt, &symbol_env)

		eval.Eval_Stmt(stmt, &symbol_env)
	}
}

func run_program(input []byte) {
	regex_string := lexer.Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	symbol_env := object.Symbol_Env{Table: object.Variable_Table{}}

	string_input := string(input)
	tokens := lexer.GetTokens(regex, string_input)
	parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}
	stmt := parser.Parse_Stmts(&parser_input)

	type_system.Type_Check_Stmts(stmt, &symbol_env)

	eval.Eval_Stmts(stmt, &symbol_env)

}

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 0 {
		dat, _ := ioutil.ReadFile(argsWithoutProg[0])
		fmt.Printf("%s", dat)

		run_program(dat)

	} else {

		closure := getClosure()

		readline.Readline(">>", closure)

	}

}
