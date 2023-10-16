package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string // Actual value of the token: x, y, 3, let, etc
}

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	FUNCTION = "FUNCTION"
	LET      = "LET"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	TRUE  = "TRUE"
	FALSE = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":    FUNCTION,
	"let":   LET,
	"true":  TRUE,
	"false": FALSE,
}

func IdentifierFinder(identifier string) TokenType {
	if identifierType, ok := keywords[identifier]; ok {
		return identifierType
	}
	return IDENT
}
