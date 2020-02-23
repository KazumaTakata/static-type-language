package main

import (
	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"
	"github.com/KazumaTakata/static-typed-language/tree-eval"
	"github.com/KazumaTakata/static-typed-language/type-system"
	"golang.org/x/crypto/openpgp"
	"io/ioutil"
	"os"
	"strconv"
)

type IR_Code struct {
	Left_Operand   Operand
	Right_Operand1 Operand
	Right_Operand2 Operand
	Op             Operator
}

var Temp_Id int = 0

func get_new_temp_name() string {

	temp_id := "_t" + strconv.Itoa(Temp_Id)
	Temp_Id++
	return temp_id

}

func Gen_IR_Binary(factor1 parser.Factor, factor2 parser.Factor, Op Operator) []IR_Code {

	temp_id := get_new_temp_name()

	temp := Operand{IfTmp: true, Type: Ident, Id: temp_id}

	code1, temp_op1 := Gen_IR_Base(factor1)
	code2, temp_op2 := Gen_IR_Base(factor2)

	code := IR_Code{Left_Operand: temp, Op: Op, Right_Operand1: temp_op1, Right_Operand2: temp_op2}

	return []IR_Code{code1, code2, code}
}

func Gen_IR_Base(factor parser.Factor) (IR_Code, Operand) {

	temp_id := get_new_temp_name()

	temp := Operand{IfTmp: true, Type: Ident, Id: temp_id}

	operand1 := Operand{IfTmp: false}

	switch factor.Type {
	case lexer.INT:
		{
			operand1.Type = Int
		}
	case lexer.DOUBLE:
		{
			operand1.Type = Float
		}
	case lexer.STRING:
		{
			operand1.Type = String
		}

	case lexer.IDENT:
		{
			operand1.Type = Ident
		}

	}

	return IR_Code{Left_Operand: temp, Op: NONE, Right_Operand1: operand1}, temp
}
