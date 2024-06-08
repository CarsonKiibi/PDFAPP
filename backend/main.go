package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type TokenType string
type ErrorType string

const (
	// text
	TokenText         TokenType = "TEXT"
	TokenTextModifier TokenType = "TEXT_MOD"

	// format
	TokenHorizontalSpacing TokenType = "HORIZ_SPACING"
	TokenVerticalSpacing   TokenType = "VERT_SPACING"
	TokenIndent            TokenType = "INDENT"
	TokenNewLine           TokenType = "NEW_LINE"

	// bullet
	TokenBulletStart TokenType = "BULLET_START"
	TokenBullet      TokenType = "BULLET"

	// maybe remove this
	TokenWarning TokenType = "WARNING"

	// end of file
	TokenEOF TokenType = "EOF"

	// ILLEGAL
	ErrorIllegalNesting   ErrorType = "ILL_NEST"
	TokenIllegalCharacter ErrorType = "ILL_CHAR"
	TokenIllegalNoParent  ErrorType = "ILL_NO_PARENT"
)

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

// ignores the command following the symbol, returns token as just text
func HandleIgnoreCommand(command string) Token {
	return Token{}
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
	// set inddex to first occurance of :
	index := strings.Index(input, ":")

	// if : is found, split content into two parts based on : location
	if index != -1 {
		beginning := input[:index]
		beginning = RemoveSpaces(beginning)
		end := input[index+1:]
		return beginning, end, nil
	}

	return input, "", fmt.Errorf("command incomplete: no colon") // and not ignored?
}

// create token bold, italic, underline, size attributes to token and return
func HandleTextModifier(command string) []Token {
	mods, content, err := SplitTextModifier(command)
	if err != nil {
		return []Token{{Attributes: TokenAttributes{Error: err}}}
	}
	modsSplit := strings.Split(mods, ",")

	// regular expression for capital letter followed by number
	re := regexp.MustCompile(`^([A-Z])(\d+)$`)

	var token Token
	token.Type = TokenTextModifier
	token.Literal = content

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
					fmt.Printf("Error converting number: %v\n", err)
				} else {
					if letter == "S" && number > 0 {
						token.Attributes.Size = number // need to modify conditionals to handle case where number is 0 or negative
					}

				}
			}
		default:
			token.Attributes.Error = fmt.Errorf("error reading text modifier")
		}
	}

	return []Token{token}
}

func HandleBulletStart(command string) Token {
	return Token{}
}

func HandleBulletChildren(command string) Token {
	return Token{}
}

// split sentence into tokens

func ProcessInput(sentence string) []Token {
	var tokens []Token
	var word []rune
	inBraces := false

	for i, char := range sentence {
		switch char {
		case '{':
			if inBraces {
				tokens = append(tokens, Token{Literal: "{", Attributes: TokenAttributes{ErrorAt: i}})
			} else {
				if len(word) > 0 {
					tokens = append(tokens, Token{Type: TokenText, Literal: string(word)})
					word = []rune{}
				}
				inBraces = true
				word = append(word, char)
			}
		case '}':
			if inBraces {
				tokens = append(tokens, HandleTextModifier(string(word[1:]))...)
				word = []rune{}
				inBraces = false
			} else {
				tokens = append(tokens, Token{Literal: "}", Attributes: TokenAttributes{ErrorAt: i}})
			}
		default:
			word = append(word, char)
		}
	}

	if len(word) > 0 {
		tokens = append(tokens, Token{Type: TokenText, Literal: string(word)})
	}

	return tokens
}

func main() {
	start := time.Now()
	str := "I like {B,S14:apples} on sandwiches with {U:Tea}"
	out := ProcessInput(str)
	fmt.Println(out)
	duration := time.Since(start)
	fmt.Println(duration.Microseconds())
}
