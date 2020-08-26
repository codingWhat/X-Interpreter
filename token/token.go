package token

type Type string

type Token struct {
	Type    Type
	Literal string
}



const  (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"  // tell parser where to stop

	//identifiers + literal
	IDENT   = "IDENT"
	INT     = "INT"

	//Operators
	ASSIGN  = "="
	PLUS    = "+"

	//Delimiters
	COMMA   = ","
	SEMICOLON = ";"


	LPAREN  = "("
	RPAREN  = ")"
	LBRACE  = "{"
	RBRACE  = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET    = "LET"
)