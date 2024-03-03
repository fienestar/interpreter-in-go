package lexer

import (
	"monkey/token"
	"monkey/types"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
let result = add(five, ten);
!-/*5;
5 < 10 > 5;
let 🌌 = "안녕하세요";
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.ADD, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.LOGICAL_NOT, "!"},
		{token.SUB, "-"},
		{token.DIV, "/"},
		{token.MUL, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LESS, "<"},
		{token.INT, "10"},
		{token.GREATER, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "🌌"},
		{token.ASSIGN, "="},
		{token.STRING, "안녕하세요"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(types.InputString(input))

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=[%q, %q] got=[%q, %q]",
				i, tt.expectedType, tt.expectedLiteral, tok.Type, string(tok.Literal))
		}

		if string(tok.Literal) != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=[%q, %q] got=[%q, %q]",
				i, tt.expectedType, tt.expectedLiteral, tok.Type, string(tok.Literal))
		}
	}
}
