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
	SPACE
	SPACEMOD
	TEXTMOD
	TEXT
)

var tokens = []string{
	EOF:         "EOF",
	ILLEGALNEST: "ILLEGALNEST",
	ILLEGAL:     "ILLEGAL",
	SPACE:     "SPACE",
	SPACEMOD: "SPACEMOD",
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
			lit, tok := l.lexText(SPACEMOD)
			if tok == EOF {
				return startPos, EOF, lit
			}
			return startPos, tok, lit
		case ']':
			return l.pos, TEXT, "]"
		case ' ':
			return l.pos, SPACE, "_"
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
				// need to pass to a new function to add to attributes
				return sb.String(), TEXTMOD
			} else {
				sb.WriteRune(r)
			}
		case SPACEMOD:
			if r == '[' {
				l.backup()
				return sb.String(), ILLEGALNEST
			} else if r == ']' {
				// need to pass to a new function to add to attributes
				return sb.String(), SPACEMOD
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

func RemoveSpaces(input string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, input)
}

func SplitTextModifier(input string) (string, string, error) {
	// set index to first occurrence of :
	index := strings.Index(input, ":")

	// if : is found, split content into two parts based on : location
	if index != -1 {
		beginning := input[:index]
		beginning = RemoveSpaces(beginning)
		end := input[index+1:]
		return beginning, end, nil
	}

	return input, "", fmt.Errorf("command incomplete: no colon")
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

// need to test mods and stuff
func main() {
	input := "text {mod}\n[[spacing]{mod2}[spacing2] text2 "
	reader := strings.NewReader(input)
	lexer := NewLexer(reader)
	// in the case of some illegal nest or other illegal token, we need to just 
	// break out of the loop with the location of the error
	for {
		pos, tok, lit := lexer.Lex()
		if tok == EOF {
			break
		}

		fmt.Printf("%d:%d | %s | %s | \n", pos.line, pos.column, tok, lit)
	}
}
