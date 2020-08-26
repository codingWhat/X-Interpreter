package lexer

import (
	"github.com/codingWhat/x-interpreter/token"
)

type Lexer struct {
	input        string
	position     int   // point to position of current node
	readPosition int   // point to position of the next node  to be read.
	ch           byte  // current byte of current position.
}

func New(input string) *Lexer {
	l :=  &Lexer{
		input:input,
	}
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
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()

	return tok
}

func newToken(t token.Type, l byte) token.Token {
	return token.Token{
		Type:t,
		Literal: string(l),
	}
}