package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"companheiro": COMPANHEIRO,
	"pt":          PT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFICADOR
}

const (
	NAO_PODE = "NAO PODE"
	TERMINOU = "TERMINOU"

	//Identifier and literals
	IDENTIFICADOR = "IDENTIFICADOR"
	INT           = "INT"

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
	COMPANHEIRO = "COMPANHEIRO"
	PT          = "pt"
)
