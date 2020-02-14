package basic_type

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/lexer"
)

type Type int

const (
	INT Type = iota + 1
	DOUBLE
	STRING
	BOOL
)

var LexerTypeToType = map[lexer.TokenType]Type{lexer.INT: INT, lexer.DOUBLE: DOUBLE, lexer.STRING: STRING}

func (e Type) String() string {

	switch e {
	case INT:
		return "INT"
	case DOUBLE:
		return "DOUBLE"
	case STRING:
		return "STRING"
	case BOOL:
		return "BOOL"

	default:
		return fmt.Sprintf("NULL")
	}
}
