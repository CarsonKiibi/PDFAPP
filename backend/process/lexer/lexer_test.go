package main

// import (
// 	"strings"
// 	"testing"

// 	"github.com/carsonkiibi/pdfapp/backend/process/commands"
// )

// func TestLexer_Lex(t *testing.T) {
// 	input := "Hello {world}!"
// 	reader := strings.NewReader(input)
// 	lexer := NewLexer(reader)

// 	pos, token, content := lexer.Lex()
// 	if content != "world" {
// 		t.Errorf("expected 'world', got '%s'", content)
// 	}

// 	// Checking Position and Token is beyond this simple test
// 	if (pos != Position{line: 1, column: 0}) || (token != commands.Token{}) {
// 		t.Errorf("unexpected Position or Token")
// 	}
// }

// func TestLexer_lexTextMod(t *testing.T) {
// 	cases := []struct {
// 		input       string
// 		expected    string
// 		expectError bool
// 	}{
// 		{"{valid}", "valid", false},
// 		{"{nested {brackets}}", "", true},
// 		{"{unclosed", "", true},
// 	}

// 	for _, tc := range cases {
// 		reader := strings.NewReader(tc.input)
// 		lexer := NewLexer(reader)

// 		_, _, _ = lexer.Lex() // This sets the lexer to the correct position to read the content within `{}`

// 		content, err := lexer.lexTextMod()
// 		if (err != nil) != tc.expectError {
// 			t.Errorf("for input '%s', unexpected error status: %v", tc.input, err)
// 		}
// 		if content != tc.expected {
// 			t.Errorf("for input '%s', expected '%s', got '%s'", tc.input, tc.expected, content)
// 		}
// 	}
// }
