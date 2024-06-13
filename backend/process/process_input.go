package process

import "fmt"

func ProcessInput(sentence string) []Token {
	var tokens []Token
	var word []rune
	inCBraces := false
	inSBraces := false

	for i, char := range sentence {
		if IsNestedError(inCBraces, inSBraces, string(char)) {
			tokens = append(tokens, Token{Type: ErrorIllegalNesting, Literal: "baba", Attributes: TokenAttributes{ErrorAt: i, Error: fmt.Errorf("illegal nesting")}})
		}
		switch char {
		case '{':
			// need to handle nested error here
			tokens, word, inCBraces = HandleOpenCurlyBrace(inCBraces, word, tokens, i)
		case '}':
			tokens, word, inCBraces = HandleClosedCurlyBrace(inCBraces, word, tokens, i)
		case '[':
			tokens, word, inSBraces = HandleOpenSquareBrace(inSBraces, word, tokens, i)
		case ']':
			tokens, word, inSBraces = HandleClosedSquareBrace(inSBraces, word, tokens, i)
		default:
			word = append(word, char)
		}
	}

	if len(word) > 0 {
		tokens = append(tokens, Token{Type: TokenText, Literal: string(word)})
	}

	return tokens
}
