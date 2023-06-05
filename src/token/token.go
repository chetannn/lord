package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"
    
    // identifier + literal
    IDENT = "IDENT"
    INT = "INT"

    // Operators
    ASSIGN = "="
    PLUS = "+"
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"

    LT = "<"
    GT = "<"

    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
    TRUE = "true"
    FALSE = "false"

)

var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
    "TRUE": TRUE,
    "FALSE": FALSE,
}

func LookupIdent(ident string) TokenType {

    if tok, ok := keywords[ident]; ok {
        return tok
    }

    return IDENT
}
