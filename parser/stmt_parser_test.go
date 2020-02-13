package parser

import (
	"fmt"
	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"testing"
)

func TestStmtParser(t *testing.T) {

	regex_string := lexer.Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	input := "var x int = 3+13.0"
	fmt.Printf("%s\n", input)

	tokens := lexer.GetTokens(regex, input)

	parser_input := Parser_Input{Tokens: tokens, Pos: 0}

	fmt.Printf("%+v\n\n", tokens)

	stmt := Parse_Stmt(&parser_input)

	fmt.Printf("%+v", stmt)
}
