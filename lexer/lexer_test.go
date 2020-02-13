package lexer

import (
	"fmt"
	"github.com/KazumaTakata/regex_virtualmachine"
	"testing"
)

func TestLexer(t *testing.T) {

	regex_string := Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	tokens := GetTokens(regex, "  3 + 13.0  ")

	fmt.Printf("%+v", tokens)

	tokens = GetTokens(regex, "  \"stringdata\"   ")

	fmt.Printf("%+v\n", tokens)

	tokens = GetTokens(regex, "var x int = 33 \n")

	fmt.Printf("%+v\n", tokens)

	tokens = GetTokens(regex, "if (33){}")

	fmt.Printf("%+v\n", tokens)

}
