package ast

import (
	"bytes"
	"mylexer/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (a *LetStatement) TokenLiteral() string {
	return a.Token.Literal
}

func (a *LetStatement) statementNode() {
}

func (a *LetStatement) String() string {

	var out bytes.Buffer

	out.WriteString(a.TokenLiteral() + " ")
	out.WriteString(a.Name.String()) //identifier
	out.WriteString(" = ")

	if a.Value != nil {
		out.WriteString(a.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string //or Value
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {
	//fmt.Println("expressionNode:", i)
}

func (i *Identifier) String() string {
	return i.Value
}

type ExpressionStatement struct {
	Token      token.Token
	Experssion Expression
}

func (es *ExpressionStatement) statementNode() {
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {

	if es.Experssion != nil {
		return es.Experssion.String()
	}
	return ""
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	return ""
}