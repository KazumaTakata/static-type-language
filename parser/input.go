package parser

import (
	"fmt"
	"os"

	"github.com/KazumaTakata/static-typed-language/lexer"
)

type Parser_Input struct {
	Tokens []lexer.Token
	Pos    int
}

func (p *Parser_Input) eat(tokentype lexer.TokenType) {
	if p.peek().Type == tokentype {
		p.Pos = p.Pos + 1
	} else {
		fmt.Printf("eat is not match:got %+v, expected %+v", tokentype, p.peek().Type)
		os.Exit(1)

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
	p.eat(token.Type)
	return token

}

func (p *Parser_Input) assert_next(tokentype lexer.TokenType) lexer.Token {
	token := p.peek()
	if token.Type != tokentype {
		fmt.Printf("assert_next is not match:got %+v, expected %+v\n", p.peek().Type, tokentype)
		os.Exit(1)
	}
	p.eat(token.Type)
	return token

}

func (p *Parser_Input) peek() lexer.Token {
	return p.Tokens[p.Pos]
}

func (p *Parser_Input) peek2() lexer.Token {
	return p.Tokens[p.Pos+1]
}
