package commands

import (
	"strconv"
)

// ignores the command following the symbol, returns token as just text
func HandleIgnoreCommand(command string) Token {
	return Token{}
}

// fix return tokens, test!
func HandleDocumentSpacing(input string) []Token {
	dir, size, err := SplitTextModifier(input)
	if err != nil {
		return []Token{{Attributes: TokenAttributes{Error: err}}}
	}

	s, err := strconv.Atoi(size)
	if err != nil {
		return []Token{{Type: TokenIllegalCharacter}}
	}
	if len(dir) != 1 {
		return []Token{{Literal: dir, Attributes: TokenAttributes{Size: s}}}
	}
	return []Token{{Literal: dir, Attributes: TokenAttributes{}}}
}

func HandleBulletStart(command string) Token {
	return Token{}
}

func HandleBulletChildren(command string) Token {
	return Token{}
}

func HandleOpenCurlyBrace(inBraces bool, word []rune, tokens []Token, i int) ([]Token, []rune, bool) {

	if len(word) > 0 {
		tokens = append(tokens, Token{Type: TokenText, Literal: string(word)})
		word = []rune{}
	}
	inBraces = true
	word = append(word, '{')
	return tokens, word, inBraces
}

func HandleClosedCurlyBrace(inBraces bool, word []rune, tokens []Token, i int) ([]Token, []rune, bool) {
	if inBraces {
		tokens = append(tokens, HandleTextModifier(string(word[1:]))...)
		word = []rune{}
		inBraces = false
	} else {
		tokens = append(tokens, Token{Type: ErrorIllegalNesting, Literal: "}", Attributes: TokenAttributes{ErrorAt: i}})
	}
	return tokens, word, inBraces
}

func HandleOpenSquareBrace(inBraces bool, word []rune, tokens []Token, i int) ([]Token, []rune, bool) {
	if len(word) > 0 {
		tokens = append(tokens, Token{Type: TokenText, Literal: string(word)})
		word = []rune{}
	}
	inBraces = true
	word = append(word, '{')
	return tokens, word, inBraces
}

func HandleClosedSquareBrace(inBraces bool, word []rune, tokens []Token, i int) ([]Token, []rune, bool) {
	if inBraces {
		tokens = append(tokens, HandleDocumentSpacing(string(word[1:]))...)
		word = []rune{}
		inBraces = false
	} else {
		tokens = append(tokens, Token{Type: ErrorIllegalNesting, Literal: "}", Attributes: TokenAttributes{ErrorAt: i}})
	}
	return tokens, word, inBraces
}

func IsNestedError(inCBraces bool, inSBraces bool, char string) bool {
	return (inCBraces || inSBraces) && (string(char) == "{" || string(char) == "[")
}
