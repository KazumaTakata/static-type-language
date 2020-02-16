package type_system

import (
	"fmt"
	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"

	"testing"
)

func TestLexer(t *testing.T) {

	regex_string := lexer.Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	string_input := "var x int =1 \n "

	tokens := lexer.GetTokens(regex, string_input)
	parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}
	stmts := parser.Parse_Stmts(&parser_input)

	variable_map := parser.Symbol_Env{Table: parser.Variable_Table{}}

	Type_Check_Stmts(stmts, &variable_map)

	fmt.Printf("%+v\n", variable_map)

}

func TestFor(t *testing.T) {

	regex_string := lexer.Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	string_input := "var x int = 3\n if 1 == 1 { var y int = 1 + x  }\n"

	tokens := lexer.GetTokens(regex, string_input)
	parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}
	stmts := parser.Parse_Stmts(&parser_input)

	variable_map := parser.Symbol_Env{Table: parser.Variable_Table{}}

	Type_Check_Stmts(stmts, &variable_map)

	fmt.Printf("%+v\n", variable_map)

	fmt.Printf("%+v\n", stmts[1].If.Symbol_Env)

}
