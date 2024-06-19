package process

import (
	"fmt"

	"github.com/carsonkiibi/pdfapp/backend/process/commands"
)

func ProcessInput(sentence string) []commands.Token {
	var tokens []commands.Token
	var word []rune
	inCBraces := false
	inSBraces := false

	for i, char := range sentence {
		if commands.IsNestedError(inCBraces, inSBraces, string(char)) {
			tokens = append(
				tokens,
				commands.Token{
					Type:    commands.ErrorIllegalNesting,
					Literal: string(char),
					Attributes: commands.TokenAttributes{
						ErrorAt: i,
						Error:   fmt.Errorf("illegal nesting"),
					},
				},
			)

			break
		}
		switch char {
		case '{':
			// need to handle nested error here
			tokens, word, inCBraces = commands.HandleOpenCurlyBrace(inCBraces, word, tokens, i)
		case '}':
			tokens, word, inCBraces = commands.HandleClosedCurlyBrace(inCBraces, word, tokens, i)
		case '[':
			tokens, word, inSBraces = commands.HandleOpenSquareBrace(inSBraces, word, tokens, i)
		case ']':
			tokens, word, inSBraces = commands.HandleClosedSquareBrace(inSBraces, word, tokens, i)
		default:
			word = append(word, char)
		}
	}

	if len(word) > 0 {
		tokens = append(tokens,
			commands.Token{
				Type:    commands.TokenText,
				Literal: string(word)})
	}

	return tokens
}
