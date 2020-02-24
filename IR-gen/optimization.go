package ir_gen

func Get_Available(codes []IR_Code) [][]IR_Code {

	available_exprs := [][]IR_Code{}

	available_exprs = append(available_exprs, []IR_Code{})

	for _, code := range codes {
		left := code.Left_Operand
		new_availables := []IR_Code{}
		for _, ir_code := range available_exprs[len(available_exprs)-1] {
			if left != ir_code.Left_Operand {
				new_availables = append(new_availables, ir_code)
			}
		}
		availables := append(new_availables, code)
		available_exprs = append(available_exprs, availables)
	}
	return available_exprs
}

func Common_Subexpression_Elimination(codes []IR_Code, availables [][]IR_Code) {

	for i, code := range codes {
		if i == 0 {
			continue
		}

		for _, available := range availables[i] {
			if available.Right_Operand1 == code.Right_Operand1 {
				if available.Op == NONE {
					codes[i].Right_Operand1 = available.Left_Operand
				} else {
					if available.Op == code.Op && available.Right_Operand2 == code.Right_Operand2 {
						codes[i].Right_Operand1 = available.Left_Operand
						codes[i].Op = NONE
					}
				}
			}

		}

	}
}

func Copy_Propagation(codes []IR_Code, availables [][]IR_Code) {

	for i, code := range codes {
		for _, available := range availables[i+1] {
			if available.Op == NONE {
				if available.Left_Operand == code.Right_Operand1 {
					codes[i].Right_Operand1 = available.Right_Operand1
				}

				if available.Left_Operand == code.Right_Operand2 {
					codes[i].Right_Operand2 = available.Right_Operand1

				}
			}
		}

	}
}

func Dead_Code_Elimination(codes []IR_Code) []IR_Code {

	eliminated := []IR_Code{}

	liveness := map[Operand]bool{}
	liveness[codes[len(codes)-1].Left_Operand] = true

	for i := len(codes) - 1; i >= 0; i-- {

		if _, ok := liveness[codes[i].Left_Operand]; ok {
			delete(liveness, codes[i].Left_Operand)
			eliminated = append([]IR_Code{codes[i]}, eliminated...)
			liveness[codes[i].Right_Operand1] = true
			if codes[i].Op != NONE {
				liveness[codes[i].Right_Operand2] = true
			}
		}
	}
	return eliminated
}

func Constant_Folding(codes []IR_Code) {

	for i, code := range codes {
		if code.Right_Operand1.Type != Ident && code.Right_Operand2.Type != Ident {
			switch code.Op {
			case ADD:
				{
					if code.Right_Operand1.Type == Int {
						if code.Right_Operand2.Type == Int {
							codes[i].Right_Operand1.Int = code.Right_Operand1.Int + code.Right_Operand2.Int
						} else if code.Right_Operand2.Type == Float {
							codes[i].Right_Operand1.Float = float64(code.Right_Operand1.Int) + code.Right_Operand2.Float
							codes[i].Right_Operand1.Type = Float
						}
					}

				}
			case SUB:
				{

				}
			case MUL:
				{
					if code.Right_Operand1.Type == Int {
						if code.Right_Operand2.Type == Int {
							codes[i].Right_Operand1.Int = code.Right_Operand1.Int * code.Right_Operand2.Int
						} else if code.Right_Operand2.Type == Float {
							codes[i].Right_Operand1.Float = float64(code.Right_Operand1.Int) * code.Right_Operand2.Float
							codes[i].Right_Operand1.Type = Float
						}
					}

				}
			case DIV:
				{

				}

			}
			codes[i].Op = NONE
		}
	}

}
