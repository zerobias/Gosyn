package lexer

import (
	//"fmt"
	. "gosyn/testingtools"
	"reflect"
	"testing"
)

var (
	litStr1 LiteralString = LiteralString{Literal{'a', 0}, Literal{'b', 1}, Literal{' ', 2}, Literal{'c', 3}}
)

func TestNewLiteralString(t *testing.T) {
	var tests = []struct {
		s string
		t LiteralString
	}{
		{"abc", LiteralString{Literal{'a', 0}, Literal{'b', 1}, Literal{'c', 2}}},
		{"AB c", litStr1},
	}
	for _, val := range tests {
		result := *NewLiteralString(val.s)
		if !reflect.DeepEqual(result, val.t) {
			PrintExpected(t, val.t.String(), result.String())
		}
	}
}

func TestAddNew(t *testing.T) {
	comp := func(l Literal, r rune, i int) bool { return l.rune == r && l.Index == i }
	ls := NewLS()

	ls.Add(Literal{'a', 3})
	result := ([]Literal(*ls))[0]
	if comp(result, 'a', 3) == false {
		PrintExpected(t, Literal{'a', 3}, result, false)
	}

	newLit := Literal{'d', 5}
	ls.Add(Literal{'c', 1}, newLit)
	result = ([]Literal(*ls))[2]
	if comp(result, 'd', 5) == false {
		PrintExpected(t, Literal{'d', 5}, result, newLit)
	}

	ls.AddNew('u', 10)
	result = ([]Literal(*ls))[3]
	if comp(result, 'u', 10) == false {
		PrintExpected(t, Literal{'u', 10}, result, 'u', 10)
	}
}

func TestLiteralType(t *testing.T) {
	var tests = []struct {
		l Literal
		t LiteralType
	}{
		{Literal{'a', 0}, LETTER},
		{Literal{'B', 0}, ANOTHER}, //means that uppercase letters can not go here
		{Literal{'\t', 0}, ANOTHER},
		{Literal{',', 0}, COMMA},
		{Literal{'_', 0}, UNDERSCORE},
		{Literal{'3', 0}, DIGIT},
	}
	for _, val := range tests {
		if val.l.Type() != val.t {
			PrintExpected(t, val.t, val.l.Type(), string(val.l.rune))
		}
	}
}

func TestLen(t *testing.T) {
	var tests = []struct {
		ls LiteralString
		ln int
	}{
		{litStr1, 4},
		{LiteralString{}, 0},
	}
	for _, val := range tests {
		result := val.ls.Len()
		if result != val.ln {
			PrintExpected(t, val.ln, result, val.ls.String(), val.ls)
		}
	}
}

func TestGetClassByName(t *testing.T) {
	var tests = []struct {
		s string
		t TokenClass
	}{
		{"KEYWORD", KEYWORD},
		{"SIGN", SIGN},
		{"DELIMETER", DELIMETER},
		{"TEXT", TEXT},
		{"NUMBER", NUMBER},
	}
	for _, val := range tests {
		result := GetClassByName(val.s)
		if result != val.t {
			PrintExpected(t, val.t, result, val.s)
		}
	}
}

func TestHandleSource(t *testing.T) {
	testsHS := [][2]string{
		{"TestString", "teststring"},
		{"space     space", "space space"},
		{"one\ttab", "one tab"},
		{"tab \t space", "tab space"},
		{"new\nline", "new line"},
		{"AT \n \t Once", "at once"},
	}
	for _, pair := range testsHS {
		result := (*HandleSource(pair[0])).String()
		if result != pair[1] {
			PrintExpected(t, pair[1], result, pair)
		}
	}
}

func TestCheckRune(t *testing.T) {
	type valChR struct {
		Ch   rune
		Text string
		Ret  bool
	}

	var testsCheckR = []valChR{
		{'a', "aaabbcdcaca", true},
		{'b', "", true},
		{' ', "abb", true},
		{' ', "aer ", false},
		{13, "abb", false},
		{13, "", false},
	}
	for _, test := range testsCheckR {
		result := checkRune(test.Ch, []rune(test.Text))
		if result != test.Ret {
			PrintExpected(t, test.Ret, result, test)
		}
	}
}

func TestTransformSpecSymbols(t *testing.T) {
	var tests = [][2]rune{
		{'a', 'a'},
		{' ', ' '},
		{'\t', ' '},
		{'\n', ' '},
	}
	for _, test := range tests {
		result := transformSpecSymbols(test[0])
		if result != test[1] {
			PrintExpected(t, string(test[1]), string(result), string(test[0]))
		}
	}
}

func TestCompare(t *testing.T) {
	var tests = []struct {
		c1 LiteralString
		c2 string
	}{
		{LiteralString{Literal{'a', 0}, Literal{'b', 1}, Literal{'c', 2}}, "abc"},
		{LiteralString{Literal{'A', 0}, Literal{'B', 1}, Literal{' ', 2}, Literal{'c', 3}}, "AB c"},
	}
	for _, val := range tests {
		result := compare(val.c1, val.c2)
		if result != reflect.DeepEqual(val.c1.String(), val.c2) {
			PrintExpected(t, reflect.DeepEqual(val.c1.String(), val.c2), result, val.c1.String(), val.c2)
		}
	}
}

/*func TestAnalyzeBuffer(t *testing.T) {
	var tests = []struct {
		s string
		l Lexeme
	}{
		{"begin", Lexeme{Token{"begin", KEYWORD, "begin"}, "begin"}},
	}
	lexemeTypes, syntaxRules := utils.ParseXML(utils.Loadfile(utils.TermsFilePath, utils.TermsFileName))
	SetLexTypes(lexemeTypes)
	for _, test := range tests {
		result := analyzeBuffer(*NewLiteralString(test.s))
		if !reflect.DeepEqual(result, test.l) {
			PrintExpected(t, test.l, result, test.s)
		}
	}
}
*/
