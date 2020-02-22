package type_system

import (
	"fmt"
	"github.com/KazumaTakata/regex_virtualmachine"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"
	basic_type "github.com/KazumaTakata/static-typed-language/type"
	"io/ioutil"
	"os"
)

func Type_Check_Stmts(stmts []parser.Stmt, symbol_env *parser.Symbol_Env) {

	for _, stmt := range stmts {
		Type_Check_Stmt(stmt, symbol_env)
	}

}

func Type_Check_Assign(assign *parser.Assign, symbol_env *parser.Symbol_Env) basic_type.Variable_Type {

	if assign.Type == parser.INIT_ASSIGN {
		switch assign.Init.Type {
		case parser.ARRAY_INIT:
			{
				arrayelement_type := assign.Init.Array.ElementType
				for _, init_value := range assign.Init.Array.InitValue {
					assign_type := Type_Check_Assign(init_value, symbol_env)
					if !basic_type.Variable_Equal(arrayelement_type, assign_type) {

						fmt.Printf("array type %+v  mismatch to initialization type \n%+v\n", arrayelement_type, assign_type)
						os.Exit(1)
					}
				}

				return basic_type.WrapWithArrayType(arrayelement_type)
			}
		case parser.MAP_INIT:
			{
			}

		}
	} else if assign.Type == parser.EXPR_ASSIGN {

		cmp_type := Type_Check_Cmp(assign.Expr, symbol_env)

		return cmp_type
	}

	return basic_type.Variable_Type{}

}

func get_Array_Element_Type(nest int, array parser.ArrayObj) basic_type.Variable_Type {
	elementtype := array.ElementType
	nest -= 1

	for nest != 0 {
		elementtype = elementtype.Array.ElementType
		nest -= 1
	}

	return elementtype

}

func type_Check_Decl(decl *parser.Decl_stmt, symbol_env *parser.Symbol_Env) {
	var_type := decl.Type
	assign_type := Type_Check_Assign(decl.Assign, symbol_env)

	if var_type.DataStructType != assign_type.DataStructType {
		fmt.Printf("data structure mismatch %+v: %+v\n", var_type.DataStructType, assign_type.DataStructType)
		os.Exit(1)

	}

	switch var_type.DataStructType {
	case basic_type.PRIMITIVE:
		{
			primitive := parser.PrimitiveObj{Type: var_type.Primitive.Type}
			symbol_env.Table[decl.Id] = parser.Object{Type: parser.PrimitiveType, Primitive: &primitive}
		}
	case basic_type.ARRAY:
		{
			array := parser.ArrayObj{ElementType: var_type.Array.ElementType}
			symbol_env.Table[decl.Id] = parser.Object{Type: parser.ArrayType, Array: &array}
		}

	}

}

func type_Check_Assign(assign *parser.Assign_stmt, symbol_env *parser.Symbol_Env) {
	variable_type := Type_Check_Assign(assign.Assign, symbol_env)
	object := resolve_name(assign.Id, symbol_env)

	switch object.Type {
	case parser.ArrayType:
		{

			for i, _ := range assign.Indexs {
				index_type := Type_Check_Arith(&assign.Indexs[i], symbol_env)
				if index_type.Primitive.Type != basic_type.INT {
					fmt.Printf("index type not int")
					os.Exit(1)
				}
			}
			if len(assign.Indexs) > 0 {
				number_of_nest := len(assign.Indexs)
				arrayelementtype := get_Array_Element_Type(number_of_nest, *object.Array)

				if !basic_type.Variable_Equal(arrayelementtype, variable_type) {
					fmt.Printf("\nassignment type mismatch %+v:%+v\n", variable_type, arrayelementtype)
					os.Exit(1)

				}
			} else {
				if !basic_type.Variable_Equal(basic_type.WrapWithArrayType(object.Array.ElementType), variable_type) {
					fmt.Printf("\nassignment type mismatch %+v:%+v\n", basic_type.WrapWithArrayType(object.Array.ElementType), variable_type)
					os.Exit(1)
				}
			}

		}
	case parser.PrimitiveType:
		{

			if variable_type.DataStructType != basic_type.PRIMITIVE {
				fmt.Printf("data structure mismatch not primitive\n")
				os.Exit(1)
			}

			if object.Primitive.Type != variable_type.Primitive.Type {
				fmt.Printf("primitive value can not assigned to %+v  varieble\n", object.Type)
				os.Exit(1)
			}
		}

	}

}

