package ir_gen

import (
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"
	"strconv"
)

type IR_Code struct {
	Left_Operand   Operand
	Right_Operand1 Operand
	Right_Operand2 Operand
	Op             Operator
}

func (code *IR_Code) String() string {
	left := code.Left_Operand.Id

	byte_code := left + " = " + code.Right_Operand1.Str() + " " + code.Op.String() + " "

	if code.Op != NONE {
		byte_code += code.Right_Operand2.Str()
	}

	return byte_code

}

var Temp_Id int = 0

func get_new_temp_name() string {

	temp_id := "_t" + strconv.Itoa(Temp_Id)
	Temp_Id++
	return temp_id

}

func Gen_IR_Binary(operand1 Operand, operand2 Operand, Op Operator) (IR_Code, Operand) {

	temp_id := get_new_temp_name()

	temp := Operand{IfTmp: true, Type: Ident, Id: temp_id}

	code := IR_Code{Left_Operand: temp, Op: Op, Right_Operand1: operand1, Right_Operand2: operand2}

	return code, temp
}

func Gen_IR_Base(factor parser.Factor) (IR_Code, Operand) {

	temp_id := get_new_temp_name()

	temp := Operand{IfTmp: true, Type: Ident, Id: temp_id}

	operand1 := Operand{IfTmp: false}

	switch factor.Type {
	case lexer.INT:
		{
			operand1.Type = Int
			operand1.Int = factor.Int
		}
	case lexer.DOUBLE:
		{
			operand1.Type = Float
			operand1.Float = factor.Float
		}
	case lexer.STRING:
		{
			operand1.Type = String
			operand1.String = factor.String
		}

	case lexer.IDENT:
		{
			operand1.Type = Ident
			operand1.Id = factor.Id
		}

	}

	return IR_Code{Left_Operand: temp, Op: NONE, Right_Operand1: operand1}, temp
}
