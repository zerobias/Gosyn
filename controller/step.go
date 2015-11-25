package controller

import (
	"errors"
	"fmt"
	"gosyn/lexer"
	"gosyn/models"
)

//Storage type for elements of rules descriptions
type Step interface {
	//Flyweight Data Type. Does not consume memory
	StepType() models.DataType
	//Stringer interface of type
	String() string
	HasChilds() bool
	//Get child steps
	Childs() *[]Step
	//Execute rule with current step
	RuleMatch(parent *SeqIterator) bool
	ParentGet() *Step
}
type NullStep struct {
	parent *Step
}

func (ns NullStep) ParentGet() *Step { return ns.parent }

//func (ns NullStep) Siblings() *[]Step { return (*ns.parent).Childs() }

/*func (ns NullStep) IsLast() bool {
	return (*ns.siblings)[len(ns.siblings()) - 1] ==
}*/

type Iter interface {
	IsLast() bool
	ParentSet(parent Step)
	ParentGet(parent Step)
}

//Each command in lang.xml, except 'seq', written as "type of data: typed string"
type SimpleStep struct {
	string
	NullStep
}

type TermStep SimpleStep
type ClassStep SimpleStep
type RuleStep SimpleStep
type SeqStep struct {
	ChildSteps []Step
	options    Options
	NullStep
}

func (ts TermStep) StepType() models.DataType  { return models.ST_TERM }
func (cs ClassStep) StepType() models.DataType { return models.ST_CLASS }
func (rs RuleStep) StepType() models.DataType  { return models.ST_RULE }
func (ss SeqStep) StepType() models.DataType   { return models.ST_SEQ }

func (ts TermStep) String() string  { return ts.string }
func (cs ClassStep) String() string { return cs.string }
func (rs RuleStep) String() string  { return rs.string }
func (ss SeqStep) String() string   { return "SEQ value" }

func (ts TermStep) HasChilds() bool  { return false }
func (cs ClassStep) HasChilds() bool { return false }
func (rs RuleStep) HasChilds() bool {
	val := (GetRule(rs.String())).TopWord.Words
	return val != nil
}
func (ss SeqStep) HasChilds() bool { return true }

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

func IsLast(step *Step) bool {
	p := *(*(*step).ParentGet()).Childs()
	return p[len(p)-1] == *step
}

func (ts TermStep) ParentGet() *Step  { return ts.NullStep.ParentGet() }
func (cs ClassStep) ParentGet() *Step { return cs.NullStep.ParentGet() }
func (rs RuleStep) ParentGet() *Step  { return rs.NullStep.ParentGet() }
func (ss SeqStep) ParentGet() *Step   { return ss.NullStep.ParentGet() }

func NewStep(value string) SimpleStep {
	return SimpleStep{value, NullStep{}}
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
	step := SeqStep{make([]Step, 0, len(se.Words)), NewOptions(se.Optional, se.Choises, se.Iterative), NullStep{}}
	for _, childSE := range se.Words {
		step.ChildSteps = append(step.ChildSteps, StepConvert(&childSE))
	}
	return step
}
