package process

func ProcessInput(sentence string) []Token {
	var tokens []Token
	var word []rune
	inBraces := false

	for i, char := range sentence {
		switch char {
		case '{':
			if inBraces {
				tokens = append(tokens, Token{Type: ErrorIllegalNesting, Literal: "{", Attributes: TokenAttributes{ErrorAt: i}})
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
				tokens = append(tokens, Token{Type: ErrorIllegalNesting, Literal: "}", Attributes: TokenAttributes{ErrorAt: i}})
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
