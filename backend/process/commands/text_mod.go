package commands

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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
