package main

import (
	"fmt"
	"github.com/KazumaTakata/Static-Typed-Language/lexer"
	"github.com/KazumaTakata/regex_virtualmachine"
	"strconv"
	"strings"
)

type FactorOp int

const (
	MUL   FactorOp = 0
	DIV   FactorOp = 1
	FNONE FactorOp = 2
)

type TermOp int

const (
	ADD   TermOp = 0
	SUB   TermOp = 1
	TNONE TermOp = 2
)

type Arith_expr struct {
	terms []ArithElement
}

type ArithElement struct {
	term Term
	op   TermOp
}

type Term struct {
	factors []TermElement
}

type TermElement struct {
	factor Factor
	op     FactorOp
}

type Factor struct {
	Number int
}

type Parser_Input struct {
	tokens []lexer.Token
	pos    int
}

func (p *Parser_Input) eat(token lexer.Token) {
	if p.peek() == token {
		p.pos = p.pos + 1
	} else {
		fmt.Errorf("eat is not match:got %+v, expected %+v", token, p.peek())
	}

}

func (p *Parser_Input) next() lexer.Token {
	token := p.peek()
	p.eat(token)
	return token

}

func (p *Parser_Input) peek() lexer.Token {
	return p.tokens[p.pos]
}
func parse_Arith_expr(tokens *Parser_Input) Arith_expr {
	terms := []ArithElement{}
	term := parse_Term(tokens)
	terms = append(terms, ArithElement{term: term, op: TNONE})

	for tokens.peek().Type == lexer.ADD || tokens.peek().Type == lexer.SUB {
		op := tokens.next()
		var top TermOp
		switch op.Type {
		case lexer.ADD:
			{
				top = ADD
			}
		case lexer.DIV:
			{
				top = SUB
			}
		default:
			{
				top = TNONE
			}

		}

		term := parse_Term(tokens)
		terms = append(terms, ArithElement{term: term, op: top})

	}

	return Arith_expr{terms: terms}
}

func parse_Term(tokens *Parser_Input) Term {
	factors := []TermElement{}
	factor := parse_Factor(tokens)
	factors = append(factors, TermElement{factor: factor, op: FNONE})

	for tokens.peek().Type == lexer.MUL || tokens.peek().Type == lexer.DIV {
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
		factors = append(factors, TermElement{factor: factor, op: fop})

	}

	return Term{factors: factors}
}

func parse_Factor(tokens *Parser_Input) Factor {

	if tokens.peek().Type == lexer.NUMBER {
		number_token := tokens.next()
		number, _ := strconv.Atoi(number_token.Value)
		return Factor{Number: number}
	}

	return Factor{}
}

func main() {

	lexer_rules := [][]string{}
	lexer_rules = append(lexer_rules, []string{"NUMBER", "\\d+"})
	lexer_rules = append(lexer_rules, []string{"ADD", "\\+"})
	lexer_rules = append(lexer_rules, []string{"SUB", "\\-"})
	lexer_rules = append(lexer_rules, []string{"MUL", "\\*"})
	lexer_rules = append(lexer_rules, []string{"DIV", "\\/"})

	regex_parts := []string{}

	for _, rule := range lexer_rules {
		regex_parts = append(regex_parts, fmt.Sprintf("(?<%s>%s)", rule[0], rule[1]))
	}

	regex_string := strings.Join(regex_parts, "|")
	//fmt.Printf("%s", regex_string)

	regex := regex.NewRegexWithParser(regex_string)

	tokens := lexer.GetTokens(regex, "13*33-35")

	fmt.Printf("%+v", tokens)

}
