package lexer

type State int
type StateTransitions map[LiteralType]State

const (
	S State = iota
	//making words
	Q01
	Q02
	//hex number
	Q03
	//number
	Q04
	//final state: correct word
	Q05
	//fs: hex
	Q06
	//fs number
	Q07
	//number
	Q08
	//ERROR word
	Q09
	//fs number
	Q10
	//ERROR num
	Q11
	//ERROR lexeme
	Q12
	//ERROR, not used
	Q13
)

//Transitions matrix of finite state machine
var StateTable = map[State]StateTransitions{
	S:   {DIGIT: Q04, LETTER: Q01, UNDERSCORE: Q12, COMMA: Q12, LEXEME_END: Q12, ANOTHER: Q12},
	Q01: {DIGIT: Q01, LETTER: Q01, UNDERSCORE: Q01, COMMA: Q09, LEXEME_END: Q05, ANOTHER: Q09},
	Q02: {DIGIT: Q04, LETTER: Q13, UNDERSCORE: Q13, COMMA: Q13, LEXEME_END: Q13, ANOTHER: Q13},
	Q03: {DIGIT: Q03, LETTER: Q11, UNDERSCORE: Q11, COMMA: Q11, LEXEME_END: Q06, ANOTHER: Q11},
	Q04: {DIGIT: Q04, LETTER: Q11, UNDERSCORE: Q11, COMMA: Q08, LEXEME_END: Q07, ANOTHER: Q11},
	Q08: {DIGIT: Q08, LETTER: Q11, UNDERSCORE: Q11, COMMA: Q11, LEXEME_END: Q10, ANOTHER: Q11},
}
