package parser

import (
	"fmt"
	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"testing"
)

func TestParser(t *testing.T) {

	fmt.Printf("\n------------------------------------------------------\n")

	regex_string := lexer.Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	input := "3+13.0"
	fmt.Printf("%s\n", input)

	tokens := lexer.GetTokens(regex, input)

	parser_input := Parser_Input{Tokens: tokens, Pos: 0}

	fmt.Printf("%+v\n", tokens)

	arith_expr := Parse_Arith_expr(&parser_input)

	fmt.Printf("%+v", arith_expr)

	fmt.Printf("\n------------------------------------------------------\n")

}
func TestCall(t *testing.T) {

	fmt.Printf("\n------------------------------------------------------\n")

	regex_string := lexer.Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	input := "add(a,b)"
	fmt.Printf("%s\n", input)

	tokens := lexer.GetTokens(regex, input)

	parser_input := Parser_Input{Tokens: tokens, Pos: 0}

	fmt.Printf("%+v\n", tokens)

	arith_expr := Parse_Arith_expr(&parser_input)

	fmt.Printf("%+v", arith_expr)

	fmt.Printf("\n------------------------------------------------------\n")

}

func TestEQUAL(t *testing.T) {

	fmt.Printf("\n------------------------------------------------------\n")

	regex_string := lexer.Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	input := "a == b"
	fmt.Printf("%s\n", input)

	tokens := lexer.GetTokens(regex, input)

	parser_input := Parser_Input{Tokens: tokens, Pos: 0}

	fmt.Printf("%+v\n", tokens)

	arith_expr := Parse_Cmp_expr(&parser_input)

	fmt.Printf("%+v", arith_expr)

	fmt.Printf("\n------------------------------------------------------\n")

}

func TestGT(t *testing.T) {

	fmt.Printf("\n------------------------------------------------------\n")

	regex_string := lexer.Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	input := "a > b"
	fmt.Printf("%s\n", input)

	tokens := lexer.GetTokens(regex, input)

	parser_input := Parser_Input{Tokens: tokens, Pos: 0}

	fmt.Printf("%+v\n", tokens)

	arith_expr := Parse_Cmp_expr(&parser_input)

	fmt.Printf("%+v", arith_expr)

	fmt.Printf("\n------------------------------------------------------\n")

}
