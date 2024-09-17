package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/jung-kurt/gofpdf"
)

// need size global

// MAKE THIS AN INTERFACE!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! goober

type TokenType int

const (
	EOF TokenType = iota
	ILLEGALNEST
	ILLEGAL
	SPACE
	SPACEMOD
	TEXTMOD
	TEXT
	TEXTMODFAILED
	SPACEMODFAILED
)

var tokenTypes = []string{
	EOF:            "EOF",
	ILLEGALNEST:    "ILLEGALNEST",
	ILLEGAL:        "ILLEGAL",
	SPACE:          "SPACE",
	SPACEMOD:       "SPACEMOD",
	TEXTMOD:        "TEXTMOD",
	TEXT:           "TEXT",
	TEXTMODFAILED:  "TEXTMODFAILED",
	SPACEMODFAILED: "SPACEMODFAILED",
}

type Position struct {
	line   int
	column int
}

// new token implementation

type Token interface {
	Type() TokenType
	Literal() string
}

type BaseToken struct {
	tokenType TokenType
	literal   string
}

func (b BaseToken) Type() TokenType {
	return b.tokenType
}

func (b BaseToken) Literal() string {
	return b.literal
}

type TextToken struct {
	BaseToken
	Bold      bool
	Italic    bool
	Underline bool
	Center    bool
}

type SpacingToken struct {
	BaseToken
	Size       int
	Horizontal bool
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
				fmt.Println("returning EOF token in lexText")
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
				token := HandleSpaceMod(sb.String())
				return token
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
		case part == "C":
			token.Attributes.Center = true
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
			token.Attributes.Error = fmt.Errorf(TEXTMODFAILED.String())
		}
	}

	token.Type = TEXTMOD
	token.Literal = content

	return token
}

func HandleSpaceMod(command string) Token {
	var token Token
	command = RemoveSpaces(command)

	mod, content, err := SplitTextMod(command)
	if err != nil {
		token.Attributes.Error = err
		return token
	}
	// need to check if content is a number
	// prob need to switch
	size, err := strconv.Atoi(content)
	if err != nil {
		token.Attributes.Error = fmt.Errorf(ILLEGAL.String())
	}
	token.Attributes.Size = size
	switch mod {
	case "H":
		token.Attributes.Horizontal = true
	case "V":
		token.Attributes.Horizontal = false
	default:
		token.Attributes.Error = fmt.Errorf(ILLEGAL.String())
	}
	token.Type = SPACEMOD
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

// TODO
// return unmodified text from first instance of command opening to last
// func (l *Lexer) ignoreNext() Token {
// 	var sb strings.Builder

// 	for {
// 		r, _, err := l.reader.ReadRune()

// 	}
// }

// put character/word/etc in right location based on current line height
// need to make sure the current line has some max text height, but then we need to know the largest character which might be at the end of the line
// OMG!!!!

// curr is current position
// pageWidth is width exluding margins
func findLineStartPos(currX float64, currY float64, pageWidth float64, pageHeight float64) (float64, float64) {
	

	return 0.0, 0.0
}

func GeneratePDF(tokens []Token) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	defaultTextSize := 11.0
	pdf.SetFont("Arial", "", defaultTextSize)

	marginLeft, marginTop, marginRight, marginBottom := pdf.GetMargins()
	pageWidth, pageHeight := pdf.GetPageSize()
	lineHeight := defaultTextSize * 0.5
	maxWidth := pageWidth - marginLeft - marginRight
	maxHeight := pageHeight - marginTop - marginBottom

	var currentX float64 = marginLeft
	var currentY float64 = marginTop
	var setStyle string = ""

	// Function to get the current font size
	getCurrentFontSize := func() float64 {
		_, fontSize := pdf.GetFontSize()
		return fontSize
	}

	// Function to adjust Y position based on font size
	adjustY := func() {
		currentY += getCurrentFontSize() * 0.3 // Approximate adjustment for ascent
	}

	for _, token := range tokens {
		switch token.Type {
		case TEXT, TEXTMOD:
			if token.Type == TEXTMOD {
				setStyle = ""
				if token.Attributes.Bold {
					setStyle += "B"
				}
				if token.Attributes.Italic {
					setStyle += "I"
				}
				if token.Attributes.Underline {
					setStyle += "U"
				}
				if token.Attributes.Size > 0 {
					pdf.SetFontSize(float64(token.Attributes.Size))
				}
				pdf.SetFontStyle(setStyle)
			}

			words := strings.Split(token.Literal, " ")
			for _, word := range words {
				wordWidth := pdf.GetStringWidth(word)
				if currentX+wordWidth > maxWidth {
					currentX = marginLeft
					currentY += lineHeight + getCurrentFontSize()
					if currentY > maxHeight {
						pdf.AddPage()
						currentY = marginTop
					}
				}
				adjustY()
				pdf.Text(currentX, currentY, word)
				currentY -= getCurrentFontSize() * 0.3 // Reset Y position
				currentX += wordWidth + pdf.GetStringWidth(" ")
			}

		case SPACE:
			spaceWidth := pdf.GetStringWidth(" ")
			if currentX+spaceWidth > maxWidth {
				currentX = marginLeft
				currentY += lineHeight + getCurrentFontSize()
				if currentY > maxHeight {
					pdf.AddPage()
					currentY = marginTop
				}
			} else {
				currentX += spaceWidth
			}

		default:
			fmt.Println("token not recognized")
		}

		pdf.SetFontSize(defaultTextSize)
		pdf.SetFontStyle("")
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {
	input := "{B,S24:HI!} hihihihi"
	reader := strings.NewReader(input)
	lexer := NewLexer(reader)

	start := time.Now()
	var tokens []Token

	for {
		pos, tok := lexer.Lex()
		if tok.Type == EOF {
			break
		}
		tokens = append(tokens, tok)
		fmt.Println(pos)
	}
	pdfBytes, err := GeneratePDF(tokens)
	if err != nil {
		fmt.Println("FAILED")
	} else {
		// Convert PDF bytes to base64
		base64PDF := base64.StdEncoding.EncodeToString(pdfBytes)
		fmt.Println(base64PDF)
	}
	duration := time.Since(start)
	fmt.Printf("Lexing took %v\n", duration)
}
