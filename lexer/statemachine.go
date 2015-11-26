package lexer

import (
	"bytes"
	"fmt"
	//"gosyn/debug"
	//"gosyn/models"
	//"gosyn/output"
	//. "gosyn/output/colorman"
	"strconv"
	//"strings"
)

var (
	lexicalTokens TokenList
	parsedLexemes LexemeSequence
)

func SetTokens(tokens TokenList) {
	lexicalTokens = tokens
}

//Zero length of transitions matrix appears only in final state
func (s State) isFinalState() bool {
	return len(StateTable[s]) == 0
}

func ScanCycle(text *LiteralString) *LexemeSequence {
	var state State
	resetState := func() { state = S }
	resetState()
	index := 0
	isScanComplete := func() bool {
		return index >= text.Len() //TODO make sending LEXEME_END(on file end) to SM
	}

	for !isScanComplete() {
		if state.isFinalState() {
			resetState()
		}
		var charT LiteralType
		state, charT = ScanStateMachine(&index, text, state)
		charT = charT
		index++
	}
	return &parsedLexemes
}

func ScanStateMachine(actualIndex *int, text *LiteralString, firstState State) (state State, charType LiteralType) {
	state = firstState
	buffer := NewLiteralString()
	index := *actualIndex
	for !state.isFinalState() && index < text.Len() {
		success, delimeter := forwardSearch(index, *text)
		if success {
			charType = LEXEME_END
			if len(*buffer) > 0 {
				parsedLexemes = append(parsedLexemes, *analyzeBuffer(buffer))
			}
			if delimeter.Name != "space" {
				parsedLexemes = append(parsedLexemes, delimeter)
			}
			buffer = NewLiteralString()
			index += (*delimeter.Value).Len() - 1
			*actualIndex = index
		} else {
			currentChar := (*text)[index]
			charType = currentChar.Type()
			*buffer = append(*buffer, currentChar)
		}
		index++
	}
	return
}

func analyzeBuffer(buf *LiteralString) *Lexeme {
	for _, token := range lexicalTokens.SelectClass(KEYWORD) {
		if compare(*buf, token.Value) {
			return token.deriveLexeme(buf)
		}
	}
	_, err := strconv.Atoi(buf.String())
	fmt.Println("New lexeme ", (*(buf)).String())
	if err == nil {
		return NewLexeme("number", NUMBER, buf)
	} else {
		return NewLexeme("text", TEXT, buf)
	}
}

//search delimeters
func forwardSearch(index int, text LiteralString) (success bool, result Lexeme) {
	checkRange := func(t Token) bool {
		return index+len(t.Value) <= text.Len()
	}
	getTokenTextPart := func(t Token) LiteralString {
		return text[index : index+len(t.Value)]
	}

	success = false
	for _, token := range lexicalTokens.SelectClass(DELIMETER, SIGN) {
		tokenText := getTokenTextPart(token)
		if checkRange(token) && compare(tokenText, token.Value) {
			success = true
			if result.Value == nil || len(token.Value) > (*(result.Value)).Len() { //prefer longest finded token
				result = *token.deriveLexeme(&tokenText)
			}
		}
	}
	return success, result
}

func compare(c1 LiteralString, c2 string) bool {
	return bytes.Equal([]byte(c1.String()), []byte(c2))
}
