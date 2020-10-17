package lexer

import (
	"testing"

	"github.com/IAPOLINARIO/lula-language/token"
)

func TestNextToken(t *testing.T) {
	input := `pt five = 5;
				pt ten = 10;
				pt add = companheiro(x, y) {
				x + y;
				};
				pt result = add(five, ten);
				!-/*5;
				5 <`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PT, "pt"},
		{token.IDENTIFICADOR, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.PT, "pt"},
		{token.IDENTIFICADOR, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.PT, "pt"},
		{token.IDENTIFICADOR, "add"},
		{token.ASSIGN, "="},
		{token.COMPANHEIRO, "companheiro"},
		{token.LPAREN, "("},
		{token.IDENTIFICADOR, "x"},
		{token.COMMA, ","},
		{token.IDENTIFICADOR, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFICADOR, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFICADOR, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.PT, "pt"},
		{token.IDENTIFICADOR, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFICADOR, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFICADOR, "five"},
		{token.COMMA, ","},
		{token.IDENTIFICADOR, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.TERMINOU, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected %q, got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
