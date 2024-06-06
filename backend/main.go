package main

import (
	"fmt"
	"time"
	"unicode"
)

type TokenType string

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

	// illegal
	TokenIllegalNesting   TokenType = "ILL_NEST"
	TokenIllegalCharacter TokenType = "ILL_CHAR"
	TokenIllegalNoParent  TokenType = "ILL_NO_PARENT"
)

type Token struct {
	Type       TokenType
	Literal    string
	Attributes TokenAttributes
}

type TokenAttributes struct {
	Bold      bool
	Italic    bool
	Underline bool
	Size      int
}

// ignores the command following the symbol, returns token as just text
func HandleIgnoreCommand(command string) (Token, error) {
	return Token{}, nil
}

// create token bold, italic, underline, size attributes to token and return
func HandleTextModifier(command string) (Token, error) {
	return Token{}, nil
}

func HandleBulletStart(command string) (Token, error) {
	return Token{}, nil
}

func HandleBulletChildren(command string) (Token, error) {
	return Token{}, nil
}

// split sentence into tokens
func ProcessInput(sentence string) ([]string, error) {
	var words []string
	var word []rune
	inBraces := false

	for _, char := range sentence {
		if char == '{' {
			if inBraces == true {
				return nil, fmt.Errorf("Error: nested command (curly) braces")
			}
			if len(word) > 0 && !inBraces {
				words = append(words, string(word))
				word = []rune{}
			}
			inBraces = true
			word = append(word, char)
		} else if char == '}' {
			word = append(word, char)
			words = append(words, string(word))
			word = []rune{}
			fmt.Printf("Returning %s", string(word))
			inBraces = false
		} else if unicode.IsSpace(char) && !inBraces {
			if len(word) > 0 {
				words = append(words, string(word))
				word = []rune{}
			}
		} else {
			word = append(word, char)
		}
	}

	if len(word) > 0 {
		words = append(words, string(word))
	}

	return words, nil
}

func main() {
	sentence := "My name is {Bob} and I like {Shiny Green Apples}"
	start := time.Now()
	words, err := ProcessInput(sentence)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(words) // Output: [My name is {Bob} and I like {Shiny Green Apples}]
	duration := time.Since(start)
	fmt.Println(duration.Microseconds())
}
