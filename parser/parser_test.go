package parser

import (
	"fmt"
	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {

	lexer_rules := [][]string{}
	lexer_rules = append(lexer_rules, []string{"DOUBLE", "\\d+\\.\\d*"})
	lexer_rules = append(lexer_rules, []string{"INT", "\\d+"})
	lexer_rules = append(lexer_rules, []string{"STRING", "\"\\w*\""})
	lexer_rules = append(lexer_rules, []string{"ADD", "\\+"})
	lexer_rules = append(lexer_rules, []string{"SUB", "\\-"})
	lexer_rules = append(lexer_rules, []string{"MUL", "\\*"})
	lexer_rules = append(lexer_rules, []string{"DIV", "\\/"})

	regex_parts := []string{}

	for _, rule := range lexer_rules {
		regex_parts = append(regex_parts, fmt.Sprintf("(?<%s>%s)", rule[0], rule[1]))
	}

	regex_string := strings.Join(regex_parts, "|")
	//fmt.Printf("%s", regex_string)

	regex := regex.NewRegexWithParser(regex_string)

	input := "3+13.0"
	fmt.Printf("%s\n", input)

	tokens := lexer.GetTokens(regex, input)

	parser_input := Parser_Input{Tokens: tokens, Pos: 0}

	fmt.Printf("%+v\n", tokens)

	arith_expr := Parse_Arith_expr(&parser_input)

	fmt.Printf("%+v", arith_expr)
}
