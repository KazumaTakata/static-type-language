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
	RETURN_STMT
	ASSIGN_STMT
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
	case RETURN_STMT:
		return "RETURN_STMT"
	case ASSIGN_STMT:
		return "ASSIGN_STMT"

	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type InitType int

const (
	ARRAY_INIT InitType = iota + 1
	MAP_INIT
)

func (e InitType) String() string {

	switch e {
	case ARRAY_INIT:
		return "ARRAY_INIT"
	case MAP_INIT:
		return "MAP_INIT"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type AssignType int

const (
	EXPR_ASSIGN AssignType = iota + 1
	INIT_ASSIGN
)

func (e AssignType) String() string {

	switch e {
	case EXPR_ASSIGN:
		return "EXPR_ASSIGN"
	case INIT_ASSIGN:
		return "INIT_ASSIGN"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type Assign struct {
	Type AssignType
	Expr *Cmp_expr
	Init *Init
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

type Return_stmt struct {
	Type     basic_type.Type
	Cmp_expr Cmp_expr
}

type Stmt struct {
	Type   Stmt_Type
	Decl   *Decl_stmt
	Expr   *Arith_expr
	For    *For_stmt
	If     *If_stmt
	Def    *Def_stmt
	Return *Return_stmt
	Assign *Assign_stmt
}
type Func_param struct {
	Ident string
	Type  basic_type.Type
}

type Def_stmt struct {
	Symbol_Env   *Symbol_Env
	Id           string
	Args         []Func_param
	Stmts        []Stmt
	Return_type  basic_type.Type
	Return_value *Object
}

type Assign_stmt struct {
	Id     string
	Assign *Assign
}

type Decl_stmt struct {
	Id     string
	Type   basic_type.Variable_Type
	Assign *Assign
}
type Array struct {
	Type      basic_type.Type
	InitValue []*Cmp_expr
}

type Map struct{}

type Init struct {
	Type  InitType
	Array *Array
	Map   *Map
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

func parse_Type(tokens *Parser_Input) basic_type.Variable_Type {

	if tokens.peek().Type == lexer.LSQUARE {
		tokens.eat(lexer.LSQUARE)
		tokens.eat(lexer.RSQUARE)
		ident_type := tokens.assert_next(lexer.DECL_TYPE)

		return basic_type.Variable_Type{Type: getBasicType[ident_type.Value], DataStructType: basic_type.ARRAY}

	} else {
		ident_type := tokens.assert_next(lexer.DECL_TYPE)
		return basic_type.Variable_Type{Type: getBasicType[ident_type.Value], DataStructType: basic_type.PRIMITIVE}
	}

}

func parse_Init(tokens *Parser_Input) Init {
	switch tokens.peek().Type {
	case lexer.LSQUARE:
		{
			tokens.eat(lexer.LSQUARE)
			tokens.eat(lexer.RSQUARE)
			ident_type := tokens.assert_next(lexer.DECL_TYPE)
			tokens.eat(lexer.LCURL)
			cmp_expr := Parse_Cmp_expr(tokens)
			cmp_exprs := []*Cmp_expr{&cmp_expr}

			for tokens.peek().Type == lexer.COMMA {
				tokens.eat(lexer.COMMA)
				cmp_expr := Parse_Cmp_expr(tokens)
				cmp_exprs = append(cmp_exprs, &cmp_expr)
			}

			tokens.eat(lexer.RCURL)

			array := Array{Type: getBasicType[ident_type.Value], InitValue: cmp_exprs}

			init := Init{Type: ARRAY_INIT, Array: &array}

			return init

		}
	case lexer.MAP:
		{
		}
	case lexer.NEW:
		{
		}

	}

	return Init{}

}

func parse_Assign(tokens *Parser_Input) Assign {

	assign := Assign{}
	if tokens.peek().Type == lexer.LSQUARE || tokens.peek().Type == lexer.NEW || tokens.peek().Type == lexer.MAP {
		init := parse_Init(tokens)
		assign.Init = &init
		assign.Type = INIT_ASSIGN
	} else {
		expr := Parse_Cmp_expr(tokens)
		assign.Expr = &expr
		assign.Type = EXPR_ASSIGN
	}
	return assign
}

func Parse_Stmt(tokens *Parser_Input) Stmt {

	stmt := Stmt{}

	if !tokens.empty() {

		switch tokens.peek().Type {
		case lexer.VAR:
			{
				tokens.eat(lexer.VAR)
				ident := tokens.assert_next(lexer.IDENT)
				variable_type := parse_Type(tokens)
				tokens.eat(lexer.ASSIGN)
				assign := parse_Assign(tokens)
				decl_stmt := Decl_stmt{Id: ident.Value, Type: variable_type, Assign: &assign}
				stmt.Decl = &decl_stmt
				stmt.Type = DECL_STMT

			}
		case lexer.IDENT:
			{
				if len(tokens.Tokens) > 1 && tokens.peek2().Type == lexer.ASSIGN {
					ident := tokens.assert_next(lexer.IDENT)
					tokens.eat(lexer.ASSIGN)
					assign := parse_Assign(tokens)
					assign_stmt := Assign_stmt{Id: ident.Value, Assign: &assign}
					stmt.Assign = &assign_stmt
					stmt.Type = ASSIGN_STMT

				} else {

					expr := Parse_Arith_expr(tokens)
					stmt.Expr = &expr
					stmt.Type = EXPR

				}
			}
		case lexer.INT, lexer.DOUBLE, lexer.STRING:
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

		case lexer.RETURN:
			{
				tokens.eat(lexer.RETURN)
				expr := Parse_Cmp_expr(tokens)
				return_stmt := Return_stmt{Cmp_expr: expr}
				stmt.Return = &return_stmt
				stmt.Type = RETURN_STMT

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
