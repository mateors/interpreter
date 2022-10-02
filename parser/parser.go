package parser

import (
	"fmt"
	"mylexer/ast"
	"mylexer/lexer"
	"mylexer/token"
)

const (
	_ int = iota
	LOWEST
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  token.Token //position
	peekToken token.Token //read position
	errors    []string
}

func New(l *lexer.Lexer) *Parser {

	p := &Parser{lexer: l}
	p.NextToken()
	p.NextToken()
	return p
}

func (p *Parser) NextToken() {

	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
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
		//p.parseExpressionStatement()
		return nil
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
	expStm.Experssion = p.parseExpression(LOWEST)
	return expStm
}

func (p *Parser) parseExpression(precedence int) ast.Expression {

	//exp := &ast.Experssion{}
	return nil
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
