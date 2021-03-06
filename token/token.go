package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"companheiro": COMPANHEIRO,
	"pt":          PT,
	"verdade":     VERDADE,
	"mentira":     MENTIRA,
	"se":          SE,
	"senao":       SENAO,
	"devorve":     DEVORVE,
}

func ProcuraIdentificador(identificador string) TokenType {
	if token, ok := keywords[identificador]; ok {
		return token
	}

	return IDENTIFICADOR
}

const (
	NAO_PODE = "NAO PODE"
	TERMINOU = "TERMINOU"

	//Identifier and literals
	IDENTIFICADOR = "IDENTIFICADOR"
	INT           = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

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

	VERDADE = "VERDADE"
	MENTIRA = "MENTIRA"
	SE      = "SE"
	SENAO   = "SENAO"
	DEVORVE = "DEVORVE"
)
