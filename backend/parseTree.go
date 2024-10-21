package main


import (
    
)

type treeitem int 

/*
ABSTRACT SYNTAX TREE RULES

1. Content

CONTENT

USAGE/TERMINATION
- Content can immediately follow anything that is not expecting some specific
  item (eg TEXT_SIZE expects a number immediately after, not content)


2. Format - Font styling

FORMAT_BOLD
FORMAT_ITALIC
FORMAT_UNDERLINE
FORMAT_NONE
FORMAT_SIZE

USAGE REQUIREMENTS
- Must be the following form:
    FORMAT_<STYLE>*X, CONTENT*X
- Must not contain duplicate styles
- FORMAT_X must not contain any children of any type, it will be the youngest
  child
- FORMAT_SIZE must have a number immediately following it
- In the case where unstyled text immediately follows styled text, use
  FORMAT_NONE before the unstyled text
- FORMAT_NONE does not need to be used for every case in which text is unstyled,
  only in above case

USAGE TERMINATION
- Format does not usually have an explicit termination item, it will terminate 
  when it reaches something that is not text

3. Format - Alignment

FORMAT_ALIGN
LEFT
CENTER
RIGHT
FORMAT_ALIGN_END

USAGE REQUIREMENTS
- Must be in the following form: 
    FORMAT_ALIGN, <DIR*1>, CONTENT*X, FORMAT_ALIGN_END 
- Must be the parent of either styled text or unstyled text
- FORMAT_ALIGN must have one of LEFT, CENTER, RIGHT immediately following it
- Cannot be a child of bullet points (BULLET START)
- Cannot be a child of another align

USAGE TERMINATION
- FORMAT_ALIGN must be terminated with FORMAT_ALIGN_END

3. Line Break

LINE_BREAK



*/

const (
    STRING treeitem = iota 

    CONTENT

    FORMAT_BOLD
    FORMAT_ITALIC
    FORMAT_UNDERLINE
    FORMAT_NONE
    FORMAT_SIZE

    FORMAT_ALIGN
    LEFT
    CENTER
    RIGHT
    FORMAT_ALIGN_END

    LINE_BREAK

    BULLET_START
    BULLET_INDENT_BEFORE
    BULLET_INDENT_AFTER
    BULLET_END

    INDENT
    VERTICAL
    HORIZONTAL

    ENDFILE
    


)

func scanTokens(tokens []Token) {
    
}