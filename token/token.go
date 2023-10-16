package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string // Actual value of the token: x, y, 3, let, etc
}

const (
	EOF       = "EOF"
	ILLEGAL   = "ILLEGAL"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	BANG      = "!"
	SLASH     = "/"
	ASTERISK  = "*"
	LT        = "<"
	GT        = ">"
	EQ        = "=="
	NE        = "!="
	LBRACKET  = "["
	RBRACKET  = "]"
	COLON     = ":"
	STRING    = "STRING"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
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

func IdentifierFinder(identifier string) TokenType {
	if identifierType, ok := keywords[identifier]; ok {
		return identifierType
	}
	return IDENT
}
