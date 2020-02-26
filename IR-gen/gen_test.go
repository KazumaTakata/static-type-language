package ir_gen

import (
	//	"fmt"
	"fmt"
	"testing"

	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"

	"io/ioutil"
)

func TestLexer(t *testing.T) {

	dat, _ := ioutil.ReadFile("test.cat")
	run_program(dat)

}

func run_program(input []byte) {
	regex_string := lexer.Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	string_input := string(input)
	tokens := lexer.GetTokens(regex, string_input)
	parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}
	stmts := parser.Parse_Stmts(&parser_input)

	codes := Gen_IR_Stmts(stmts)

	for _, code := range codes {
		fmt.Printf("%s\n", code.String())
	}

}
