package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
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
	EOF: 	"EOF",
	ILLEGALNEST: "ILLEGALNEST",
	ILLEGAL: "ILLEGAL",
	SPACING: "SPACING",
	TEXTMOD: "TEXTMOD",
	TEXT: "TEXT",
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
		l.pos.column++
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}
			panic(err)
		}

		switch r {
		case '\\':
			// found it!
		case '\n':
			l.resetPosition()
		case '{':
			startPos := l.pos
			lit, tok := l.lexText(TEXTMOD)
			return startPos, tok, lit 
		case '}':
			return l.pos, TEXT, "}"
		case '[':
			startPos := l.pos 
			lit, tok := l.lexText(SPACING)
			return startPos, tok, lit
		case ']':
			return l.pos, TEXT, "]"
		case ' ':
			return l.pos, TEXT, "_"
		default:
			startPos := l.pos
			l.backup()
			lit, tok := l.lexText(TEXT)
			return startPos, tok, lit
		}
	}
	
}

func (l *Lexer) resetPosition() {
	l.pos.line++
	l.pos.column = 0
}

// need to test this
func (l *Lexer) lexText(tokenType int) (string, Token) {
	var sb strings.Builder
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return sb.String(), EOF
			}
		}
		l.pos.column++ 
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			sb.WriteRune(r)

		// TEXT MOD
		} else if tokenType == TEXTMOD {
			if r == '{' {
				return sb.String(), ILLEGALNEST
			} else if r == '}' {
				return sb.String(), TEXTMOD
			} else {
				sb.WriteRune(r)
			}

		// TEXT
		} else if tokenType == SPACING {
			if r == '[' {
				return sb.String(), ILLEGALNEST 
			} else if r == ']' {
				return sb.String(), SPACING
			} else {
				sb.WriteRune(r)
			}

		} else if tokenType == TEXT {
			if unicode.IsSpace(r) {
				l.backup() // space between one and two doesnt show but after two does (when removed)
				return sb.String(), TEXT
			} else {
				sb.WriteRune(r)
			}
		}
	}
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
