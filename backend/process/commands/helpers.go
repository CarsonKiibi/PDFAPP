package commands

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
		tokens = append(tokens, HandleTextModifier(string(word[1:]))...)
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
