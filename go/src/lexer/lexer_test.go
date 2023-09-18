package lexer

import (
    "testing"

    "orca/src/token"
)

func TestNextToken(test *testing.T) {
    input := `=+(){},;`

    tests := []struct {
        expectedType    token.TokenType
        expectedLiteral string
    }{
        {token.ASSIGN, "="},
        {token.PLUS, "+"},
        {token.LPAREN, "("},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.RBRACE, "}"},
        {token.COMMA, ","},
        {token.SEMICOLON, ";"},
        {token.EOF, ""},
    }

    lexer := New(input)

    for index, testToken := range tests {
        tok := lexer.NextToken()

        if tok.Type != testToken.expectedType {
            test.Fatalf("tests[%d] - tokentype wrong. expected=%q, got %q",
            index, testToken.expectedType, tok.Type)
        }

        if tok.Literal != testToken.expectedLiteral {
            test.Fatalf("tests[%d] - literal wrong. expected %q, got %q",
            index, testToken.expectedLiteral, tok.Literal)
        }
    }
}
