package lexer

import (
	"fmt"
	"strings"
)

func HandleSource(recieved string) *LiteralString {
	return removeSpecialChars(strings.ToLower(recieved))
}

func removeSpecialChars(recieved string) *LiteralString { //TODO remove result sending
	trim := strings.TrimSpace(recieved)
	result := *NewLiteralString()
	for i, ch := range []rune(trim) {
		r := transformSpecSymbols(ch)
		if checkRune(r, result.AsRune()) {
			result.AddNew(r, i)
		}
	}
	fmt.Println("end removing")
	return &result
}

func checkRune(ch rune, result []rune) bool { //TODO remove result recieving
	if ch == 13 {
		return false
	}
	if len(result) > 0 {
		if ch == ' ' && result[len(result)-1] == ' ' {
			return false
		}
	}
	return true
}

const (
	rune_SPACE rune = 32
	rune_TAB        = 9
	rune_LINE       = 10
)

func transformSpecSymbols(ch rune) rune {
	if ch == rune_SPACE || ch == rune_TAB || ch == rune_LINE {
		return ' '
	}
	return ch
}
