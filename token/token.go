package token

const (
	ILLEGAL TokenType = iota
	LBR
	RBR
	EQS
	//NEWLINE
	LET
	RETURN
	IDENT
	INTEGAR   //digit
	SEMICOLON //;
	PLUS      //+
	MINUS     //-
	ASTERISK  //*
	SLASH     //
	BANG      //!
	EOF
)

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"let":    LET,
	"return": RETURN,
}

func LookupIdentifier(key string) TokenType {

	if val, isFound := keywords[key]; isFound {
		return val
	}
	return IDENT
}

/*
mutation createMovie {

	createMovie(
	  input: {
		title: "Rise of GraphQL Warrior Pt1"
		url: "https://riseofgraphqlwarriorpt1.com/"
	  }
	){
	  id
	}

}
*/

/*
query getMovies {

 movies {
     title
     url
     releaseDate
 }

}
*/

// var tokens = [...]string{
// 	ILLEGAL: "ILLEGAL",

// 	EOF:     "EOF",
// 	COMMENT: "COMMENT"Println

// 	IDENT:  "IDENT",
// 	INT:    "INT",
// 	FLOAT:  "FLOAT",
// 	IMAG:   "IMAG",
// 	CHAR:   "CHAR",
// 	STRING: "STRING",
// }

/*
type Table{
	Name
	ID
}
*/
