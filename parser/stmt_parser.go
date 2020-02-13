package parser

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/type"
)

type Stmt_Type int

const (
	DECL_STMT Stmt_Type = iota + 1
	EXPR
)

func (e Stmt_Type) String() string {

	switch e {
	case DECL_STMT:
		return "DECL_STMT"
	case EXPR:
		return "EXPR"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type Stmt struct {
	Type Stmt_Type
	Decl *Decl_stmt
	Expr *Arith_expr
}

type Decl_stmt struct {
	Id   string
	Type basic_type.Type
	Expr Arith_expr
}

var getBasicType = map[string]basic_type.Type{"int": basic_type.INT, "double": basic_type.DOUBLE, "string": basic_type.STRING}

func Parse_Stmt(tokens *Parser_Input) Stmt {

	stmt := Stmt{}

	if !tokens.empty() {

		switch tokens.peek().Type {
		case lexer.VAR:
			{
				tokens.eat(lexer.VAR)
				ident := tokens.assert_next(lexer.IDENT)
				ident_type := tokens.assert_next(lexer.DECL_TYPE)
				tokens.eat(lexer.EQUAL)
				expr := Parse_Arith_expr(tokens)
				decl_stmt := Decl_stmt{Id: ident.Value, Type: getBasicType[ident_type.Value], Expr: expr}
				stmt.Decl = &decl_stmt
				stmt.Type = DECL_STMT

			}
		case lexer.INT, lexer.DOUBLE, lexer.STRING, lexer.IDENT:
			{
				expr := Parse_Arith_expr(tokens)
				stmt.Expr = &expr
				stmt.Type = EXPR
			}

		}
	}
	return stmt
}
