package main

import (
	"fmt"
	"strings"

	"github.com/KazumaTakata/readline"
	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"
	basic_type "github.com/KazumaTakata/static-typed-language/type"
	type_checker "github.com/KazumaTakata/static-typed-language/type-system"
)

func Arith_Factors_INT(factors []parser.TermElement) int {

	if len(factors) == 1 {
		return factors[0].Factor.Int
	}

	var result int

	for i, factor := range factors {
		if i == 0 {
			result = factor.Factor.Int
			continue
		}
		switch factor.Op {
		case parser.MUL:
			{
				result = result * factor.Factor.Int
			}

		case parser.DIV:
			{
				result = result / factor.Factor.Int
			}
		}

	}
	return result
}

func Arith_Factors_DOUBLE(factors []parser.TermElement) float64 {

	if len(factors) == 1 {
		return factors[0].Factor.Float
	}

	var result float64

	for i, factor := range factors {
		if i == 0 {
			if factor.Factor.Type == lexer.INT {
				result = float64(factor.Factor.Int)
			} else if factor.Factor.Type == lexer.DOUBLE {
				result = factor.Factor.Float
			}
			continue
		}
		switch factor.Op {
		case parser.MUL:
			{
				if factor.Factor.Type == lexer.INT {
					result = result * float64(factor.Factor.Int)
				} else if factor.Factor.Type == lexer.DOUBLE {
					result = result * factor.Factor.Float
				}
			}

		case parser.DIV:
			{
				if factor.Factor.Type == lexer.INT {
					result = result / float64(factor.Factor.Int)
				} else if factor.Factor.Type == lexer.DOUBLE {
					result = result / factor.Factor.Float
				}
			}
		}

	}
	return result
}

func Arith_Terms_INT(terms []parser.ArithElement) int {

	if len(terms) == 1 {
		return Arith_Factors_INT(terms[0].Term.Factors)
	}

	var result int

	for i, term := range terms {
		if i == 0 {
			result = Arith_Factors_INT(term.Term.Factors)
			continue
		}
		switch term.Op {
		case parser.ADD:
			{
				result = result + Arith_Factors_INT(term.Term.Factors)
			}

		case parser.SUB:
			{
				result = result - Arith_Factors_INT(term.Term.Factors)
			}
		}

	}
	return result
}

func Arith_Terms_DOUBLE(terms []parser.ArithElement) float64 {

	if len(terms) == 1 {
		return Arith_Factors_DOUBLE(terms[0].Term.Factors)
	}

	var result float64

	for i, term := range terms {
		//		fmt.Printf("%+v\n", term)
		if i == 0 {
			if term.Term.Type == basic_type.INT {
				result = float64(Arith_Factors_INT(term.Term.Factors))
			} else if term.Term.Type == basic_type.DOUBLE {
				result = Arith_Factors_DOUBLE(term.Term.Factors)
			}
			continue
		}
		switch term.Op {
		case parser.ADD:
			{
				if term.Term.Type == basic_type.INT {
					result = result + float64(Arith_Factors_INT(term.Term.Factors))
				} else if term.Term.Type == basic_type.DOUBLE {
					result = result + Arith_Factors_DOUBLE(term.Term.Factors)
				}
			}

		case parser.SUB:
			{
				if term.Term.Type == basic_type.INT {
					result = result - float64(Arith_Factors_INT(term.Term.Factors))
				} else if term.Term.Type == basic_type.DOUBLE {
					result = result - Arith_Factors_DOUBLE(term.Term.Factors)
				}
			}
		}

	}
	return result
}

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

		resolved_type := type_checker.Type_Check_Arith(&arith_expr)

		if resolved_type == basic_type.INT {
			result := Arith_Terms_INT(arith_expr.Terms)
			fmt.Printf("%+v", result)
		} else if resolved_type == basic_type.DOUBLE {
			result := Arith_Terms_DOUBLE(arith_expr.Terms)
			fmt.Printf("%+v", result)

		}

	}
}

func main() {

	closure := getClosure()

	readline.Readline(">>", closure)

}
