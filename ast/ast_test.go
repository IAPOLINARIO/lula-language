package ast

import (
	"testing"

	"github.com/IAPOLINARIO/lula-language/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&PtStatement{
				Token: token.Token{Type: token.PT, Literal: "pt"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENTIFICADOR, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENTIFICADOR, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "pt myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
