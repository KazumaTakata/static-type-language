package eval

import (
	"fmt"
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

func Eval_Cmp(cmp_expr parser.Cmp_expr, symbol_env *parser.Symbol_Env) bool {
	if len(cmp_expr.Ariths) == 1 {
		return Arith_Terms_BOOL(cmp_expr.Ariths[0].Arith.Terms, symbol_env)
	}

	switch cmp_expr.Ariths[0].Arith.Type {
	case basic_type.INT:
		{
			cmp1 := Arith_Terms_INT(cmp_expr.Ariths[0].Arith.Terms, symbol_env)

			if cmp_expr.Ariths[1].Arith.Type != basic_type.INT {
				os.Exit(1)
			} else {
				cmp2 := Arith_Terms_INT(cmp_expr.Ariths[1].Arith.Terms, symbol_env)

				if cmp1 == cmp2 {
					return true
				} else {
					return false
				}

			}

		}
	case basic_type.DOUBLE:
		{
			cmp1 := Arith_Terms_DOUBLE(cmp_expr.Ariths[1].Arith.Terms, symbol_env)

			if cmp_expr.Ariths[1].Arith.Type != basic_type.DOUBLE {
				os.Exit(1)
			} else {
				cmp2 := Arith_Terms_DOUBLE(cmp_expr.Ariths[1].Arith.Terms, symbol_env)

				if cmp1 == cmp2 {
					return true
				} else {
					return false
				}

			}

		}
	case basic_type.STRING:
		{
			cmp1 := Arith_Terms_STRING(cmp_expr.Ariths[1].Arith.Terms, symbol_env)

			if cmp_expr.Ariths[1].Arith.Type != basic_type.STRING {
				os.Exit(1)
			} else {
				cmp2 := Arith_Terms_STRING(cmp_expr.Ariths[1].Arith.Terms, symbol_env)

				if cmp1 == cmp2 {
					return true
				} else {
					return false
				}

			}
		}

	}

	return true
}

func Eval_Stmts(stmts []parser.Stmt, symbol_env *parser.Symbol_Env) {

	for _, stmt := range stmts {
		if_return := Eval_Stmt(stmt, symbol_env)

		if if_return {
			break
		}
	}
}

func Calc_Arith(expr *parser.Arith_expr, symbol_env *parser.Symbol_Env) parser.Object {

	switch expr.Type {
	case basic_type.INT:
		{
			result := Arith_Terms_INT(expr.Terms, symbol_env)
			return parser.Object{Type: parser.VariableObj, Variable: &parser.Variable{Int: result, Type: basic_type.INT}}

		}
	case basic_type.DOUBLE:
		{
			result := Arith_Terms_DOUBLE(expr.Terms, symbol_env)
			return parser.Object{Type: parser.VariableObj, Variable: &parser.Variable{Double: result, Type: basic_type.DOUBLE}}

		}
	case basic_type.STRING:
		{
			result := Arith_Terms_STRING(expr.Terms, symbol_env)
			return parser.Object{Type: parser.VariableObj, Variable: &parser.Variable{String: result, Type: basic_type.STRING}}

		}

	case basic_type.BOOL:
		{
			result := Arith_Terms_BOOL(expr.Terms, symbol_env)
			return parser.Object{Type: parser.VariableObj, Variable: &parser.Variable{Bool: result, Type: basic_type.BOOL}}

		}

	}

	return parser.Object{}
}

func Eval_Stmt(stmt parser.Stmt, symbol_env *parser.Symbol_Env) bool {

	switch stmt.Type {

	case parser.DEF_STMT:
		{

		}
	case parser.RETURN_STMT:
		{
			return_value := Calc_Arith(&stmt.Return.Cmp_expr.Ariths[0].Arith, symbol_env)
			symbol_env.Return_Value = &return_value

			return true

		}

	case parser.FOR_STMT:
		{
			for Eval_Cmp(stmt.For.Cmp_expr, symbol_env) {
				Eval_Stmts(stmt.For.Stmts, stmt.For.Symbol_Env)
			}

		}
	case parser.IF_STMT:
		{
			if Eval_Cmp(stmt.If.Cmp_expr, symbol_env) {
				Eval_Stmts(stmt.If.Stmts, stmt.If.Symbol_Env)
			}
		}

	case parser.DECL_STMT:
		{

			result := Calc_Arith(stmt.Decl.Expr, symbol_env)
			symbol_env.Table[stmt.Decl.Id] = result
			//fmt.Printf("%+v", result)

		}

	case parser.EXPR:
		{

			switch stmt.Expr.Type {
			case basic_type.INT:
				{
					result := Arith_Terms_INT(stmt.Expr.Terms, symbol_env)
					fmt.Printf("%+v\n", result)

				}
			case basic_type.DOUBLE:
				{
					result := Arith_Terms_DOUBLE(stmt.Expr.Terms, symbol_env)

					fmt.Printf("%+v\n", result)

				}
			case basic_type.STRING:
				{
					result := Arith_Terms_STRING(stmt.Expr.Terms, symbol_env)

					fmt.Printf("%+v\n", result)

				}
			case basic_type.BOOL:
				{
					result := Arith_Terms_BOOL(stmt.Expr.Terms, symbol_env)

					fmt.Printf("%+v\n", result)

				}

			}
		}

	}

	return false
}
