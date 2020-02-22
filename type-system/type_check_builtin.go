package type_system

import (
	"fmt"
	"github.com/KazumaTakata/static-typed-language/lexer"
	"github.com/KazumaTakata/static-typed-language/parser"
	"github.com/KazumaTakata/static-typed-language/type"
	"os"
)

func type_Check_Builtin(name string, factor parser.Factor, symbol_env *parser.Symbol_Env) basic_type.Variable_Type {

	switch name {
	case "len":
		{
			return len_typecheck(factor, symbol_env)
		}
	case "time":
		{
			return time_typecheck(factor, symbol_env)
		}
	case "print":
		{
			return print_typecheck(factor, symbol_env)
		}

	}

	return basic_type.Variable_Type{}

}

func print_typecheck(factor parser.Factor, symbol_env *parser.Symbol_Env) basic_type.Variable_Type {

	if len(factor.Args) != 1 {
		fmt.Printf("\nbuiltin func print does accept only one argument\n")
		os.Exit(1)
	}

	if factor.Args[0].Type == lexer.IDENT {
		param := resolve_name(factor.Args[0].Value, symbol_env)
		if param.Type != parser.PrimitiveType {
			fmt.Printf("\nbuiltin func print does not accept argument\n")
			os.Exit(1)

		}
		if param.Primitive.Type != basic_type.STRING {
			fmt.Printf("\nbuiltin func print does not accept argument of type of %v \n", param.Primitive.Type)
			os.Exit(1)

		}
	} else {
		if factor.Args[0].Type != lexer.STRING {
			fmt.Printf("\nbuiltin func print  accept argument of type of string \n")
			os.Exit(1)

		}
	}

	return basic_type.Variable_Type{}
}

func time_typecheck(factor parser.Factor, symbol_env *parser.Symbol_Env) basic_type.Variable_Type {

	if len(factor.Args) != 0 {
		fmt.Printf("\nbuiltin func time does not accept argument\n")
		os.Exit(1)
	}

	return basic_type.IntPrimitiveType
}

func len_typecheck(factor parser.Factor, symbol_env *parser.Symbol_Env) basic_type.Variable_Type {

	if len(factor.Args) != 1 {
		fmt.Printf("\nbuiltin func len only accept one argument\n")
		os.Exit(1)
	}

	for _, param := range factor.Args {
		if param.Type == lexer.IDENT {
			param_object := resolve_name(param.Value, symbol_env)
			if param_object.Type != parser.ArrayType {
				fmt.Printf("\nbuiltin func len only accept array type\n")
				os.Exit(1)

			}
		}

	}
	return basic_type.IntPrimitiveType
}
