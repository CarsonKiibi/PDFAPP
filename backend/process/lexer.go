package process

import (
	"bufio"
	"fmt"
	"io"

	"github.com/carsonkiibi/pdfapp/backend/process/commands"
)

type Position struct {
	line   int
	column int
}

type Lexer struct {
	pos     Position
	nextPos Position
	reader  *bufio.Reader
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
			fmt.Println(r)
		}
	}
}
