package main

import (
	"fmt"
	"strings"

	"github.com/KazumaTakata/readline"
	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"
	type_checker "github.com/KazumaTakata/static-typed-language/type-system"
)

func getClosure() func([]byte) {

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

	return func(input []byte) {
		string_input := string(input)
		tokens := lexer.GetTokens(regex, string_input)
		parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}
		arith_expr := parser.Parse_Arith_expr(&parser_input)

		//fmt.Printf("%+v\n", arith_expr)
		resolved_type := type_checker.Type_Check_Arith(&arith_expr)
		fmt.Printf("\n%+v\n", resolved_type)
		fmt.Printf("\n%+v\n", arith_expr.Type)

	}
}

func main() {

	closure := getClosure()

	readline.Readline(">>", closure)

}
