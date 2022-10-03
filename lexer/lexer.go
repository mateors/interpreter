package lexer

import (
	"mylexer/token"
)

// or Scanner
type Lexer struct {
	input       string
	read        int //
	currentRead int
	chr         byte
}

func New(in string) *Lexer {

	l := &Lexer{input: in}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.currentRead >= len(l.input) {
		l.chr = 0
		return
	}

	//fmt.Println(string(l.chr), l.chr)
	l.chr = l.input[l.currentRead]
	l.read = l.currentRead
	l.currentRead += 1
}

func (l *Lexer) skipWhitespace() {

	for l.chr == ' ' || l.chr == '\t' || l.chr == '\r' || l.chr == '\n' {
		//for l.chr == ' ' || l.chr == '\t' || l.chr == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {

	var tok token.Token
	l.skipWhitespace()

	switch l.chr {

	case '=':
		tok = token.Token{Type: token.EQS, Literal: string(l.chr)}

	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(l.chr)}

	case '-':
		tok = token.Token{Type: token.MINUS, Literal: string(l.chr)}

	case '!':
		tok = token.Token{Type: token.BANG, Literal: string(l.chr)}

	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: string(l.chr)}

	//case '\n':
	//tok = token.Token{Type: token.NEWLINE, Literal: string(l.chr)}

	case 0:
		tok = token.Token{Type: token.EOF, Literal: string(l.chr)}

	default:

		if isLetter(l.chr) {

			// position := l.read
			// for isLetter(l.chr) {
			// 	l.readChar2()
			// }
			// if l.chr == ' ' {
			// 	fmt.Println("SPACE")
			// }
			literal := l.identifier()
			tokType := token.LookupIdentifier(literal)
			//fmt.Println("**", literal, tokType)
			//return token.Token{Type: tokType, Literal: literal}
			tok.Type = tokType
			tok.Literal = literal
			return tok

		} else if isDigit(l.chr) {

			// position := l.read
			// for isDigit(l.chr) {
			// 	l.readChar()
			// }
			digit := l.integar()
			tok.Literal = digit
			tok.Type = token.INTEGAR
			return tok

		} else {

			//fmt.Printf("# %v %T\n", l.chr, l.chr)
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.chr)}

		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) identifier() string {

	position := l.read
	for isLetter(l.chr) {
		l.readChar()
	}
	// if l.chr == ' ' {
	// 	fmt.Println("SPACE")
	// }
	//fmt.Println(">>", l.input[position:l.currentRead])
	return l.input[position:l.read]
}

func (l *Lexer) integar() string {

	position := l.read
	for isDigit(l.chr) {
		l.readChar()
	}
	//fmt.Println(position, l.read, l.currentRead)
	return l.input[position:l.read]
}

func isLetter(ch byte) bool {

	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {

	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() byte {

	if l.currentRead >= len(l.input) {
		return 0
	}
	return l.input[l.currentRead]
}
