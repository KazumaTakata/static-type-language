package lexer

import (
	"fmt"
	"github.com/KazumaTakata/regex_virtualmachine"
)

type TokenType int

const (
	NUMBER TokenType = 0
	ADD    TokenType = 1
	SUB    TokenType = 2
	MUL    TokenType = 3
	DIV    TokenType = 4
)

func (e TokenType) String() string {

	switch e {
	case NUMBER:
		return "NUMBER"
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case MUL:
		return "MUL"
	case DIV:
		return "DIV"

	default:
		return fmt.Sprintf("%d", int(e))
	}
}

var tokenmap map[string]TokenType = map[string]TokenType{
	"NUMBER": NUMBER,
	"ADD":    ADD,
	"SUB":    SUB,
	"MUL":    MUL,
	"DIV":    DIV,
}

type Token struct {
	Type  TokenType
	Value string
}

func GetMatched(named_group map[string]*regex.Group_cap) (string, *regex.Group_cap) {

	for name, group_range := range named_group {
		if group_range != nil {
			return name, group_range
		}
	}

	return "", nil

}

func GetNextToken(regex regex.Regex, input string) (Token, string) {

	_, _, named := regex.Match(input)

	name, group_range := GetMatched(named)

	matched_string := input[group_range.Begin:group_range.End]

	token := Token{Type: tokenmap[name], Value: matched_string}
	return token, input[group_range.End:]

}

func GetTokens(regex regex.Regex, input string) []Token {
	tokens := []Token{}
	for len(input) > 0 {
		token, next_input := GetNextToken(regex, input)
		input = next_input
		tokens = append(tokens, token)
	}
	return tokens
}

func NewLexer() {
}
