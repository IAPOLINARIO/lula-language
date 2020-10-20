package ast

import (
	"github.com/IAPOLINARIO/lula-language/token"
)

type Node interface {
	TokenLiteral() string
}
type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type PtStatement struct {
	Token token.Token // the token.PT token
	Name  *Identifier
	Value Expression
}

func (ls *PtStatement) statementNode()       {}
func (ls *PtStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENTIFICADOR token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type DevorveStatement struct {
	Token        token.Token // the 'devorve' token
	DevorveValue Expression
}

func (rs *DevorveStatement) statementNode()       {}
func (rs *DevorveStatement) TokenLiteral() string { return rs.Token.Literal }
