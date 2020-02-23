package parser

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/type"
)

type ArithOp int

const (
	EQUAL ArithOp = iota + 1
	NOTEQUAL
	GT
	LT
	ANONE
)

func (e ArithOp) String() string {

	switch e {
	case EQUAL:
		return "EQUAL"
	case NOTEQUAL:
		return "NOTEQUAL"
	case GT:
		return "GT"
	case LT:
		return "LT"
	case ANONE:
		return "ANONE"

	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type Cmp_expr struct {
	Left    *Arith_expr
	Right   *Arith_expr
	Op      ArithOp
	LogicOp LogicOp
	Type    basic_type.Variable_Type
}

func Parse_Cmp_expr(tokens *Parser_Input) Cmp_expr {

	cmp_expr := Cmp_expr{}
	arith := Parse_Arith_expr(tokens)
	cmp_expr.Left = &arith

	if !tokens.empty() && (tokens.peek().Type == lexer.EQUAL || tokens.peek().Type == lexer.NOTEQUAL || tokens.peek().Type == lexer.GT || tokens.peek().Type == lexer.LT) {
		op := tokens.next()
		aop := tokenToArithOp[op.Type]
		arith := Parse_Arith_expr(tokens)
		cmp_expr.Right = &arith
		cmp_expr.Op = aop
	}

	if !tokens.empty() && tokens.peek().Type == lexer.NEWLINE {
		tokens.eat(lexer.NEWLINE)
	}

	return cmp_expr
}

type Logic_expr struct {
	Cmps []Cmp_expr
	Op   LogicOp
	Type basic_type.Variable_Type
}

type LogicOp int

const (
	AND LogicOp = iota + 1
	OR
	LNONE
)

func (e LogicOp) String() string {

	switch e {
	case AND:
		return "AND"
	case OR:
		return "OR"
	case LNONE:
		return "LNONE"

	default:
		return fmt.Sprintf("%d", int(e))
	}
}

var tokenToLogicOp = map[lexer.TokenType]LogicOp{lexer.AND: AND, lexer.OR: OR}

func Parse_Logic_expr(tokens *Parser_Input) Logic_expr {

	logic_expr := Logic_expr{}
	cmp := Parse_Cmp_expr(tokens)
	logic_expr.Cmps = append(logic_expr.Cmps, cmp)

	if !tokens.empty() && (tokens.peek().Type == lexer.ADD || tokens.peek().Type == lexer.OR) {
		op := tokens.next()
		cmp := Parse_Cmp_expr(tokens)
		cmp.LogicOp = tokenToLogicOp[op.Type]
		logic_expr.Cmps = append(logic_expr.Cmps, cmp)
	}

	if !tokens.empty() && tokens.peek().Type == lexer.NEWLINE {
		tokens.eat(lexer.NEWLINE)
	}

	return logic_expr
}
