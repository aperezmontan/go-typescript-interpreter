package lexer

import (
    "orca/src/token"
)

type Lexer struct {
    input           string
    position        int
    readPosition    int
    character       byte
}

func New(input string) *Lexer {
    lexer := &Lexer{input: input}
    lexer.readChar()
    return lexer
}

func (lexer *Lexer) readChar() {
    if lexer.readPosition >= len(lexer.input) {
        lexer.character = 0
    } else {
        lexer.character = lexer.input[lexer.readPosition]
    }

    lexer.position = lexer.readPosition
    lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
    var tok token.Token

    switch lexer.character {
    case '=':
        tok = newToken(token.ASSIGN, lexer.character)
    case ';':
        tok = newToken(token.SEMICOLON, lexer.character)
    case '(':
        tok = newToken(token.LPAREN, lexer.character)
    case ')':
        tok = newToken(token.RPAREN, lexer.character)
    case '{':
        tok = newToken(token.LBRACE, lexer.character)
    case '}':
        tok = newToken(token.RBRACE, lexer.character)
    case ',':
        tok = newToken(token.COMMA, lexer.character)
    case '+':
        tok = newToken(token.PLUS, lexer.character)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    }

    lexer.readChar()
    return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(char)}
}