func Type_Check_Stmt(stmt parser.Stmt, symbol_env *parser.Symbol_Env) {
	switch stmt.Type {
	case parser.EXPR:
		{
			_ = Type_Check_Arith(stmt.Expr, symbol_env)

		}

	case parser.IMPORT_STMT:
		{

			regex_string := lexer.Get_Regex_String()

			regex := regex.NewRegexWithParser(regex_string)

			module_symbol_env := parser.Symbol_Env{Table: parser.Symbol_Table{}}

			dat, _ := ioutil.ReadFile(stmt.Import.Module_name + ".cat")
			string_input := string(dat)

			tokens := lexer.GetTokens(regex, string_input)
			parser_input := parser.Parser_Input{Tokens: tokens, Pos: 0}
			stmts := parser.Parse_Stmts(&parser_input)

			Type_Check_Stmts(stmts, &module_symbol_env)

			symbol_env.Table[stmt.Import.Module_name] = parser.Object{Type: parser.EnvType, Env: &module_symbol_env}

		}
	case parser.ASSIGN_STMT:
		{
			type_Check_Assign(stmt.Assign, symbol_env)
		}

	case parser.DECL_STMT:
		{
			type_Check_Decl(stmt.Decl, symbol_env)

		}
	case parser.FOR_STMT:
		{
			switch stmt.For.Type {
			case parser.Cmp:
				{

					_ = Type_Check_Cmp(&stmt.For.Cmp_expr, symbol_env)

					if !basic_type.Variable_Equal(stmt.For.Cmp_expr.Type, basic_type.BoolPrimitiveType) {
						fmt.Printf("if conditional expression should return bool type: return %+v\n", stmt.For.Cmp_expr.Type)
						os.Exit(1)
					}

					Child_env := &parser.Symbol_Env{Table: parser.Symbol_Table{}, Parent_Env: symbol_env}
					stmt.For.Symbol_Env = Child_env
					Type_Check_Stmts(stmt.For.Stmts, Child_env)
				}
			case parser.DeclCmpAssign:
				{
					type_Check_Decl(&stmt.For.Decl, symbol_env)
					type_Check_Assign(&stmt.For.Assign, symbol_env)

					_ = Type_Check_Cmp(&stmt.For.Cmp_expr, symbol_env)

					if !basic_type.Variable_Equal(stmt.For.Cmp_expr.Type, basic_type.BoolPrimitiveType) {
						fmt.Printf("if conditional expression should return bool type: return %+v\n", stmt.For.Cmp_expr.Type)
						os.Exit(1)
					}

					Child_env := &parser.Symbol_Env{Table: parser.Symbol_Table{}, Parent_Env: symbol_env}
					stmt.For.Symbol_Env = Child_env
					Type_Check_Stmts(stmt.For.Stmts, Child_env)

				}
			}
		}
	case parser.IF_STMT:
		{
			_ = Type_Check_Cmp(&stmt.If.Cmp_expr, symbol_env)

			if !basic_type.Variable_Equal(stmt.If.Cmp_expr.Type, basic_type.BoolPrimitiveType) {
				fmt.Printf("if conditional expression should return bool type: return %+v\n", stmt.If.Cmp_expr.Type)
				os.Exit(1)
			}

			Child_env := &parser.Symbol_Env{Table: parser.Symbol_Table{}, Parent_Env: symbol_env}
			stmt.If.Symbol_Env = Child_env
			Type_Check_Stmts(stmt.If.Stmts, Child_env)

		}

	case parser.RETURN_STMT:
		{
			//_ = Type_Check_Arith(&stmt.Return.Cmp_expr.Ariths[0].Arith, symbol_env)
			_ = Type_Check_Cmp(&stmt.Return.Cmp_expr, symbol_env)

			stmt.Return.Type = stmt.Return.Cmp_expr.Type

			if !basic_type.Variable_Equal(symbol_env.Return_Type, stmt.Return.Type) {
				fmt.Printf("func return type mismatch :expect %+v, got%+v\n", symbol_env.Return_Type.Primitive, stmt.Return.Type.Primitive)
				os.Exit(1)

			}

		}
	case parser.DEF_STMT:
		{

			function := parser.Object{Type: parser.FunctionType, Function: stmt.Def}
			symbol_env.Table[stmt.Def.Id] = function

			Child_env := &parser.Symbol_Env{Table: parser.Symbol_Table{}, Parent_Env: symbol_env}

			for _, arg := range stmt.Def.Args {

				switch arg.Type.DataStructType {
				case basic_type.PRIMITIVE:
					{
						primitive := parser.PrimitiveObj{Type: arg.Type.Primitive.Type}
						Child_env.Table[arg.Ident] = parser.Object{Type: parser.PrimitiveType, Primitive: &primitive}
					}
				case basic_type.ARRAY:
					{
						array := parser.ArrayObj{ElementType: arg.Type.Array.ElementType}
						Child_env.Table[arg.Ident] = parser.Object{Type: parser.ArrayType, Array: &array}
					}
				}
			}

			Child_env.Return_Type = function.Function.Return_type
			stmt.Def.Symbol_Env = Child_env
			Type_Check_Stmts(stmt.Def.Stmts, Child_env)

		}

	}
}
