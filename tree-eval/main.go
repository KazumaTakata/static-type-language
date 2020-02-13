package main

import (
	"fmt"
	"github.com/KazumaTakata/readline"
	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"
	basic_type "github.com/KazumaTakata/static-typed-language/type"
	type_checker "github.com/KazumaTakata/static-typed-language/type-system"
	"os"
)

func Arith_Factors_INT(factors []parser.TermElement, variable_map type_checker.Variable_Table) int {

	if len(factors) == 1 {

		if factors[0].Factor.Type == lexer.IDENT {
			return variable_map[factors[0].Factor.Id].Int
		}
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

func Arith_Factors_STRING(factors []parser.TermElement, variable_map type_checker.Variable_Table) string {

	if len(factors) == 1 {

		if factors[0].Factor.Type == lexer.IDENT {
			return variable_map[factors[0].Factor.Id].String
		}

		return factors[0].Factor.String
	}
	os.Exit(1)
	return ""
}

func Arith_Factors_DOUBLE(factors []parser.TermElement, variable_map type_checker.Variable_Table) float64 {

	if len(factors) == 1 {

		if factors[0].Factor.Type == lexer.IDENT {
			if variable_map[factors[0].Factor.Id].Type == basic_type.DOUBLE {
				return variable_map[factors[0].Factor.Id].Double
			}
		}

		return factors[0].Factor.Float
	}

	var result float64

	for i, factor := range factors {
		if i == 0 {

			switch factor.Factor.Type {

			case lexer.IDENT:
				{
					if variable_map[factor.Factor.Id].Type == basic_type.DOUBLE {
						result = variable_map[factor.Factor.Id].Double
					} else if variable_map[factor.Factor.Id].Type == basic_type.INT {
						result = float64(variable_map[factor.Factor.Id].Int)
					}

				}

			case lexer.INT:
				{
					result = float64(factor.Factor.Int)
				}
			case lexer.DOUBLE:
				{
					result = factor.Factor.Float
				}
			}
			continue
		}
		switch factor.Op {
		case parser.MUL:
			{
				switch factor.Factor.Type {
				case lexer.IDENT:
					{
						if variable_map[factor.Factor.Id].Type == basic_type.DOUBLE {
							result = result * variable_map[factor.Factor.Id].Double
						} else if variable_map[factor.Factor.Id].Type == basic_type.INT {
							result = result * float64(variable_map[factor.Factor.Id].Int)
						}

					}

				case lexer.INT:
					{
						result = result * float64(factor.Factor.Int)
					}
				case lexer.DOUBLE:
					{
						result = result * factor.Factor.Float
					}
				}
			}

		case parser.DIV:
			{

				switch factor.Factor.Type {
				case lexer.IDENT:
					{
						if variable_map[factor.Factor.Id].Type == basic_type.DOUBLE {
							result = result / variable_map[factor.Factor.Id].Double
						} else if variable_map[factor.Factor.Id].Type == basic_type.INT {
							result = result / float64(variable_map[factor.Factor.Id].Int)
						}

					}

				case lexer.INT:
					{
						result = result / float64(factor.Factor.Int)
					}
				case lexer.DOUBLE:
					{
						result = result / factor.Factor.Float
					}
				}
			}
		}

	}
	return result
}

func Arith_Terms_INT(terms []parser.ArithElement, variable_map type_checker.Variable_Table) int {

	if len(terms) == 1 {
		return Arith_Factors_INT(terms[0].Term.Factors, variable_map)
	}

	var result int

	for i, term := range terms {
		if i == 0 {
			result = Arith_Factors_INT(term.Term.Factors, variable_map)
			continue
		}
		switch term.Op {
		case parser.ADD:
			{
				result = result + Arith_Factors_INT(term.Term.Factors, variable_map)
			}

		case parser.SUB:
			{
				result = result - Arith_Factors_INT(term.Term.Factors, variable_map)
			}
		}

	}
	return result
}

func Arith_Terms_DOUBLE(terms []parser.ArithElement, variable_map type_checker.Variable_Table) float64 {

	if len(terms) == 1 {
		return Arith_Factors_DOUBLE(terms[0].Term.Factors, variable_map)
	}

	var result float64

	for i, term := range terms {
		//		fmt.Printf("%+v\n", term)
		if i == 0 {
			if term.Term.Type == basic_type.INT {
				result = float64(Arith_Factors_INT(term.Term.Factors, variable_map))
			} else if term.Term.Type == basic_type.DOUBLE {
				result = Arith_Factors_DOUBLE(term.Term.Factors, variable_map)
			}
			continue
		}
		switch term.Op {
		case parser.ADD:
			{
				if term.Term.Type == basic_type.INT {
					result = result + float64(Arith_Factors_INT(term.Term.Factors, variable_map))
				} else if term.Term.Type == basic_type.DOUBLE {
					result = result + Arith_Factors_DOUBLE(term.Term.Factors, variable_map)
				}
			}

		case parser.SUB:
			{
				if term.Term.Type == basic_type.INT {
					result = result - float64(Arith_Factors_INT(term.Term.Factors, variable_map))
				} else if term.Term.Type == basic_type.DOUBLE {
					result = result - Arith_Factors_DOUBLE(term.Term.Factors, variable_map)
				}
			}
		}

	}
	return result
}

func Arith_Terms_STRING(terms []parser.ArithElement, variable_map type_checker.Variable_Table) string {

	if len(terms) == 1 {
		return Arith_Factors_STRING(terms[0].Term.Factors, variable_map)
	}

	var result string

	for i, term := range terms {
		//		fmt.Printf("%+v\n", term)
		if i == 0 {
			result = Arith_Factors_STRING(term.Term.Factors, variable_map)
			continue
		}
		switch term.Op {
		case parser.ADD:
			{
				result = result + Arith_Factors_STRING(term.Term.Factors, variable_map)
			}

		case parser.SUB:
			{

			}
		}

	}
	return result
}

func Eval_Stmt(stmt parser.Stmt, variable_map type_checker.Variable_Table) {

	switch stmt.Type {
	case parser.DECL_STMT:
		{

			switch stmt.Decl.Expr.Type {
			case basic_type.INT:
				{
					result := Arith_Terms_INT(stmt.Decl.Expr.Terms, variable_map)
					variable_map[stmt.Decl.Id] = type_checker.Variable{Int: result, Type: basic_type.INT}

					fmt.Printf("%+v", result)

				}
			case basic_type.DOUBLE:
				{
					result := Arith_Terms_DOUBLE(stmt.Decl.Expr.Terms, variable_map)
					variable_map[stmt.Decl.Id] = type_checker.Variable{Double: result, Type: basic_type.DOUBLE}

					fmt.Printf("%+v", result)

				}
			case basic_type.STRING:
				{
					result := Arith_Terms_STRING(stmt.Decl.Expr.Terms, variable_map)
					variable_map[stmt.Decl.Id] = type_checker.Variable{String: result, Type: basic_type.STRING}

					fmt.Printf("%+v", result)

				}
			}

		}

	case parser.EXPR:
		{

			switch stmt.Expr.Type {
			case basic_type.INT:
				{
					result := Arith_Terms_INT(stmt.Expr.Terms, variable_map)
					fmt.Printf("%+v", result)

				}
			case basic_type.DOUBLE:
				{
					result := Arith_Terms_DOUBLE(stmt.Expr.Terms, variable_map)

					fmt.Printf("%+v", result)

				}
			case basic_type.STRING:
				{
					result := Arith_Terms_STRING(stmt.Expr.Terms, variable_map)

					fmt.Printf("%+v", result)

				}
			}
		}

	}
}

func getClosure() func([]byte) {

	regex_string := lexer.Get_Regex_String()

	regex := regex.NewRegexWithParser(regex_string)

	variable_table := type_checker.Variable_Table{}

	return func(input []byte) {
		string_input := string(input)
		tokens := lexer.GetTokens(regex, string_input)
		parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}
		stmt := parser.Parse_Stmt(&parser_input)

		_ = type_checker.Type_Check_Stmt(stmt, variable_table)

		Eval_Stmt(stmt, variable_table)
	}
}

func main() {

	closure := getClosure()

	readline.Readline(">>", closure)

}
