package process

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

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
	// set index to first occurrence of :
	index := strings.Index(input, ":")

	// if : is found, split content into two parts based on : location
	if index != -1 {
		beginning := input[:index]
		beginning = RemoveSpaces(beginning)
		end := input[index+1:]
		return beginning, end, nil
	}

	return input, "", fmt.Errorf("command incomplete: no colon")
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
						token.Attributes.Size = number
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

func HandleOpenBrace(inBraces bool, word []rune, tokens []Token, i int) ([]Token, []rune, bool) {
	if inBraces {
		tokens = append(tokens, Token{Type: ErrorIllegalNesting, Literal: "{", Attributes: TokenAttributes{ErrorAt: i, Error: fmt.Errorf("illegal nesting")}})
	} else {
		if len(word) > 0 {
			tokens = append(tokens, Token{Type: TokenText, Literal: string(word)})
			word = []rune{}
		}
		inBraces = true
		word = append(word, '{')
	}
	return tokens, word, inBraces
}

func HandleCloseBrace(inBraces bool, word []rune, tokens []Token, i int) ([]Token, []rune, bool) {
	if inBraces {
		tokens = append(tokens, HandleTextModifier(string(word[1:]))...)
		word = []rune{}
		inBraces = false
	} else {
		tokens = append(tokens, Token{Type: ErrorIllegalNesting, Literal: "}", Attributes: TokenAttributes{ErrorAt: i}})
	}
	return tokens, word, inBraces
}
