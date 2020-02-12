package parser

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/type"
	"strconv"
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
	Type   lexer.TokenType
}

type Parser_Input struct {
	Tokens []lexer.Token
	Pos    int
}

func (p *Parser_Input) eat(token lexer.Token) {
	if p.peek() == token {
		p.Pos = p.Pos + 1
	} else {
		fmt.Errorf("eat is not match:got %+v, expected %+v", token, p.peek())
	}

}

func (p *Parser_Input) empty() bool {
	if len(p.Tokens)-1 < p.Pos {
		return true
	}
	return false

}

func (p *Parser_Input) next() lexer.Token {
	token := p.peek()
	p.eat(token)
	return token

}

func (p *Parser_Input) peek() lexer.Token {
	return p.Tokens[p.Pos]
}
func Parse_Arith_expr(tokens *Parser_Input) Arith_expr {

	terms := []ArithElement{}
	term := parse_Term(tokens)
	terms = append(terms, ArithElement{Term: term, Op: TNONE})

	for !tokens.empty() && (tokens.peek().Type == lexer.ADD || tokens.peek().Type == lexer.SUB) {
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

	}

	return Factor{}
}
