package controller

import (
	"errors"
	"fmt"
	"gosyn/lexer"
	"gosyn/models"
)

type Step interface {
	StepType() models.DataType
	String() string
	HasChilds() bool
	Childs() *[]Step
	RuleMatch(parent *SeqIterator) bool
}

//Each command in lang.xml, except 'seq', written as "type of data: typed value"
type SimpleStep string

type TermStep SimpleStep
type ClassStep SimpleStep
type RuleStep SimpleStep
type SeqStep struct {
	ChildSteps []Step
	options    Options
}

//Flyweight Data Type. Does not consume memory
func (ts TermStep) StepType() models.DataType  { return models.ST_TERM }
func (cs ClassStep) StepType() models.DataType { return models.ST_CLASS }
func (rs RuleStep) StepType() models.DataType  { return models.ST_RULE }
func (ss SeqStep) StepType() models.DataType   { return models.ST_SEQ }

//Stringer interface of type
func (ts TermStep) String() string  { return string(ts) }
func (cs ClassStep) String() string { return string(cs) }
func (rs RuleStep) String() string  { return string(rs) }
func (ss SeqStep) String() string   { return "SEQ value" }

func (ts TermStep) HasChilds() bool  { return false }
func (cs ClassStep) HasChilds() bool { return false }
func (rs RuleStep) HasChilds() bool {
	val := (GetRule(rs.String())).TopWord.Words
	return val != nil
}
func (ss SeqStep) HasChilds() bool { return true }

//Get child steps
func (ts TermStep) Childs() *[]Step  { return nil }
func (cs ClassStep) Childs() *[]Step { return nil }
func (rs RuleStep) Childs() *[]Step {
	val := (GetRule(rs.String())).TopWord.Words
	fmt.Println(val)
	if val != nil {
		return StepSliceConv(val)
	} else {
		return nil
	}
}
func (ss SeqStep) Childs() *[]Step { return &ss.ChildSteps }

func NewStep(value string) SimpleStep {
	ss := *new(SimpleStep)
	ss = SimpleStep(value)
	return ss
}

func GetRuleStep(name string) (*Step, error) {
	ruleStep := RuleStep(NewStep(name))
	step := Step(&ruleStep)
	has := step.HasChilds()
	if has {
		return &step, nil
	} else {
		return &step, errors.New("Empty rule! Name %s")
	}
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

//--------------------

func StepSliceConv(ses []models.StringElement) *[]Step {
	result := make([]Step, 0)
	for _, val := range ses {
		result = append(result, StepConvert(&val))
	}
	return &result
}

func StepConvert(se *models.StringElement) Step {
	val := NewStep(se.Value)
	switch se.Type() {
	case models.ST_SEQ:
		return seqConvert(se)
	case models.ST_CLASS:
		return ClassStep(val)
	case models.ST_RULE:
		return RuleStep(val)
	case models.ST_TERM:
		return TermStep(val)
	default:
		return nil
	}
}

func seqConvert(se *models.StringElement) SeqStep {
	step := SeqStep{make([]Step, 0, len(se.Words)), NewOptions(se.Optional, se.Choises, se.Iterative)}
	for _, childSE := range se.Words {
		step.ChildSteps = append(step.ChildSteps, StepConvert(&childSE))
	}
	return step
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
