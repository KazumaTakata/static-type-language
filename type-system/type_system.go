package main

import (
	"fmt"
	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"
	"strings"
)

type Type int

const (
	Int    Type = 0
	Double Type = 1
)

func Type_Check_Arith_Term(term1 parser.Term, term2 parser.Term) {

}

func Type_Check_Arith_Factor(factor1 parser.Factor, factor2 parser.Factor) {

}

func main() {

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

	input := "13.3+33"
	fmt.Printf("%s\n", input)

	tokens := lexer.GetTokens(regex, input)

	parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}

	//fmt.Printf("%+v", tokens)

	arith_expr := parser.Parse_Arith_expr(&parser_input)

	fmt.Printf("%+v", arith_expr)

}
