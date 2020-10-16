package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	NAO_PODE = "ILLEGAL"
	TERMINOU = "EOF"

	//Identifier and literals
	IDENT = "IDENTIFICADOR"
	INT   = "INT"

	//Operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	PT       = "LET"
)
