package main

import (
	"fmt"
	"mylexer/ast"
	"mylexer/lexer"
	"mylexer/parser"
)

func main() {

	//step-1
	//read the input
	//read character by character

	//input := "let x = 500 ;" //type Table{ Column1\nColumn2\nColumn3}
	input := `	
	let myx = another;
	`
	//myLexer(input)
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
	//fmt.Println(len(prog.Statements))
	for i, stm := range prog.Statements {
		letstm, ok := stm.(*ast.LetStatement)
		if letstm == nil {
			return
		}
		//fmt.Println(i, ok, letstm, stm.TokenLiteral(), letstm.Name.TokenLiteral(), letstm.Token)
		fmt.Println(i, ok, letstm)
	}

	// i := 0
	// for {
	// 	tok := l.NextToken()
	// 	if tok.Type == token.EOF {
	// 		break
	// 	}
	// 	fmt.Println(i, "->", tok.Type, tok.Literal)
	// 	i++
	// }

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
