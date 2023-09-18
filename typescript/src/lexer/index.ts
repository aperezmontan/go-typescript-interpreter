export const TokenType = {
    ASSIGN: "=",
    PLUS: "+",
    LPAREN: "(",
    RPAREN: ")",
    LBRACE: "{",
    RBRACE: "}",
    COMMA: ",",
    SEMICOLON: ";",
    COLON: ":",
    EOF: "",
}

type TokenItem = typeof TokenType[keyof typeof TokenType];

export type Token = {
    type: TokenItem;
    literal: string;
}

export const createToken = (type: TokenItem, literal: string): Token => {
    return { type, literal } as Token;
}

export class Tokenizer {
    private position: number = 0;
    private readPosition: number = 0;
    private character: string;

    constructor(private input: string) {
        this.input = input;
        this.readChar();
    }

    /** @throws Error */
    public getNextToken(): Token {
        let tokenItem: TokenItem;

        switch (this.character) {
            case "=":
                tokenItem = TokenType.ASSIGN;
                break;
            case "+":
                tokenItem = TokenType.PLUS;
                break;
            case "(":
                tokenItem = TokenType.LPAREN;
                break;
            case ")":
                tokenItem = TokenType.RPAREN;
                break;
            case "{":
                tokenItem = TokenType.LBRACE;
                break;
            case "}":
                tokenItem = TokenType.RBRACE;
                break;
            case ",":
                tokenItem = TokenType.COMMA;
                break;
            case ";":
                tokenItem = TokenType.SEMICOLON;
                break;
            case ":":
                tokenItem = TokenType.COLON;
                break;
            case "\0":
                this.readChar();
                return createToken(TokenType.EOF, "");
            default:
                throw new Error("Unhandled token")
        }
        
        this.readChar();
        return createToken(tokenItem, this.character);
    }

    private readChar(): void {
        if (this.readPosition >= this.input.length) {
            this.character = "\0";
        } else {
            this.character = this.input[this.readPosition];
        }

        this.position = this.readPosition;
        this.readPosition += 1;
    }
}
