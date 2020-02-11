package lexer

import (
	"fmt"
	"github.com/KazumaTakata/regex_virtualmachine"
)

type TokenType int

const (
	INT TokenType = iota + 1
	DOUBLE
	STRING
	ADD
	SUB
	MUL
	DIV
)

func (e TokenType) String() string {

	switch e {
	case INT:
		return "INT"
	case DOUBLE:
		return "DOUBLE"
	case STRING:
		return "STRING"

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
	"INT":    INT,
	"STRING": STRING,
	"DOUBLE": DOUBLE,
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

	if name == "STRING" {
		matched_string = matched_string[1 : len(matched_string)-1]
	}

	token := Token{Type: tokenmap[name], Value: matched_string}
	return token, input[group_range.End:]

}

func eatSpace(regex regex.Regex, input string) string {

	match_range, ifmatch, _ := regex.Match(input)

	if !ifmatch {
		input = input
	} else {
		input = input[match_range[1]:]
	}

	return input

}

func GetTokens(Regex regex.Regex, input string) []Token {

	white_space_regex := regex.NewRegexWithParser("(\\s+)")
	input = eatSpace(white_space_regex, input)

	tokens := []Token{}
	for len(input) > 0 {
		token, next_input := GetNextToken(Regex, input)
		input = eatSpace(white_space_regex, next_input)
		tokens = append(tokens, token)
	}
	return tokens
}

func NewLexer() {
}
