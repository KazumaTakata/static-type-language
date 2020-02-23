package eval

import (
	"github.com/KazumaTakata/static-typed-language/parser"
)

func Eval_Logic_Bool(logic_expr parser.Logic_expr, symbol_env *parser.Symbol_Env) bool {

	if len(logic_expr.Cmps) == 1 {
		return Eval_Cmp_Bool(logic_expr.Cmps[0], symbol_env)
	}

	var result bool

	for i, cmp_expr := range logic_expr.Cmps {

		if i == 0 {
			result = Eval_Cmp_Bool(cmp_expr, symbol_env)
			continue
		}

		tmp_bool := Eval_Cmp_Bool(cmp_expr, symbol_env)

		switch cmp_expr.LogicOp {
		case parser.AND:
			{
				result = result && tmp_bool
			}
		case parser.OR:
			{
				result = result || tmp_bool
			}
		}

	}

	return result
}
