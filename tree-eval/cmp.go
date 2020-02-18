package eval

import (
	"github.com/KazumaTakata/static-typed-language/parser"
	basic_type "github.com/KazumaTakata/static-typed-language/type"
	"os"
)

func Eval_Cmp_Bool(cmp_expr parser.Cmp_expr, symbol_env *parser.Symbol_Env) bool {
	if cmp_expr.Left == nil {
		return Arith_Terms_BOOL(cmp_expr.Left.Terms, symbol_env)
	}

	switch cmp_expr.Left.Type {
	case basic_type.INT:
		{
			cmp1 := Arith_Terms_INT(cmp_expr.Left.Terms, symbol_env)

			if cmp_expr.Right.Type != basic_type.INT {
				os.Exit(1)
			} else {
				cmp2 := Arith_Terms_INT(cmp_expr.Right.Terms, symbol_env)

				if cmp1 == cmp2 {
					return true
				} else {
					return false
				}

			}

		}
	case basic_type.DOUBLE:
		{
			cmp1 := Arith_Terms_DOUBLE(cmp_expr.Left.Terms, symbol_env)

			if cmp_expr.Right.Type != basic_type.DOUBLE {
				os.Exit(1)
			} else {
				cmp2 := Arith_Terms_DOUBLE(cmp_expr.Right.Terms, symbol_env)

				if cmp1 == cmp2 {
					return true
				} else {
					return false
				}

			}

		}
	case basic_type.STRING:
		{
			cmp1 := Arith_Terms_STRING(cmp_expr.Left.Terms, symbol_env)

			if cmp_expr.Right.Type != basic_type.STRING {
				os.Exit(1)
			} else {
				cmp2 := Arith_Terms_STRING(cmp_expr.Right.Terms, symbol_env)

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
