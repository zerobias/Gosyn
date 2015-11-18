package lexer

import (
	"bytes"
	"fmt"
	//"llsp/debug"
	//"llsp/models"
	//"llsp/output"
	//. "llsp/output/colorman"
	"strconv"
	//"strings"
)

var (
	lexicalTokens TokenList
	parsedLexemes LexemeSequence
)

type State int
type StateTransitions map[LiteralType]State

const (
	S   State = iota
	Q01       //making words
	Q02
	Q03 //hex number
	Q04 //number
	Q05 //final state correct word
	Q06 //fs hex
	Q07 //fs number
	Q08 //number
	Q09 //ERROR word
	Q10 //fs number
	Q11 //ERROR num
	Q12 //ERROR lexeme
	Q13 //ERROR, not used
)

var StateTable = map[State]StateTransitions{
	S:   {DIGIT: Q04, LETTER: Q01, UNDERSCORE: Q12, COMMA: Q12, LEXEME_END: Q12, ANOTHER: Q12},
	Q01: {DIGIT: Q01, LETTER: Q01, UNDERSCORE: Q01, COMMA: Q09, LEXEME_END: Q05, ANOTHER: Q09},
	Q02: {DIGIT: Q04, LETTER: Q13, UNDERSCORE: Q13, COMMA: Q13, LEXEME_END: Q13, ANOTHER: Q13},
	Q03: {DIGIT: Q03, LETTER: Q11, UNDERSCORE: Q11, COMMA: Q11, LEXEME_END: Q06, ANOTHER: Q11},
	Q04: {DIGIT: Q04, LETTER: Q11, UNDERSCORE: Q11, COMMA: Q08, LEXEME_END: Q07, ANOTHER: Q11},
	Q08: {DIGIT: Q08, LETTER: Q11, UNDERSCORE: Q11, COMMA: Q11, LEXEME_END: Q10, ANOTHER: Q11},
}

func SetTokens(tokens TokenList) {
	lexicalTokens = tokens
}

func (s State) isFinalState() bool {
	return len(StateTable[s]) == 0
}

func ScanCycle(text *LiteralString) *LexemeSequence {
	state := S
	index := 0
	isScanComplete := func() bool {
		return index >= text.Len() //TODO make sending LEXEME_END(on file end) to SM
	}

	for !isScanComplete() {
		if state.isFinalState() {
			state = S
		}
		var charT LiteralType
		state, charT = ScanStateMachine(&index, text, state)
		charT = charT
		/*debug.Print([]string{"i", "state", "char", "char type"},
		index, debug.GetEnumText(debug.STATE, int(state)), text[index],
		debug.GetEnumText(debug.SYMB, int(charT)))*/
		index++
	}
	//	output.PrintSectionTitle("PARSED", "TOKENS")
	//	output.PrintInLine(0, parsedLexemes)

	/*output.Par = *new(output.Paragraph)
	output.Par.Init()
	output.Par.Send(0, COL_RED, "CATEGORY", "VALUE", "NAME")
	for _, token := range parsedLexemes {
		output.Par.Send(0, RESET,
			debug.GetEnumText(debug.CATEGORY, int(token.LCat)),
			debug.AddQuotes(token.Value), token.Name)
	}
	output.Par.Write()*/

	return &parsedLexemes
}

func ScanStateMachine(actualIndex *int, text *LiteralString, firstState State) (state State, charType LiteralType) {
	state = firstState
	buffer := NewLS()
	index := *actualIndex
	for !state.isFinalState() && index < text.Len() {
		success, delimeter := forwardSearch(index, *text)
		//debug.PrintString(index, *actualIndex)
		if success {
			//debug.PrintString(" ", " ", "delim:"+debug.AddQuotes(delimeter.Value), index, *actualIndex)
			charType = LEXEME_END
			if len(*buffer) > 0 {
				parsedLexemes = append(parsedLexemes, *analyzeBuffer(buffer))
			}
			if delimeter.Name != "space" {
				parsedLexemes = append(parsedLexemes, delimeter)
			}
			buffer = NewLS()
			/*debug.ColorPrint(debug.BOLDRED, index, *actualIndex, delimeter.Value,
			len(delimeter.Value))*/
			index += (*delimeter.Value).Len() - 1
			*actualIndex = index
		} else {
			currentChar := (*text)[index]
			charType = currentChar.Type()
			*buffer = append(*buffer, currentChar)
			//debug.PrintString(" ", " ", debug.GetEnumText(debug.SYMB, int(charType)), index, *actualIndex, debug.AddQuotes(buffer))
		}
		index++
		//debug.ColorPrint(debug.GRAY, " ", " ", " ", " ", index)
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
