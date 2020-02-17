package eval

import (
	"github.com/KazumaTakata/static-typed-language/parser"
	basic_type "github.com/KazumaTakata/static-typed-language/type"
	"os"
)

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
