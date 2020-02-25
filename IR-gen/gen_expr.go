package ir_gen

import (
	"github.com/KazumaTakata/static-typed-language/parser"
)

func Gen_IR_Expr(cmp parser.Cmp_expr) {
}

func Gen_IR_Arith(arith parser.Arith_expr) ([]IR_Code, Operand) {

	if len(arith.Terms) == 1 {
		codes, operand := Gen_IR_Term(arith.Terms[0].Term)
		return codes, operand
	}

	codes := []IR_Code{}
	code := IR_Code{}

	var operand1 Operand

	for i, term := range arith.Terms {
		if i == 0 {
			code, operand := Gen_IR_Term(term.Term)
			operand1 = operand
			codes = append(codes, code...)
			continue
		}
		term_code, operand2 := Gen_IR_Term(term.Term)

		codes = append(codes, term_code...)

		code, operand1 = Gen_IR_Binary(operand1, operand2, termOpToOp[term.Op])
		codes = append(codes, code)
	}

	return codes, operand1

}
func Gen_IR_Term(term parser.Term) ([]IR_Code, Operand) {

	if len(term.Factors) == 1 {
		code, operand := Gen_IR_Base(term.Factors[0].Factor)
		return []IR_Code{code}, operand

	}

	codes := []IR_Code{}
	code := IR_Code{}

	var operand1 Operand
	var operand2 Operand

	for i, factor := range term.Factors {
		if i == 0 {
			code, operand := Gen_IR_Base(factor.Factor)
			operand1 = operand
			codes = append(codes, code)
			continue
		}
		code, operand2 = Gen_IR_Base(factor.Factor)
		codes = append(codes, code)

		code, operand1 = Gen_IR_Binary(operand1, operand2, factorOpToOp[factor.Op])
		codes = append(codes, code)
	}

	return codes, operand1

}
