package eval

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/object"
	"github.com/KazumaTakata/static-typed-language/parser"

	basic_type "github.com/KazumaTakata/static-typed-language/type"
	"os"
)

func resolve_variable_int(id string, symbol_env *object.Symbol_Env) int {
	if factor, ok := symbol_env.Table[id]; ok {
		return factor.Int
	} else {
		if symbol_env.Parent_Env != nil {
			return resolve_variable_int(id, symbol_env.Parent_Env)
		}
		fmt.Printf("\nnot defined variable %v\n", id)
		os.Exit(1)
	}
	return 0

}

func resolve_variable_string(id string, symbol_env *object.Symbol_Env) string {
	if factor, ok := symbol_env.Table[id]; ok {
		return factor.String
	} else {
		if symbol_env.Parent_Env != nil {
			return resolve_variable_string(id, symbol_env.Parent_Env)
		}

		fmt.Printf("\nnot defined variable %v\n", id)
		os.Exit(1)

	}
	return ""

}

func resolve_variable_bool(id string, symbol_env *object.Symbol_Env) bool {
	if factor, ok := symbol_env.Table[id]; ok {
		return factor.Bool
	} else {
		if symbol_env.Parent_Env != nil {
			return resolve_variable_bool(id, symbol_env.Parent_Env)
		}

		fmt.Printf("\nnot defined variable %v\n", id)
		os.Exit(1)

	}
	return true

}

func Arith_Factors_INT(factors []parser.TermElement, symbol_env *object.Symbol_Env) int {

	if len(factors) == 1 {

		if factors[0].Factor.Type == lexer.IDENT {
			return resolve_variable_int(factors[0].Factor.Id, symbol_env)
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
				if factor.Factor.Type == lexer.IDENT {
					result = result * resolve_variable_int(factors[0].Factor.Id, symbol_env)
				} else {
					result = result * factor.Factor.Int
				}
			}

		case parser.DIV:
			{
				if factor.Factor.Type == lexer.IDENT {
					result = result / resolve_variable_int(factors[0].Factor.Id, symbol_env)
				} else {
					result = result / factor.Factor.Int
				}

			}
		}

	}
	return result
}

func Arith_Factors_STRING(factors []parser.TermElement, symbol_env *object.Symbol_Env) string {

	if len(factors) == 1 {

		if factors[0].Factor.Type == lexer.IDENT {
			return resolve_variable_string(factors[0].Factor.Id, symbol_env)
		}

		return factors[0].Factor.String
	}
	os.Exit(1)
	return ""
}

func Arith_Factors_BOOL(factors []parser.TermElement, symbol_env *object.Symbol_Env) bool {

	if len(factors) == 1 {

		if factors[0].Factor.Type == lexer.IDENT {
			return resolve_variable_bool(factors[0].Factor.Id, symbol_env)
		}

		return factors[0].Factor.Bool
	}
	os.Exit(1)
	return true
}

func resolve_variable_double(id string, symbol_env *object.Symbol_Env) object.Variable {
	if factor, ok := symbol_env.Table[id]; ok {
		return factor
	} else {
		if symbol_env.Parent_Env != nil {
			return resolve_variable_double(id, symbol_env.Parent_Env)

		}

		fmt.Printf("\nnot defined variable %v\n", id)
		os.Exit(1)

	}
	return object.Variable{}

}

func Arith_Factors_DOUBLE(factors []parser.TermElement, symbol_env *object.Symbol_Env) float64 {

	if len(factors) == 1 {

		if factors[0].Factor.Type == lexer.IDENT {
			double := resolve_variable_double(factors[0].Factor.Id, symbol_env)
			return double.Double
		}

		return factors[0].Factor.Float
	}

	var result float64

	for i, factor := range factors {
		if i == 0 {

			switch factor.Factor.Type {

			case lexer.IDENT:
				{
					variable := resolve_variable_double(factor.Factor.Id, symbol_env)

					if variable.Type == basic_type.DOUBLE {
						result = variable.Double
					} else if variable.Type == basic_type.INT {
						result = float64(variable.Int)
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
						variable := resolve_variable_double(factor.Factor.Id, symbol_env)

						if variable.Type == basic_type.DOUBLE {
							result = result * variable.Double
						} else if variable.Type == basic_type.INT {
							result = result * float64(variable.Int)
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
						variable := resolve_variable_double(factor.Factor.Id, symbol_env)

						if variable.Type == basic_type.DOUBLE {
							result = result / variable.Double
						} else if variable.Type == basic_type.INT {
							result = result / float64(variable.Int)
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

func Arith_Terms_INT(terms []parser.ArithElement, symbol_env *object.Symbol_Env) int {

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

func Arith_Terms_DOUBLE(terms []parser.ArithElement, symbol_env *object.Symbol_Env) float64 {

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

func Arith_Terms_STRING(terms []parser.ArithElement, symbol_env *object.Symbol_Env) string {

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

func Arith_Terms_BOOL(terms []parser.ArithElement, symbol_env *object.Symbol_Env) bool {

	if len(terms) == 1 {
		return Arith_Factors_BOOL(terms[0].Term.Factors, symbol_env)
	}

	os.Exit(1)

	return true

}

func Eval_Cmp(cmp_expr parser.Cmp_expr, symbol_env *object.Symbol_Env) bool {
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

func Eval_Stmts(stmts []parser.Stmt, symbol_env *object.Symbol_Env) {

	for _, stmt := range stmts {
		Eval_Stmt(stmt, symbol_env)
	}
}

func Eval_Stmt(stmt parser.Stmt, symbol_env *object.Symbol_Env) {

	switch stmt.Type {

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

			switch stmt.Decl.Expr.Type {
			case basic_type.INT:
				{
					result := Arith_Terms_INT(stmt.Decl.Expr.Terms, symbol_env)
					symbol_env.Table[stmt.Decl.Id] = object.Variable{Int: result, Type: basic_type.INT}

					fmt.Printf("%+v", result)

				}
			case basic_type.DOUBLE:
				{
					result := Arith_Terms_DOUBLE(stmt.Decl.Expr.Terms, symbol_env)
					symbol_env.Table[stmt.Decl.Id] = object.Variable{Double: result, Type: basic_type.DOUBLE}

					fmt.Printf("%+v", result)

				}
			case basic_type.STRING:
				{
					result := Arith_Terms_STRING(stmt.Decl.Expr.Terms, symbol_env)
					symbol_env.Table[stmt.Decl.Id] = object.Variable{String: result, Type: basic_type.STRING}

					fmt.Printf("%+v", result)

				}

			case basic_type.BOOL:
				{
					result := Arith_Terms_BOOL(stmt.Decl.Expr.Terms, symbol_env)
					symbol_env.Table[stmt.Decl.Id] = object.Variable{Bool: result, Type: basic_type.STRING}

					fmt.Printf("%+v", result)

				}

			}

		}

	case parser.EXPR:
		{

			switch stmt.Expr.Type {
			case basic_type.INT:
				{
					result := Arith_Terms_INT(stmt.Expr.Terms, symbol_env)
					fmt.Printf("%+v", result)

				}
			case basic_type.DOUBLE:
				{
					result := Arith_Terms_DOUBLE(stmt.Expr.Terms, symbol_env)

					fmt.Printf("%+v", result)

				}
			case basic_type.STRING:
				{
					result := Arith_Terms_STRING(stmt.Expr.Terms, symbol_env)

					fmt.Printf("%+v", result)

				}
			case basic_type.BOOL:
				{
					result := Arith_Terms_BOOL(stmt.Expr.Terms, symbol_env)

					fmt.Printf("%+v", result)

				}

			}
		}

	}
}
