package parser

import (
	"strconv"

	"github.com/KazumaTakata/static-typed-language/lexer"
)

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
