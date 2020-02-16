package parser

import (
	"fmt"

	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/type"
)

type Stmt_Type int

const (
	DECL_STMT Stmt_Type = iota + 1
	FOR_STMT
	IF_STMT
	DEF_STMT
	EXPR
)

func (e Stmt_Type) String() string {

	switch e {
	case DECL_STMT:
		return "DECL_STMT"
	case EXPR:
		return "EXPR"
	case FOR_STMT:
		return "FOR_STMT"
	case IF_STMT:
		return "IF_STMT"
	case DEF_STMT:
		return "DEF_STMT"

	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type For_stmt struct {
	Symbol_Env *Symbol_Env
	Cmp_expr   Cmp_expr
	Stmts      []Stmt
}

type If_stmt struct {
	Symbol_Env *Symbol_Env
	Cmp_expr   Cmp_expr
	Stmts      []Stmt
}

type Stmt struct {
	Type Stmt_Type
	Decl *Decl_stmt
	Expr *Arith_expr
	For  *For_stmt
	If   *If_stmt
	Def  *Def_stmt
}
type Func_param struct {
	Ident string
	Type  basic_type.Type
}
type Def_stmt struct {
	Id          string
	Args        []Func_param
	Stmts       []Stmt
	Return_type basic_type.Type
}

type Decl_stmt struct {
	Id   string
	Type basic_type.Type
	Expr *Arith_expr
}

var getBasicType = map[string]basic_type.Type{"int": basic_type.INT, "double": basic_type.DOUBLE, "string": basic_type.STRING}

func eat_newline(tokens *Parser_Input) {
	for !tokens.empty() && tokens.peek().Type == lexer.NEWLINE {
		tokens.eat(lexer.NEWLINE)
	}
}

func Parse_Stmts(tokens *Parser_Input) []Stmt {

	stmts := []Stmt{}
	_ = []lexer.TokenType{lexer.IDENT, lexer.INT, lexer.DOUBLE, lexer.STRING, lexer.VAR, lexer.FOR}

	for !tokens.empty() && tokens.peek().Type != lexer.RCURL {
		eat_newline(tokens)

		stmt := Parse_Stmt(tokens)
		stmts = append(stmts, stmt)

		eat_newline(tokens)

	}

	return stmts
}

func Parse_Stmt(tokens *Parser_Input) Stmt {

	stmt := Stmt{}

	if !tokens.empty() {

		switch tokens.peek().Type {
		case lexer.VAR:
			{
				tokens.eat(lexer.VAR)
				ident := tokens.assert_next(lexer.IDENT)
				ident_type := tokens.assert_next(lexer.DECL_TYPE)
				tokens.eat(lexer.ASSIGN)
				expr := Parse_Arith_expr(tokens)
				decl_stmt := Decl_stmt{Id: ident.Value, Type: getBasicType[ident_type.Value], Expr: &expr}
				stmt.Decl = &decl_stmt
				stmt.Type = DECL_STMT

			}
		case lexer.INT, lexer.DOUBLE, lexer.STRING, lexer.IDENT:
			{
				expr := Parse_Arith_expr(tokens)
				stmt.Expr = &expr
				stmt.Type = EXPR
			}
		case lexer.FOR:
			{
				tokens.eat(lexer.FOR)
				expr := Parse_Cmp_expr(tokens)
				tokens.eat(lexer.LCURL)
				stmts := Parse_Stmts(tokens)
				tokens.eat(lexer.RCURL)
				for_expr := For_stmt{Cmp_expr: expr, Stmts: stmts}
				stmt.For = &for_expr
				stmt.Type = FOR_STMT

			}
		case lexer.IF:
			{
				tokens.eat(lexer.IF)
				expr := Parse_Cmp_expr(tokens)
				tokens.eat(lexer.LCURL)
				stmts := Parse_Stmts(tokens)
				tokens.eat(lexer.RCURL)
				if_expr := If_stmt{Cmp_expr: expr, Stmts: stmts}
				stmt.If = &if_expr
				stmt.Type = IF_STMT

			}
		case lexer.DEF:
			{
				tokens.eat(lexer.DEF)
				func_name := tokens.assert_next(lexer.IDENT)
				tokens.eat(lexer.LPAREN)
				args := []Func_param{}
				for tokens.peek().Type != lexer.RPAREN {
					id := tokens.assert_next(lexer.IDENT)
					id_type := tokens.assert_next(lexer.DECL_TYPE)
					param := Func_param{Ident: id.Value, Type: getBasicType[id_type.Value]}
					args = append(args, param)

					if tokens.peek().Type != lexer.RPAREN {
						tokens.eat(lexer.COMMA)
					}
				}
				tokens.eat(lexer.RPAREN)
				return_type := tokens.assert_next(lexer.DECL_TYPE)

				tokens.eat(lexer.LCURL)
				stmts := Parse_Stmts(tokens)
				tokens.eat(lexer.RCURL)

				def_expr := Def_stmt{Id: func_name.Value, Args: args, Stmts: stmts, Return_type: getBasicType[return_type.Value]}
				stmt.Def = &def_expr
				stmt.Type = DEF_STMT

			}

		}
	}
	return stmt
}
