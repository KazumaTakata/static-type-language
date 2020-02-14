package type_checker

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

	string_input := "var x int  = 3 + 3\n "

	tokens := lexer.GetTokens(regex, string_input)
	parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}
	stmts := parser.Parse_Stmts(&parser_input)

	variable_map := Variable_Table{}

	Type_Check_Stmts(stmts, variable_map)

	fmt.Printf("%+v\n", variable_map)

}

func TestFor(t *testing.T) {

	regex_string := lexer.Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	string_input := "for 1 == 1 { var x double = 1 } "

	tokens := lexer.GetTokens(regex, string_input)
	parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}
	stmts := parser.Parse_Stmts(&parser_input)

	variable_map := Variable_Table{}

	Type_Check_Stmts(stmts, variable_map)

	fmt.Printf("%+v\n", variable_map)

}
