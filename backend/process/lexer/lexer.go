package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Token int

const (
	EOF = iota
	ILLEGALNEST
	ILLEGAL
	SPACING
	TEXTMOD
	TEXT
)

var tokens = []string{
	EOF:         "EOF",
	ILLEGALNEST: "ILLEGALNEST",
	ILLEGAL:     "ILLEGAL",
	SPACING:     "SPACING",
	TEXTMOD:     "TEXTMOD",
	TEXT:        "TEXT",
}

type Position struct {
	line   int
	column int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func (t Token) String() string {
	return tokens[t]
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, column: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (Position, Token, string) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}
			panic(err)
		}
		l.pos.column++

		switch r {
		case '\\':
			// found it!
		case '\n':
			l.resetPosition()
		case '{':
			startPos := l.pos
			lit, tok := l.lexText(TEXTMOD)
			if tok == EOF {
				return startPos, EOF, lit
			}
			return startPos, tok, lit
		case '}':
			return l.pos, TEXT, "}"
		case '[':
			startPos := l.pos
			lit, tok := l.lexText(SPACING)
			if tok == EOF {
				return startPos, EOF, lit
			}
			return startPos, tok, lit
		case ']':
			return l.pos, TEXT, "]"
		case ' ':
			return l.pos, SPACING, "_"
		default:
			startPos := l.pos
			l.backup()
			lit, tok := l.lexText(TEXT)
			if tok == EOF {
				return startPos, EOF, lit
			}
			return startPos, tok, lit
		}
	}
}

func (l *Lexer) lexText(tokenType int) (string, Token) {
	var sb strings.Builder
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// If we've accumulated any text, return it before EOF
				if sb.Len() > 0 {
					return sb.String(), Token(tokenType)
				}
				return sb.String(), EOF
			}
			panic(err)
		}
		l.pos.column++

		switch tokenType {
		case TEXTMOD:
			if r == '{' {
				l.backup()
				return sb.String(), ILLEGALNEST
			} else if r == '}' {
				l.backup()
				return sb.String(), TEXTMOD
			} else {
				sb.WriteRune(r)
			}
		case SPACING:
			if r == '[' {
				l.backup()
				return sb.String(), ILLEGALNEST
			} else if r == ']' {
				l.backup()
				return sb.String(), SPACING
			} else {
				sb.WriteRune(r)
			}
		case TEXT:
			if r == ' ' || r == '{' || r == '}' || r == '[' || r == ']' {
				l.backup()
				return sb.String(), TEXT
			} else {
				sb.WriteRune(r)
			}
		}
	}
}

func (l *Lexer) resetPosition() {
	l.pos.line++
	l.pos.column = 0
}
func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.column--
}

// maybe just pass some variable into lexText etc that makes it ignore stuff
func (l *Lexer) ignoreNext() {
	// ??
}

func main() {
	input := "text {mod}[spacing] second text text what what"
	reader := strings.NewReader(input)
	lexer := NewLexer(reader)
	for {
		pos, tok, lit := lexer.Lex()
		if tok == EOF {
			break
		}

		fmt.Printf("%d:%d | %s | %s | \n", pos.line, pos.column, tok, lit)
	}
}
