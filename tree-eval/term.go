package eval

import (
	"github.com/KazumaTakata/static-typed-language/parser"
	basic_type "github.com/KazumaTakata/static-typed-language/type"
	"os"
)

func Arith_Terms_INT(terms []parser.ArithElement, symbol_env *parser.Symbol_Env) int {

	if len(terms) == 1 {
		return Arith_Factors_INT(terms[0].Term.Factors, symbol_env)
	}

	var result int

	for i, term := range terms {
		if i == 0 {
			result = Arith_Factors_INT(term.Term.Factors, symbol_env)
			continue
		}
		switch term.Op {
		case parser.ADD:
			{
				result = result + Arith_Factors_INT(term.Term.Factors, symbol_env)
			}

		case parser.SUB:
			{
				result = result - Arith_Factors_INT(term.Term.Factors, symbol_env)
			}
		}

	}
	return result
}

func Arith_Terms_DOUBLE(terms []parser.ArithElement, symbol_env *parser.Symbol_Env) float64 {

	if len(terms) == 1 {
		return Arith_Factors_DOUBLE(terms[0].Term.Factors, symbol_env)
	}

	var result float64

	for i, term := range terms {
		//		fmt.Printf("%+v\n", term)
		if i == 0 {
			if term.Term.Type == basic_type.INT {
				result = float64(Arith_Factors_INT(term.Term.Factors, symbol_env))
			} else if term.Term.Type == basic_type.DOUBLE {
				result = Arith_Factors_DOUBLE(term.Term.Factors, symbol_env)
			}
			continue
		}
		switch term.Op {
		case parser.ADD:
			{
				if term.Term.Type == basic_type.INT {
					result = result + float64(Arith_Factors_INT(term.Term.Factors, symbol_env))
				} else if term.Term.Type == basic_type.DOUBLE {
					result = result + Arith_Factors_DOUBLE(term.Term.Factors, symbol_env)
				}
			}

		case parser.SUB:
			{
				if term.Term.Type == basic_type.INT {
					result = result - float64(Arith_Factors_INT(term.Term.Factors, symbol_env))
				} else if term.Term.Type == basic_type.DOUBLE {
					result = result - Arith_Factors_DOUBLE(term.Term.Factors, symbol_env)
				}
			}
		}

	}
	return result
}

func Arith_Terms_STRING(terms []parser.ArithElement, symbol_env *parser.Symbol_Env) string {

	if len(terms) == 1 {
		return Arith_Factors_STRING(terms[0].Term.Factors, symbol_env)
	}

	var result string

	for i, term := range terms {
		//		fmt.Printf("%+v\n", term)
		if i == 0 {
			result = Arith_Factors_STRING(term.Term.Factors, symbol_env)
			continue
		}
		switch term.Op {
		case parser.ADD:
			{
				result = result + Arith_Factors_STRING(term.Term.Factors, symbol_env)
			}

		case parser.SUB:
			{

			}
		}

	}
	return result
}

func Arith_Terms_BOOL(terms []parser.ArithElement, symbol_env *parser.Symbol_Env) bool {

	if len(terms) == 1 {
		return Arith_Factors_BOOL(terms[0].Term.Factors, symbol_env)
	}

	os.Exit(1)

	return true

}
