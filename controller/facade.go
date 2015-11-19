package controller

import (
	"gosyn/lexer"
	"gosyn/models"
)

func NewFacade(step *Step, lexeme *lexer.Lexeme, childs *[]DataFacade) *DataFacade {
	return &DataFacade{step, lexeme, childs}
}

type DataFacade struct {
	step   *Step
	lexeme *lexer.Lexeme
	childs *[]DataFacade
}

func (df *DataFacade) StepValue() string          { return (*df.step).String() }
func (df *DataFacade) StepType() models.DataType  { return (*df.step).StepType() }
func (df *DataFacade) HasChilds() (bool, *[]Step) { return (*df.step).HasChilds() }

/*func (dc *DClass) RuleMatch(parent *SeqIterator) bool {
	_, element, e := GetCurrent(parent)
	if isError(e) {
		return false
	}
	return lexer.GetClassByName(dc.value) == element.Cat
}

func (dt *DTerm) RuleMatch(parent *SeqIterator) bool {
	_, element, e := GetCurrent(parent)
	if isError(e) {
		return false
	}
	return dt.Value() == element.Name
}

func (dr *DRule) RuleMatch(parent *SeqIterator) bool {
	cursor, _, e := GetCurrent(parent)
	if isError(e) {
		return false
	}
	return Translate(&GetRule(dr.Value()).TopWord, cursor)
}*/
/*
func GetCurrent(parent *SeqIterator) (*SeqIterator, *lexer.Lexeme, error) {
	i, e := InitIter(parent)
	if isError(e) {
		return nil, nil, e
	}
	l, e := i.GetElement()
	if isError(e) {
		return i, nil, e
	}
	return i, l, e
}*/
