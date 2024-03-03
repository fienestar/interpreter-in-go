package token

import (
	"monkey/types"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal types.InputString
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN      = "="
	ADD         = "+"
	SUB         = "-"
	DIV         = "/"
	MUL         = "*"
	LOGICAL_NOT = "!"
	EQ          = "=="
	NOT_EQ      = "!="
	GREATER     = ">"
	GREATER_EQ  = ">="
	LESS        = "<"
	LESS_EQ     = "<="

	TRUE  = "TRUE"
	FALSE = "FALSE"

	IF   = "IF"
	ELSE = "ELSE"

	RETURN = "RETURN"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"

	STRING = "STRING"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident types.InputString) TokenType {
	if tok, ok := keywords[string(ident)]; ok {
		return tok
	}
	return IDENT
}
