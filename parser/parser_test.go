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
	checkParserErrors(t, p)

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

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
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

func TestReturnStatements(t *testing.T) {
	input := `
	devorve 5;
	devorve 10;
	devorve 993322;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	for _, stmt := range program.Statements {
		devorveStmt, ok := stmt.(*ast.DevorveStatement)
		if !ok {
			t.Errorf("stmt not *ast.devorveStatement. got=%T", stmt)
			continue
		}
		if devorveStmt.TokenLiteral() != "devorve" {
			t.Errorf("devorveStmt.TokenLiteral not 'devorve', got %q",
				devorveStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",

			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}
