package optimize

import (
	//	"fmt"
	"fmt"
	"github.com/KazumaTakata/static-typed-language/IR-gen"
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

	codes, _ := ir_gen.Gen_IR_Arith(*stmts[0].Expr)

	//for _, available := range availables {
	//fmt.Printf("{")
	//for _, code := range available {
	//fmt.Printf(" %s, ", code.String())
	//}
	//fmt.Printf("}\n")

	/*}*/

}
