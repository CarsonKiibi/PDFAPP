package process

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
				return l.pos, commands.Token{}, ""
			}
			panic(err)
		}

		if r == '{' {
			content, err := l.lexTextMod()
			if err != nil {
				panic(err)
			}
			return l.pos, commands.Token{}, content
		}
	}
}

// need to test this
func (l *Lexer) lexTextMod() (string, error) {
	var sb strings.Builder
	nestedBrackets := 0

	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			return "", fmt.Errorf("unexpected EOF")
		}

		if r == '{' {
			nestedBrackets++
		} else if r == '}' {
			if nestedBrackets > 0 {
				return "", fmt.Errorf("nested curly brackets are not allowed")
			}
			break
		}

		sb.WriteRune(r)
	}

	return sb.String(), nil
}
