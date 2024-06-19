package commands

// todo
func HandleDocumentSpacing(input string) []Token {
	dir, size, err := SplitTextModifier(input)
	if err != nil {
		return []Token{{Attributes: TokenAttributes{Error: err}}}
	}

	if len(dir) != 1 {
		return []Token{{Literal: size, Attributes: TokenAttributes{}}}
	}
	return []Token{{Literal: size, Attributes: TokenAttributes{}}}
}
