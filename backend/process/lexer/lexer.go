package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type TokenType int

const (
	EOF TokenType = iota
	ILLEGALNEST
	ILLEGAL
	SPACE
	SPACEMOD
	TEXTMOD
	TEXT
)

var tokenTypes = []string{
	EOF:         "EOF",
	ILLEGALNEST: "ILLEGALNEST",
	ILLEGAL:     "ILLEGAL",
	SPACE:       "SPACE",
	SPACEMOD:    "SPACEMOD",
	TEXTMOD:     "TEXTMOD",
	TEXT:        "TEXT",
}

type Position struct {
	line   int
	column int
}

type Token struct {
	Type       TokenType
	Literal    string
	Attributes TokenAttributes
}

type TokenAttributes struct {
	Bold        bool
	Italic      bool
	Underline   bool
	BulletChild bool
	Size        int
	Error       error
	ErrorAt     int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func (t TokenType) String() string {
	return tokenTypes[t]
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, column: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (Position, Token) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, Token{Type: EOF, Literal: ""}
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
			token := l.lexText(TEXTMOD)
			if token.Type == EOF {
				return startPos, token
			}
			return startPos, token
		case '}':
			return l.pos, Token{Type: TEXT, Literal: "}"}
		case '[':
			startPos := l.pos
			token := l.lexText(SPACEMOD)
			if token.Type == EOF {
				return startPos, token
			}
			return startPos, token
		case ']':
			return l.pos, Token{Type: TEXT, Literal: "]"}
		case ' ':
			return l.pos, Token{Type: SPACE, Literal: "_"}
		default:
			startPos := l.pos
			l.backup()
			token := l.lexText(TEXT)
			if token.Type == EOF {
				return startPos, token
			}
			return startPos, token
		}
	}
}

func (l *Lexer) lexText(tokenType TokenType) Token {
	var sb strings.Builder
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// If we've accumulated any text, return it before EOF
				if sb.Len() > 0 {
					return Token{Type: tokenType, Literal: sb.String()}
				}
				return Token{Type: EOF, Literal: sb.String()}
			}
			panic(err)
		}
		l.pos.column++

		switch tokenType {
		case TEXTMOD:
			if r == '{' {
				l.backup()
				return Token{Type: ILLEGALNEST, Literal: sb.String()}
			} else if r == '}' {
				token := HandleTextMod(sb.String())

				return token
			} else {
				sb.WriteRune(r)
			}
		case SPACEMOD:
			if r == '[' {
				l.backup()
				return Token{Type: ILLEGALNEST, Literal: sb.String()}
			} else if r == ']' {
				return Token{Type: SPACEMOD, Literal: sb.String()}
			} else {
				sb.WriteRune(r)
			}
		case TEXT:
			if r == ' ' || r == '{' || r == '}' || r == '[' || r == ']' {
				l.backup()
				return Token{Type: TEXT, Literal: sb.String()}
			} else {
				sb.WriteRune(r)
			}
		}
	}
}

func HandleTextMod(command string) Token {
	var token Token

	mods, content, err := SplitTextMod(command)
	if err != nil {
		token.Attributes.Error = err
		return token
	}
	modsSplit := strings.Split(mods, ",")
	re := regexp.MustCompile(`^([A-Z])(\d+)$`)

	for _, part := range modsSplit {
		switch {
		case part == "B":
			token.Attributes.Bold = true
		case part == "I":
			token.Attributes.Italic = true
		case part == "U":
			token.Attributes.Underline = true
		case re.MatchString(part):
			matches := re.FindStringSubmatch(part)
			if len(matches) == 3 {
				letter := matches[1]
				number, err := strconv.Atoi(matches[2])
				if err != nil {
					token.Attributes.Error = fmt.Errorf("text mod details cannot compile")
				} else {
					if letter == "S" && number > 0 {
						token.Attributes.Size = number
					}
				}
			}
		default:
			token.Attributes.Error = fmt.Errorf("error ")
		}
	}

	token.Type = TEXTMOD
	token.Literal = content

	return token
}

func SplitTextMod(input string) (string, string, error) {
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

func (l *Lexer) ignoreNext() {
	// ??
}

func main() {
	input := "text {S:mod} [spacing][spacing2] text2 \n hi"
	reader := strings.NewReader(input)
	lexer := NewLexer(reader)
	for {
		pos, tok := lexer.Lex()
		if tok.Type == EOF {
			break
		} else if tok.Attributes.Error != nil {
			fmt.Printf("Error: %d \n", tok.Attributes.Error)
		} else {
			fmt.Printf("%d:%d | %s | %s | \n", pos.line, pos.column, tok.Type, tok.Literal)
		}
	}
}
