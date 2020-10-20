package parser

import (
	"testing"

	"github.com/IAPOLINARIO/lula-language/ast"
	"github.com/IAPOLINARIO/lula-language/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
				pt x = 5;
				pt y = 10;
				pt foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "pt" {
		t.Errorf("s.TokenLiteral not 'pt'. got=%q", s.TokenLiteral())
		return false
	}
	ptStmt, ok := s.(*ast.PtStatement)
	if !ok {
		t.Errorf("s not *ast.PtStatement. got=%T", s)
		return false
	}
	if ptStmt.Name.Value != name {
		t.Errorf("ptStmt.Name.Value not '%s'. got=%s", name, ptStmt.Name.Value)
		return false
	}
	if ptStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, ptStmt.Name)
		return false
	}
	return true
}
