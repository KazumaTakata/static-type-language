package main

import (
	"fmt"
	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"
	"os"
	"strings"
)

type Type int

const (
	INT Type = iota + 1
	DOUBLE
	STRING
)

func (e Type) String() string {

	switch e {
	case INT:
		return "INT"
	case DOUBLE:
		return "DOUBLE"
	case STRING:
		return "STRING"

	default:
		return fmt.Sprintf("%d", int(e))
	}
}

var lexerTypeToType = map[lexer.TokenType]Type{lexer.INT: INT, lexer.DOUBLE: DOUBLE}
var NumberType = map[Type]bool{INT: true, DOUBLE: true}

func Type_Check_Arith_Factor(factor1_type Type, factor2_type Type, op parser.FactorOp) Type {

	is_factor1_number := false
	is_factor2_number := false

	if _, ok := NumberType[factor1_type]; ok {
		is_factor1_number = true
	}

	if _, ok := NumberType[factor2_type]; ok {
		is_factor2_number = true
	}

	if is_factor1_number && is_factor2_number {
		if factor1_type < factor2_type {
			return factor2_type
		} else {
			return factor1_type
		}
	} else if is_factor1_number && !is_factor2_number {

		fmt.Printf("\ntype mismatch: %+v can not be %ved with %v\n", factor1_type, op, factor2_type)
		os.Exit(1)
	} else if !is_factor1_number && is_factor2_number {
		fmt.Printf("\ntype mismatch: %+v can not be %ved with %v\n", factor1_type, op, factor2_type)
		os.Exit(1)
	} else {
		fmt.Printf("\ntype mismatch: %v can not be %ved with %v\n", factor1_type, op, factor2_type)
		os.Exit(1)

	}

	return INT

}

func Type_Check_Arith_Term(term1_type Type, term2_type Type, op parser.TermOp) Type {

	is_factor1_number := false
	is_factor2_number := false

	if _, ok := NumberType[term1_type]; ok {
		is_factor1_number = true
	}

	if _, ok := NumberType[term2_type]; ok {
		is_factor2_number = true
	}

	if is_factor1_number && is_factor2_number {
		if term1_type < term2_type {
			return term2_type
		} else {
			return term1_type
		}
	} else if is_factor1_number && !is_factor2_number {

		fmt.Printf("\ntype mismatch: %+v can not be %ved with %v\n", term1_type, op, term2_type)
		os.Exit(1)
	} else if !is_factor1_number && is_factor2_number {
		fmt.Printf("\ntype mismatch: %+v can not be %ved with %v\n", term1_type, op, term2_type)
		os.Exit(1)
	} else {

		if term1_type == term2_type && term1_type == STRING && op == parser.ADD {
			return STRING
		} else {
			fmt.Printf("\ntype mismatch: %v can not be %ved with %v\n", term1_type, op, term2_type)
			os.Exit(1)
		}

	}

	return INT

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

	input := "13+\"hello\""
	fmt.Printf("%s\n", input)

	tokens := lexer.GetTokens(regex, input)

	parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}

	//fmt.Printf("%+v", tokens)

	arith_expr := parser.Parse_Arith_expr(&parser_input)

	fmt.Printf("%+v", arith_expr)

	_ = Type_Check_Arith_Factor(STRING, INT, parser.MUL)

	factorType := Type_Check_Arith_Term(STRING, INT, parser.ADD)
	fmt.Printf("%+v\n", factorType)

}
