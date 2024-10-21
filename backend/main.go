package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"time"
	"unicode"
	"fmt"
)

type TokenType int

const (
	EOF TokenType = iota
	LCURLYBRACE
	RCURLYBRACE
	LROUNDBRACE
	RROUNDBRACE
	LSQUAREBRACE
	RSQUAREBRACE
	BSLASH
	FSLASH
	LESSTHAN
	GREATERTHAN
	COLON
	ADD
	SUB
	MUL
	SPACE
	TEXT
	NUMBER
	ILLEGAL
)

var tokenTypes = []string{
	EOF:         "EOF",
	LCURLYBRACE: "{",
	RCURLYBRACE: "}",
	LROUNDBRACE: "(",
	RROUNDBRACE: ")",
	LSQUAREBRACE: "[",
	RSQUAREBRACE: "]",
	BSLASH: "\\",
	FSLASH: "/",
	LESSTHAN: "<",
	GREATERTHAN: ">",
	COLON: ":",
	ADD: "+",
	SUB: "-",
	MUL: "*",
	SPACE: " ",
	TEXT: "TEXT",
	NUMBER: "NUMBER",
	ILLEGAL: "ILLEGAL",
}

type Position struct {
	line   int
	column int
}

// new token implementation


type Token interface {
	Type() TokenType
}

type BaseToken struct {
	tokenType TokenType
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func (b BaseToken) Type() TokenType {
	return b.tokenType
}

type TextToken struct {
	BaseToken
	text string 
}

func (t TextToken) Text() string {
	return t.text
}

type NumberToken struct {
	BaseToken
	number float64
}

func (n NumberToken) Number() float64 {
	return n.number
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
				return l.pos, BaseToken{tokenType: EOF}
			}

			panic(err)
		}

		l.pos.column++

		switch r {
		case '\n':
			l.resetPosition()
		case '{':
			return l.pos, BaseToken{tokenType: LCURLYBRACE}
		case '}':
			return l.pos, BaseToken{tokenType: RCURLYBRACE}
		case '(':
			return l.pos, BaseToken{tokenType: LROUNDBRACE}
		case ')':
			return l.pos, BaseToken{tokenType: RROUNDBRACE}
		case '[':
			return l.pos, BaseToken{tokenType: LSQUAREBRACE}
		case ']':
			return l.pos, BaseToken{tokenType: RSQUAREBRACE}
		case '\\':
			return l.pos, BaseToken{tokenType: BSLASH}
		case '/':
			return l.pos, BaseToken{tokenType: FSLASH}
		case '<':
			return l.pos, BaseToken{tokenType: LESSTHAN}
		case '>':
			return l.pos, BaseToken{tokenType: GREATERTHAN}
		case ':':
			return l.pos, BaseToken{tokenType: COLON}
		case '+':
			return l.pos, BaseToken{tokenType: ADD}
		case '-':
			return l.pos, BaseToken{tokenType: SUB}
		case '*':
			return l.pos, BaseToken{tokenType: MUL}
		case ' ':
			return l.pos, BaseToken{tokenType: SPACE}
		default:
			l.backup()
			lit := l.lexChar()
			valid, num := isNumeric(lit)
			if valid {
				return l.pos, NumberToken{BaseToken: BaseToken{tokenType: NUMBER}, number: num}
			} else {
				return l.pos, TextToken{BaseToken: BaseToken{tokenType: TEXT}, text: lit}
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

func (l *Lexer) lexChar() string {
	var lit string 
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.column++
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

// Helper function to check if a string contains only numbers or a single decimal point
func isNumeric(s string) (bool, float64) {
	hasDecimal := false

	// Iterate through the string to check each character
	for _, r := range s {
		if r == '.' {
			// If there's more than one decimal, return false
			if hasDecimal {
				return false, 0
			}
			hasDecimal = true
		} else if !unicode.IsDigit(r) {
			// If the character is not a digit, return false
			return false, 0
		}
	}

	// Try to convert the string to a float
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false, 0
	}

	return true, num
}

func main() {
	input := "{}asdasd[ {hihi} hihi: 12sd"
	reader := strings.NewReader(input)
	lexer := NewLexer(reader)
	
	start := time.Now()
	var tokens []Token

	for {
		_, tok := lexer.Lex()
		if tok.Type() == EOF {
			break
		}
		tokens = append(tokens, tok)
		fmt.Printf("%v", tok.Type())
		if tok.Type() == TEXT {
			textToken, ok := tok.(TextToken)
			if ok {
				fmt.Printf(" - %s\n", textToken.Text()) // Print the text of the token
			} else {
				fmt.Println() // fallback just in case something goes wrong
			}
		}

		if tok.Type() == NUMBER {
			numberToken, ok := tok.(NumberToken)
			if ok {
				fmt.Printf(" - %f\n", numberToken.Number()) // Print the text of the token
			} else {
				fmt.Println() // fallback just in case something goes wrong
			}
		}
	}

	duration := time.Since(start)
	fmt.Println("")
	fmt.Println(duration)
}
