package lexer

import (
	"strings"
)

type LiteralType int

type Literal struct {
	rune
	Index int
}

type LiteralString []Literal
type PLString *LiteralString

func (ls *LiteralString) AsRune() (result []rune) { //TODO remove "As"
	for _, e := range []Literal(*ls) {
		result = append(result, e.rune)
	}
	return
}

func (ls *LiteralString) Add(l ...Literal) {
	for _, e := range l {
		*ls = append([]Literal(*ls), e)
	}
}

func (ls *LiteralString) AddNew(r rune, n int) {
	*ls = append([]Literal(*ls), Literal{r, n})
}

func (ls LiteralString) String() string {
	result := make([]rune, 0, len([]Literal(ls)))
	for _, l := range []Literal(ls) {
		result = append(result, l.rune)
	}
	return string(result)
}

func (ls LiteralString) Len() int {
	return len([]Literal(ls))
}

func NewLiteralString() *LiteralString {
	slice := make([]Literal, 0)
	newLs := *new(LiteralString)
	newLs = slice
	return &newLs
}

func NewLStringText(str string) *LiteralString { //TODO remove, possibly not used
	result := NewLiteralString()
	for i, val := range []rune(strings.ToLower(str)) {
		(*result).AddNew(val, i)
	}
	return result
}

const (
	DIGIT LiteralType = iota
	LETTER
	UNDERSCORE
	COMMA
	LEXEME_END //TODO remove?
	ANOTHER
)

var charGroups = map[LiteralType]string{
	DIGIT:      "0123456789",
	LETTER:     "abcdefghijklmnopqrstuvwxyz",
	UNDERSCORE: "_",
	COMMA:      ",",
}

func (l Literal) Type() LiteralType {
	for currentType, includedChars := range charGroups {
		if strings.ContainsRune(includedChars, l.rune) {
			return currentType
		}
	}
	return ANOTHER
}
