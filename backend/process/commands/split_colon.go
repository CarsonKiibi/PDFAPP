package commands

import (
	"fmt"
	"strings"
	"unicode"
)

// Removes spaces and returns modified string
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

// Splits text across first instance of colon (:) 
// SUCCESS: return content split into beginning and end, nil error
// FAIL: returns input as first string (beginning), "" as end, and error
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
