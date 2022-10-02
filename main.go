package main

import (
	"fmt"
	"mylexer/ast"
	"mylexer/lexer"
	"mylexer/parser"
)

func LetOrExpressStatementManual() {

	input := "foobar;" //let x = myfunc ; type Table{ Column1\nColumn2\nColumn3}
	l := lexer.New(input)
	// for {
	// 	tok := l.NextToken()
	// 	if tok.Type == token.EOF {
	// 		break
	// 	}
	// 	fmt.Println(tok)
	// }
	//os.Exit(1)
	p := parser.New(l)
	prog := p.ParseProgram()
	fmt.Println(len(prog.Statements))
	for i, stm := range prog.Statements {

		letstm, ok := stm.(*ast.LetStatement)
		if ok {
			fmt.Println(i, ok, letstm)
			return
		}

		stmt, ok := stm.(*ast.ExpressionStatement)
		if ok {
			ident, ok := stmt.Expression.(*ast.Identifier)
			fmt.Println(i, ok, ident)
		}

	}

}

func ReturnStatementManul() {

	input := `	
	return 5; 
	`
	l := lexer.New(input)
	// for {
	// 	tok := l.NextToken()
	// 	if tok.Type == token.EOF {
	// 		break
	// 	}
	// 	fmt.Println(tok)
	// }
	//os.Exit(1)

	p := parser.New(l)

	prog := p.ParseProgram()
	fmt.Println(len(prog.Statements))
	for i, stm := range prog.Statements {

		letstm, ok := stm.(*ast.ReturnStatement)
		if letstm == nil {
			return
		}
		//fmt.Println(i, ok, letstm, stm.TokenLiteral(), letstm.Name.TokenLiteral(), letstm.Token)
		fmt.Println(i, ok, letstm)
	}
}

func main() {

	//step-1
	//read the input
	//read character by character
	LetOrExpressStatementManual()
	//ReturnStatementManul()

}

func myLexer(input string) {

	fmt.Println(rune('অ'), byte('a'))
	fmt.Printf("%v %T\n", 'a', 'a')
	fmt.Printf("%v %T\n", "a", "a")
	fmt.Printf("%v %T\n", byte('a'), byte('a'))
	fmt.Printf("%v %T\n", []byte("abc"), []byte("a"))
	fmt.Printf("%v\n", fmt.Sprintf("%c", byte('a')))
	fmt.Printf("%v\n", fmt.Sprintf("%c", 2437))
	fmt.Println(string(97), string(2437))   //takes single byte or rune
	fmt.Println(string([]byte{97, 98, 99})) //takes byte slice

	// rslc := []rune{}
	// for i, r := range input {
	// 	fmt.Printf("%d %v = %T %s\n", i, r, r, rune2str(r))
	// 	rslc = append(rslc, r)
	// }

	// fmt.Println(rslc)
}

func rune2str(r rune) string {
	return fmt.Sprintf("%c", r)
}
