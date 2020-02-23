package parser

import (
	"fmt"

	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/type"
)

type FactorOp int

const (
	MUL   FactorOp = 0
	DIV   FactorOp = 1
	FNONE FactorOp = 2
)

func (e FactorOp) String() string {

	switch e {
	case MUL:
		return "MUL"
	case DIV:
		return "DIV"
	case FNONE:
		return "FNONE"

	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type TermOp int

const (
	ADD   TermOp = 0
	SUB   TermOp = 1
	TNONE TermOp = 2
)

func (e TermOp) String() string {

	switch e {
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case TNONE:
		return "TNONE"

	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type Arith_expr struct {
	Terms []ArithElement
	Type  basic_type.Variable_Type
}

type ArithElement struct {
	Term Term
	Op   TermOp
}

type Term struct {
	Factors []TermElement
	Type    basic_type.Variable_Type
}

type TermElement struct {
	Factor Factor
	Op     FactorOp
}

var tokenToArithOp = map[lexer.TokenType]ArithOp{lexer.EQUAL: EQUAL, lexer.NOTEQUAL: NOTEQUAL, lexer.GT: GT, lexer.LT: LT}

func Parse_Arith_expr(tokens *Parser_Input) Arith_expr {

	terms := []ArithElement{}
	term := parse_Term(tokens)
	terms = append(terms, ArithElement{Term: term, Op: TNONE})

	for !tokens.empty() && tokens.peek().Type != lexer.LCURL && tokens.peek().Type != lexer.NEWLINE && (tokens.peek().Type == lexer.ADD || tokens.peek().Type == lexer.SUB) {
		op := tokens.next()
		var top TermOp
		switch op.Type {
		case lexer.ADD:
			{
				top = ADD
			}
		case lexer.SUB:
			{
				top = SUB
			}
		default:
			{
				top = TNONE
			}

		}

		term := parse_Term(tokens)
		terms = append(terms, ArithElement{Term: term, Op: top})

	}

	if !tokens.empty() && tokens.peek().Type == lexer.NEWLINE {
		tokens.eat(lexer.NEWLINE)
	}

	return Arith_expr{Terms: terms}
}

func parse_Term(tokens *Parser_Input) Term {
	factors := []TermElement{}
	factor := parse_Factor(tokens)
	factors = append(factors, TermElement{Factor: factor, Op: FNONE})

	for !tokens.empty() && (tokens.peek().Type == lexer.MUL || tokens.peek().Type == lexer.DIV) {
		op := tokens.next()
		var fop FactorOp
		switch op.Type {
		case lexer.MUL:
			{
				fop = MUL
			}
		case lexer.DIV:
			{
				fop = DIV
			}
		default:
			{
				fop = FNONE
			}

		}

		factor := parse_Factor(tokens)
		factors = append(factors, TermElement{Factor: factor, Op: fop})

	}

	return Term{Factors: factors}
}
