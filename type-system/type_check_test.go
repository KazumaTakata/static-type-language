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

	string_input := "var x int  = 3 + 3"

	tokens := lexer.GetTokens(regex, string_input)
	parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}
	stmt := parser.Parse_Stmt(&parser_input)

	resolved_type, variable_type_map := Type_Check_Stmt(stmt)

	fmt.Printf("%+v, %+v\n", resolved_type, variable_type_map)

}
