package process

func ProcessInput(sentence string) []Token {
	var tokens []Token
	var word []rune
	inBraces := false

	for i, char := range sentence {
		switch char {
		case '{':
			tokens, word, inBraces = HandleOpenBrace(inBraces, word, tokens, i)
		case '}':
			tokens, word, inBraces = HandleCloseBrace(inBraces, word, tokens, i)
		default:
			word = append(word, char)
		}
	}

	if len(word) > 0 {
		tokens = append(tokens, Token{Type: TokenText, Literal: string(word)})
	}

	return tokens
}
