package ir_gen

import (
	"github.com/KazumaTakata/static-typed-language/parser"
	"strconv"
)

var label_id int = 0

func get_new_label_name() string {
	temp_id := "L" + strconv.Itoa(label_id)
	label_id++
	return temp_id

}

func Gen_IR_Stmts(stmts []parser.Stmt) []IR_Code {

	codes := []IR_Code{}

	for _, stmt := range stmts {
		codes = append(codes, Gen_IR_Stmt(stmt)...)
	}

	return codes
}

func Gen_IR_Stmt(stmt parser.Stmt) []IR_Code {

	switch stmt.Type {

	case parser.DECL_STMT:
		{
			codes := Gen_IR_Assign(*stmt.Decl.Assign)
			codes[len(codes)-1].Left_Operand = Operand{Type: Ident, Id: stmt.Decl.Id}

			return codes

		}
	case parser.ASSIGN_STMT:
		{
			codes := Gen_IR_Assign(*stmt.Assign.Assign)
			codes[len(codes)-1].Left_Operand = Operand{Type: Ident, Id: stmt.Assign.Id}
			return codes
		}

	case parser.IF_STMT:
		{

			new_label := get_new_label_name()
			cmp_codes := Gen_IR_Cmp(stmt.If.Cmp_expr)
			if_code := IR_Code{Type: Ifz, Right_Operand1: cmp_codes[len(cmp_codes)-1].Left_Operand, Right_Operand2: Operand{Type: String, String: new_label}}
			stmts_codes := Gen_IR_Stmts(stmt.If.Stmts)
			label_code := IR_Code{Type: Label, Right_Operand1: Operand{Type: String, String: new_label}}

			codes := append(cmp_codes, if_code)
			codes = append(codes, stmts_codes...)
			codes = append(codes, label_code)

			return codes

		}

	}

	return []IR_Code{}

}

func Gen_IR_Assign(assign parser.Assign) []IR_Code {

	switch assign.Type {
	case parser.EXPR_ASSIGN:
		{
			return Gen_IR_Cmp(*assign.Expr)
		}
	case parser.INIT_ASSIGN:
		{
		}
	}

	return []IR_Code{}
}
