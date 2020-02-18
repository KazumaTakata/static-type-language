package parser

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/type"
	"strconv"
)

type ArithOp int

const (
	EQUAL ArithOp = iota + 1
	NOTEQUAL
	MTHAN
	METHAN
	LTHAN
	LETHAN
	ANONE
)

func (e ArithOp) String() string {

	switch e {
	case EQUAL:
		return "EQUAL"
	case NOTEQUAL:
		return "NOTEQUAL"
	case MTHAN:
		return "MTHAN"
	case METHAN:
		return "METHAN"
	case LTHAN:
		return "LTHAN"
	case LETHAN:
		return "LETHAN"
	case ANONE:
		return "ANONE"

	default:
		return fmt.Sprintf("%d", int(e))
	}
}

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

type Cmp_expr struct {
	Left  *Arith_expr
	Right *Arith_expr
	Op    ArithOp
	Type  basic_type.Type
}

type Arith_expr struct {
	Terms []ArithElement
	Type  basic_type.Type
}

type ArithElement struct {
	Term Term
	Op   TermOp
}

type Term struct {
	Factors []TermElement
	Type    basic_type.Type
}

type TermElement struct {
	Factor Factor
	Op     FactorOp
}

type Factor struct {
	Int    int
	Float  float64
	String string
	Id     string
	Bool   bool
	Type   lexer.TokenType
	IsCall bool
	Args   []lexer.Token
}

func Parse_Cmp_expr(tokens *Parser_Input) Cmp_expr {

	cmp_expr := Cmp_expr{}
	arith := Parse_Arith_expr(tokens)
	cmp_expr.Left = &arith

	if !tokens.empty() && (tokens.peek().Type == lexer.EQUAL || tokens.peek().Type == lexer.NOTEQUAL) {
		op := tokens.next()
		var aop ArithOp
		switch op.Type {
		case lexer.EQUAL:
			{
				aop = EQUAL
			}
		case lexer.NOTEQUAL:
			{
				aop = NOTEQUAL
			}
		default:
			{
				aop = ANONE
			}

		}

		arith := Parse_Arith_expr(tokens)
		cmp_expr.Right = &arith
		cmp_expr.Op = aop

	}

	if !tokens.empty() && tokens.peek().Type == lexer.NEWLINE {
		tokens.eat(lexer.NEWLINE)
	}

	return cmp_expr
}

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

func parse_Factor(tokens *Parser_Input) Factor {

	switch tokens.peek().Type {
	case lexer.INT:
		{
			number_token := tokens.next()
			number, _ := strconv.Atoi(number_token.Value)
			return Factor{Int: number, Type: lexer.INT}
		}
	case lexer.DOUBLE:
		{
			number_token := tokens.next()
			number, _ := strconv.ParseFloat(number_token.Value, 64)
			return Factor{Float: number, Type: lexer.DOUBLE}
		}
	case lexer.STRING:
		{
			string_token := tokens.next()
			return Factor{String: string_token.Value, Type: lexer.STRING}
		}
	case lexer.IDENT:
		{
			ident_token := tokens.next()
			if !tokens.empty() && tokens.peek().Type == lexer.LPAREN {
				tokens.eat(lexer.LPAREN)
				args := []lexer.Token{}
				for tokens.peek().Type != lexer.RPAREN {
					id := tokens.next()
					args = append(args, id)

					if tokens.peek().Type != lexer.RPAREN {
						tokens.eat(lexer.COMMA)
					}
				}
				tokens.eat(lexer.RPAREN)

				return Factor{Id: ident_token.Value, Type: lexer.IDENT, IsCall: true, Args: args}

			} else {

				return Factor{Id: ident_token.Value, Type: lexer.IDENT}
			}
		}
	case lexer.BOOL:
		{
			bool_token := tokens.next()
			var bool_value bool
			if bool_token.Value == "true" {
				bool_value = true
			} else {
				bool_value = false
			}

			return Factor{Bool: bool_value, Type: lexer.IDENT}
		}
	}

	return Factor{}
}
