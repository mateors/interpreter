package parser

import (
	"fmt"
	"mylexer/ast"
	"mylexer/lexer"
	"mylexer/token"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)

)

type prefixParseFunction func() ast.Expression

type infixParseFunction func(ast.Expression) ast.Expression

type Parser struct {
	lexer          *lexer.Lexer
	curToken       token.Token //position
	peekToken      token.Token //read position
	errors         []string
	prefixParseFns map[token.TokenType]prefixParseFunction
	infixParseFns  map[token.TokenType]infixParseFunction
}

func New(l *lexer.Lexer) *Parser {

	p := &Parser{lexer: l}
	p.prefixParseFns = make(map[token.TokenType]prefixParseFunction)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INTEGAR, p.parseIntegerLiteral)

	p.infixParseFns = make(map[token.TokenType]infixParseFunction)

	p.NextToken()
	p.NextToken()
	return p
}

func (p *Parser) parseIntegerLiteral() ast.Expression {

	//fmt.Println(p.curToken)
	str := p.curToken.Literal
	intval, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		//fmt.Println(err)
		emsg := fmt.Sprintf("unable to parse %v as integer", str)
		p.errors = append(p.errors, emsg)
		return nil
	}
	return &ast.IntegerLiteral{Token: p.curToken, Value: intval}
}

func (p *Parser) registerPrefix(key token.TokenType, pfn prefixParseFunction) {

	p.prefixParseFns[key] = pfn
}

func (p *Parser) registerInfix(key token.TokenType, ifn infixParseFunction) {

	p.infixParseFns[key] = ifn
}

func (p *Parser) NextToken() {

	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseIdentifier() ast.Expression {
	//fmt.Println("parseIdentifier", p.curToken, p.curToken.Literal)
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) ParseProgram() *ast.Program {

	prog := &ast.Program{}
	prog.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {

		stmt := p.parseStatement()
		//fmt.Println("--", stmt)
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		p.NextToken()
	}
	return prog
}

func (p *Parser) parseStatement() ast.Statement {

	//fmt.Println("parseStatement", p.curToken)
	//if p.curToken.Type == token.LET {
	//fmt.Println(">>", p.curToken.Type, ">", p.peekToken.Type)
	//}

	switch p.curToken.Type {

	case token.LET:
		return p.parseLetStatement()

	case token.RETURN:
		return p.parseReturnStatement()

	default:
		//fmt.Println("parseStatement Nill", p.curToken, p.curToken.Type, p.curToken.Literal)
		return p.parseExpressionStatement()
	}

}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {

	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.NextToken()
	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.NextToken()
	}
	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {

	expStm := &ast.ExpressionStatement{}
	expStm.Token = p.curToken
	expStm.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.NextToken()
	}
	return expStm
}

func (p *Parser) parseExpression(precedence int) ast.Expression {

	//fmt.Println("parseExpression", p.curToken.Type)
	prefixFunction := p.prefixParseFns[p.curToken.Type]
	//fmt.Println(prefixFunction)
	if prefixFunction == nil {
		return nil
	}
	leftExp := prefixFunction()

	return leftExp
}

// func (p *Parser) parseLetStatement() *ast.LetStatement {

// 	stm := &ast.LetStatement{Token: p.curToken}
// 	//fmt.Println("parseLetStatement:", p.peekToken.Literal, p.peekToken.Type, p.curToken)
// 	if p.peekToken.Type != token.IDENT {
// 		return nil
// 	}

// 	//fmt.Println(">", p.curToken, p.peekToken) // --> x
// 	p.NextToken()
// 	stm.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
// 	//fmt.Println("?1", p.curToken)

// 	if p.peekToken.Type != token.EQS {
// 		return nil
// 	}

// 	p.NextToken()
// 	//fmt.Println("?2", p.curToken, p.peekToken) // --> 500
// 	for p.curToken.Type != token.SEMICOLON {
// 		p.NextToken()
// 	}
// 	return stm
// }

func (p *Parser) parseLetStatement() *ast.LetStatement {

	//fmt.Println("parseLetStatement:", p.curToken)
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	//fmt.Println("2>", p.curToken, p.curToken.Literal) //myx
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.EQS) {
		return nil
	}

	p.NextToken() //=
	stmt.Value = p.parseExpression(LOWEST)

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	if p.peekTokenIs(token.SEMICOLON) {
		//fmt.Println("5>", p.curToken)
		p.NextToken()
	}
	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %v, got %v instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.NextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}
