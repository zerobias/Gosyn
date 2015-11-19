package controller

import (
	"fmt"
	"gosyn/lexer"
	"gosyn/models"
)

func StepSliceConv(ses []models.StringElement) *[]Step {
	result := make([]Step, 0)
	for _, val := range ses {
		result = append(result, *StepConvert(&val))
	}
	return &result
}

func StepConvert(se *models.StringElement) *Step {
	val := NewStep(se.Value)
	switch se.Type() {
	case models.ST_SEQ:
		result := *seqConvert(se)
		ret := Step(&result)
		return &ret
	case models.ST_CLASS:
		result := ClassStep(val)
		ret := Step(&result)
		return &ret
	case models.ST_RULE:
		result := RuleStep(val)
		ret := Step(&result)
		return &ret
	case models.ST_TERM:
		result := TermStep(val)
		ret := Step(&result)
		return &ret
	}

	return nil
}

func seqConvert(se *models.StringElement) *SeqStep {
	step := SeqStep{make([]Step, 0, len(se.Words)), NewOptions(se.Optional, se.Choises, se.Iterative)}
	for _, childSE := range se.Words {
		step.ChildSteps = append(step.ChildSteps, *StepConvert(&childSE))
	}
	return &step
}

type Step interface {
	StepType() models.DataType
	String() string
	HasChilds() (bool, *[]Step)
}

type SeqStep struct {
	ChildSteps []Step
	options    Options
}

//Flyweight Data Type. Does not consume memory
func (ss *SeqStep) StepType() models.DataType  { return models.ST_SEQ }
func (ss *SeqStep) String() string             { return "SEQ value" }
func (ss *SeqStep) HasChilds() (bool, *[]Step) { return true, &ss.ChildSteps }

//Each command in lang.xml, except 'seq', written as "type of data: typed value"
type SimpleStep string

type RuleStep SimpleStep

//Flyweight Data Type. Does not consume memory
func (rs *RuleStep) StepType() models.DataType { return models.ST_RULE }
func (rs *RuleStep) String() string            { return string(*rs) }
func (rs *RuleStep) HasChilds() (bool, *[]Step) {
	val := (GetRule(rs.String())).TopWord.Words
	fmt.Println(val)
	if val != nil {
		return true, StepSliceConv(val)
	} else {
		return false, nil
	}
}

type TermStep SimpleStep

//Flyweight Data Type. Does not consume memory
func (ts *TermStep) StepType() models.DataType  { return models.ST_TERM }
func (ts *TermStep) String() string             { return string(*ts) }
func (ts *TermStep) HasChilds() (bool, *[]Step) { return false, nil }

type ClassStep SimpleStep

//Flyweight Data Type. Does not consume memory
func (cs *ClassStep) StepType() models.DataType  { return models.ST_CLASS }
func (cs *ClassStep) String() string             { return string(*cs) }
func (cs *ClassStep) HasChilds() (bool, *[]Step) { return false, nil }

func NewStep(value string) SimpleStep {
	ss := *new(SimpleStep)
	ss = SimpleStep(value)
	return ss
}

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
}

//---------------------

type OptionType string

func NewOptions(o, c, i bool) Options {
	opt := make(map[OptionType]bool, 3)
	opt[Optional] = o
	opt[Choises] = c
	opt[Iterative] = i
	return Options(opt)
}

func (ot OptionType) String() string { return string(ot) }

const (
	Optional  OptionType = "optional"
	Choises              = "choises"
	Iterative            = "iterative"
)

type Options map[OptionType]bool
type Optioner interface {
	HasNonDefaultOpt() bool
	Options() Options
}
