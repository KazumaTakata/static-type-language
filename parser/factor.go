package parser

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"strconv"
)

type FactorType int

const (
	FuncCall FactorType = iota + 1
	ArrayMapAccess
	Primitive
	Resolve
)

func (e FactorType) String() string {

	switch e {
	case FuncCall:
		return "FuncCall"
	case Resolve:
		return "Resolve"
	case ArrayMapAccess:
		return "ArrayMapAccess"
	case Primitive:
		return "Primitive"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type Factor struct {
	Int         int
	Float       float64
	String      string
	Id          string
	Bool        bool
	Type        lexer.TokenType
	FactorType  FactorType
	Args        []lexer.Token
	AccessIndex *Arith_expr
	Factor      *Factor
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

				return Factor{Id: ident_token.Value, Type: lexer.IDENT, FactorType: FuncCall, Args: args}

			} else if !tokens.empty() && tokens.peek().Type == lexer.LSQUARE {
				tokens.eat(lexer.LSQUARE)

				index := Parse_Arith_expr(tokens)
				tokens.eat(lexer.RSQUARE)

				return Factor{AccessIndex: &index, Id: ident_token.Value, Type: lexer.IDENT, FactorType: ArrayMapAccess}

			} else if !tokens.empty() && tokens.peek().Type == lexer.DOT {
				tokens.eat(lexer.DOT)

				childfactor := parse_Factor(tokens)
				return Factor{Type: lexer.IDENT, FactorType: Resolve, Id: ident_token.Value, Factor: &childfactor}

			} else {

				return Factor{Id: ident_token.Value, Type: lexer.IDENT, FactorType: Primitive}
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

			return Factor{Bool: bool_value, Type: lexer.BOOL}
		}
	}

	return Factor{}
}
