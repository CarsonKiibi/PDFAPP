package process

type TokenType string
type ErrorType string

const (
	// text
	TokenText         TokenType = "TEXT"
	TokenTextModifier TokenType = "TEXT_MOD"

	// format
	TokenHorizontalSpacing TokenType = "HORIZ_SPACING"
	TokenVerticalSpacing   TokenType = "VERT_SPACING"
	TokenIndent            TokenType = "INDENT"
	TokenNewLine           TokenType = "NEW_LINE"

	// bullet
	TokenBulletStart TokenType = "BULLET_START"
	TokenBullet      TokenType = "BULLET"

	// maybe remove this
	TokenWarning TokenType = "WARNING"

	// end of file
	TokenEOF TokenType = "EOF"

	// ILLEGAL
	ErrorIllegalNesting   TokenType = "ILL_NEST"
	TokenIllegalCharacter TokenType = "ILL_CHAR"
	TokenIllegalNoParent  TokenType = "ILL_NO_PARENT"
)

type Token struct {
	Type       TokenType
	Literal    string
	Attributes TokenAttributes
}

type TokenAttributes struct {
	Bold        bool
	Italic      bool
	Underline   bool
	BulletChild bool
	Size        int
	Error       error
	ErrorAt     int
}
