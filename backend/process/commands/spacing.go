package commands

import "strconv"

// Splits content into direction and size
// SUCCESS: return token with Literal, Type, TokenAttribute{Size} updated
// FAIL: return token with error updated
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
	return []Token{{Attributes: TokenAttributes{Error: err}}}
}