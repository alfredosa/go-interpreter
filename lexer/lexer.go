package lexer

import (
	"github.com/alfredosa/go-interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int // going rogue here (not in the book)
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1}
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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace() // In Monkey - we don't care about whitespaces

	switch l.ch {
	case '=':
		if l.nextChar() == '=' {
			firstChar := l.ch
			l.readChar()
			secondChar := l.ch
			literal := string(firstChar) + string(secondChar)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
		return tok
	case '!':
		if l.nextChar() == '=' {
			firstChar := l.ch
			l.readChar()
			secondChar := l.ch
			literal := string(firstChar) + string(secondChar)
			tok = token.Token{Type: token.NE, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if IsLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.IdentifierFinder(tok.Literal)
			return tok
		} else if isNumber(l.ch) {
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

func (l *Lexer) readString() string {
	l.readChar() // read the opening double quote
	currentPosition := l.position
	for readDoubleQuote(l.ch) {
		l.readChar()
	}

	sequence := l.input[currentPosition:l.position] // up until that point is the string
	l.readChar()                                    // skip the closing double quote

	return sequence
}

func (l *Lexer) readIdentifier() string {
	currentPosition := l.position
	for IsLetter(l.ch) {
		l.readChar()
	}
	return l.input[currentPosition:l.position]
}

func (l *Lexer) readNumber() string {

	currentPosition := l.position
	for isNumber(l.ch) {
		l.readChar()
	}
	return l.input[currentPosition:l.position]
}

func readDoubleQuote(ch byte) bool {
	return ch != '"'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) Line() int {
	return l.line
}

func (l *Lexer) IncrementLine() {
	l.line += 1
}

func (l *Lexer) nextChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}
