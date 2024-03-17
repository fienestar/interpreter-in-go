package ast

import (
	"monkey/token"
	"monkey/types"
)

type Node interface {
	TokenLiteral() types.InputString
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

func (p *Program) TokenLiteral() types.InputString {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return types.InputString{}
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()                  {}
func (ls *LetStatement) TokenLiteral() types.InputString { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value types.InputString
}

func (i *Identifier) expressionNode()                 {}
func (i *Identifier) TokenLiteral() types.InputString { return i.Token.Literal }
