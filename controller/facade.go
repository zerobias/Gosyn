package controller

import (
	"gosyn/lexer"
	"gosyn/models"
)

func NewFacade(step *Step, lexeme *lexer.Lexeme, childs *[]DataFacade) *DataFacade {
	return &DataFacade{step, lexeme, childs}
}

func NewFacadeList(step *Step, childs *[]DataFacade) *DataFacade {
	return &DataFacade{step, nil, childs}
}

func NewFacadeText(step *Step, lexeme *lexer.Lexeme) *DataFacade {
	return &DataFacade{step, lexeme, nil}
}

func NewFacadeStep(step Step) *DataFacade {
	return &DataFacade{&step, nil, nil}
}

type DataFacade struct {
	step   *Step
	lexeme *lexer.Lexeme
	childs *[]DataFacade
}

func (df *DataFacade) StepValue() string          { return (*df.step).String() }
func (df *DataFacade) StepType() models.DataType  { return (*df.step).StepType() }
func (df *DataFacade) HasChilds() (bool, *[]Step) { return (*df.step).HasChilds() }

func (df *DataFacade) Add(child DataFacade) {
	*(df.childs) = append(*(df.childs), child)
}

func (df *DataFacade) AddText(word *models.StringElement, lex *lexer.Lexeme) {
	df.Add(*NewFacadeText(StepConvert(word), lex))
}

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
