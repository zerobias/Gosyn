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

func (df *DataFacade) StepValue() string         { return (*df.step).String() }
func (df *DataFacade) StepType() models.DataType { return (*df.step).StepType() }
func (df *DataFacade) HasChilds() bool           { return (*df.step).HasChilds() }
func (df *DataFacade) Childs() *[]Step           { return (*df.step).Childs() }
func (df *DataFacade) Child() *lexer.Lexeme      { return df.lexeme }
func (df *DataFacade) Add(child DataFacade) {
	*(df.childs) = append(*(df.childs), child)
}

func (df *DataFacade) AddText(step *Step, lex *lexer.Lexeme) {
	df.Add(*NewFacadeText(step, lex))
}

func (df *DataFacade) GetStep() *Step { return df.step }

func (df *DataFacade) GetLex() *lexer.Lexeme { return df.lexeme }
