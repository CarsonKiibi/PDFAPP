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
			return l.pos, TEXT, "{"
		case '}':
			return l.pos, TEXT, "}"
		case '[':
			return l.pos, TEXT, "["
		case ']':
			return l.pos, TEXT, "]"
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsLetter(r) || unicode.IsNumber(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexTextMod()
				return startPos, TEXTMOD, lit 
			}
		}
	}
	
}

func (l *Lexer) resetPosition() {
	l.pos.line++
	l.pos.column = 0
}

// need to test this
func (l *Lexer) lexTextMod() string {
	var sb strings.Builder
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return sb.String()
			}
		}
		l.pos.column++ 
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			sb.WriteRune(r)
		} else {
			return sb.String()
		}
	}
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}
	
	l.pos.column--
}

func main() {
	input := "tok1 { tok2 [ \n tok3"
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
