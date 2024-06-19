package commands

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// create token bold, italic, underline, size attributes to token and return
// receives text modifier in form mods:content
func HandleTextModifier(command string) []Token {

	var token Token

	// split content into either mods or content by colon
	// m:stuff -> mods = m, content = stuff
	mods, content, err := SplitTextModifier(command)
	if err != nil {
		token.Attributes.Error = err
		return []Token{token}
	}
	modsSplit := strings.Split(mods, ",") // split modifications

	// regular expression for capital letter followed by number
	re := regexp.MustCompile(`^([A-Z])(\d+)$`)

	for _, part := range modsSplit {
		switch {
		case part == "B":
			token.Attributes.Bold = true
		case part == "I":
			token.Attributes.Italic = true
		case part == "U":
			token.Attributes.Underline = true
		// clean this up
		case re.MatchString(part):
			matches := re.FindStringSubmatch(part)
			if len(matches) == 3 {
				letter := matches[1]
				number, err := strconv.Atoi(matches[2])
				if err != nil {
					token.Attributes.Error = fmt.Errorf("Something wrong with text mod specifications")
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
	// if loop exits w no error then give token attributes
	token.Type = TokenTextModifier
	token.Literal = content

	return []Token{token}
}
