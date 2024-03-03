package lexer

import (
	"monkey/token"
	"monkey/types"
)

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch { // TODO: ==, !=, >=, <=
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.ADD, l.ch)
	case '-':
		tok = newToken(token.SUB, l.ch)
	case '/':
		tok = newToken(token.DIV, l.ch)
	case '*':
		tok = newToken(token.MUL, l.ch)
	case '!':
		tok = newToken(token.LOGICAL_NOT, l.ch)
	case '>':
		tok = newToken(token.GREATER, l.ch)
	case '<':
		tok = newToken(token.LESS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '"':
		literal, ok := l.readString()
		tok.Literal = literal
		if ok {
			tok.Type = token.STRING
		} else {
			tok.Type = token.ILLEGAL
		}
	case 0:
		tok.Literal = types.InputString("")
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func isDigit(ch types.InputChar) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readString() (types.InputString, bool) {
	position := l.position + 1
	var before types.InputChar = ' '
	for {
		l.readChar()
		if (before != '\\' && l.ch == '"') || l.ch == 0 {
			break
		}
		if l.ch == '\\' && before == '\\' {
			before = ' '
		} else {
			before = types.InputChar(l.ch)
		}
	}
	return l.input[position:l.position], l.ch != 0
}

func (l *Lexer) readNumber() types.InputString {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() types.InputString {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch types.InputChar) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= 128
}

func newToken(tokenType token.TokenType, ch types.InputChar) token.Token {
	return token.Token{Type: tokenType, Literal: types.ToInputString(ch)}
}

type Lexer struct {
	input        types.InputString
	position     int
	readPosition int
	ch           types.InputChar
}

func New(input types.InputString) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
