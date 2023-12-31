package ast

import (
	"asm/token"
)

type Node interface {
	TokenLiteral() string
}

// Statement represents statement node in AST.
type Statement interface {
	Node
	statementNode()
}

// Expression represents expression node in AST.
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
	}
	return ""
}

type Opcode struct {
	Token    token.Token
	Operands []Expression
}

func (o *Opcode) statementNode() {}
func (o *Opcode) TokenLiteral() string {
	return o.Token.Literal
}

type Label struct {
	Token token.Token
}

func (l *Label) statementNode() {}
func (l *Label) TokenLiteral() string {
	return l.Token.Literal
}

type LabelRef struct {
	Token token.Token
}

func (lr *LabelRef) expressionNode() {}
func (lr *LabelRef) TokenLiteral() string {
	return lr.Token.Literal
}

type NumberLiteral struct {
	Token token.Token
	Value byte
}

func (nl *NumberLiteral) expressionNode() {}
func (nl *NumberLiteral) TokenLiteral() string {
	return nl.Token.Literal
}
