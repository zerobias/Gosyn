package controller

import (
	"llsp/lexer"
	//"llsp/models"
	. "llsp/testingtools"
	"llsp/utils"
	"testing"
)

func init() {
	InitTest()
}

func InitTest() {
	var lexType_test lexer.TokenList
	lexType_test, rules = utils.ParseXML(utils.Loadfile(utils.TermsFilePath, utils.TermsFileName))
	lexer.SetTokens(lexType_test)
	seq = *lexer.ScanCycle(lexer.HandleSource(string(utils.Loadfile("C:\\_projects\\Go\\bin", "code1.txt"))))
}

func TestIsError(t *testing.T) {
	iter, error := InitIter(NewIterator(len(seq)))
	if !isError(error) {
		PrintExpected(t, true, false, len(seq), error, *iter)
	}
	iter_, error := InitIter(NewIterator(0))
	if isError(error) {
		PrintExpected(t, false, true, error, *iter_)
	}
}

func TestGetRule(t *testing.T) {
	if GetRule("x") != nil {
		PrintExpected(t, "GetRule must be nil", GetRule("x"), "x")
	}
	var tests = []struct {
		rname     string
		childType string
	}{
		{"label", "seq"},
		{"program", "seq"},
	}
	for _, val := range tests {
		result := *GetRule(val.rname)
		if result.Name != val.rname || string(result.TopWord.Type()) != val.childType {
			PrintExpected(t, val, result, val.rname)
		}
	}
}

func TestGetElement(t *testing.T) {
	curs := *NewIterator(0)
	str := "program"
	result, err := curs.GetElement()
	if result.Name != str {
		PrintExpected(t, str, result, err)
	}
	curs = *NewIterator(len(seq))
	res, err := curs.GetElement()
	if !isError(err) {
		PrintExpected(t, true, false, len(seq), err, *res, curs)
	}
}
