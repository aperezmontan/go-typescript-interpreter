import { TokenType, Tokenizer } from '../../lexer';


test("test getNextToken()", () => {
    const input = `=+(){},;`;

    const tokens = [
        TokenType.ASSIGN,
        TokenType.PLUS,
        TokenType.LPAREN,
        TokenType.RPAREN,
        TokenType.LBRACE,
        TokenType.RBRACE,
        TokenType.COMMA,
        TokenType.SEMICOLON,
        TokenType.EOF
    ]

    const lexer = new Tokenizer(input);

    tokens.forEach((token) => {
        expect(lexer.getNextToken().type).toBe(token);
    })
})
