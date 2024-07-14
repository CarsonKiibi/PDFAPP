package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/carsonkiibi/pdfapp/backend/process/commands"
)

type Position struct {
	line   int
	column int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, column: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (Position, commands.Token, string) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, commands.Token{Literal: "end"}, ""
			}
			panic(err)
		}

		if r == '{' {
			content, err := l.lexTextMod()
			if err != nil {
				panic(err)
			}
			return l.pos, commands.Token{Literal: string(r)}, content
		}
	}
}

// need to test this
func (l *Lexer) lexTextMod() (string, error) {
	var sb strings.Builder

	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			return "", fmt.Errorf("unexpected EOF")
		}

		switch string(r) {
		case "{", "}":
			return string(r), nil
		}

		sb.WriteRune(r)
	}

	return sb.String(), nil
}

func main() {
	input := "hello {hello} hello"
	reader := strings.NewReader(input)
	lexer := NewLexer(reader)
	for {
		pos, tok, lit := lexer.Lex()
		if tok.Literal == "end" {
			break
		}

		fmt.Printf("%d:%d | \t%s | \t%s | \n", pos.line, pos.column, tok.Literal, lit)
	}
}
