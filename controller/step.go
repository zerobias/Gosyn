package controller

import (
	"errors"
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
	HasChilds() bool
	Childs() *[]Step
	RuleMatch(parent *SeqIterator) bool
}

type SeqStep struct {
	ChildSteps []Step
	options    Options
}

//Flyweight Data Type. Does not consume memory
func (ss SeqStep) StepType() models.DataType { return models.ST_SEQ }
func (ss SeqStep) String() string            { return "SEQ value" }
func (ss SeqStep) HasChilds() bool           { return true }
func (ss SeqStep) Childs() *[]Step           { return &ss.ChildSteps }

//Each command in lang.xml, except 'seq', written as "type of data: typed value"
type SimpleStep string

type RuleStep SimpleStep

//Flyweight Data Type. Does not consume memory
func (rs *RuleStep) StepType() models.DataType { return models.ST_RULE }
func (rs *RuleStep) String() string            { return string(*rs) }
func (rs *RuleStep) HasChilds() bool {
	val := (GetRule(rs.String())).TopWord.Words
	return val != nil
}

func (rs *RuleStep) Childs() *[]Step {
	val := (GetRule(rs.String())).TopWord.Words
	fmt.Println(val)
	if val != nil {
		return StepSliceConv(val)
	} else {
		return nil
	}
}

type TermStep SimpleStep

//Flyweight Data Type. Does not consume memory
func (ts *TermStep) StepType() models.DataType { return models.ST_TERM }
func (ts *TermStep) String() string            { return string(*ts) }
func (ts *TermStep) HasChilds() bool           { return false }
func (ts *TermStep) Childs() *[]Step           { return nil }

type ClassStep SimpleStep

//Flyweight Data Type. Does not consume memory
func (cs *ClassStep) StepType() models.DataType { return models.ST_CLASS }
func (cs *ClassStep) String() string            { return string(*cs) }
func (cs *ClassStep) HasChilds() bool           { return false }
func (cs *ClassStep) Childs() *[]Step           { return nil }

func (cs *ClassStep) RuleMatch(parent *SeqIterator) bool {
	_, element, e := GetCurrent(parent)
	if isError(e) {
		return false
	}
	return lexer.GetClassByName(cs.String()) == element.Cat
}

func (ts *TermStep) RuleMatch(parent *SeqIterator) bool {
	_, element, e := GetCurrent(parent)
	if isError(e) {
		return false
	}
	return ts.String() == element.Name
}

func (rs *RuleStep) RuleMatch(parent *SeqIterator) bool {
	/*cursor*/ _, _, e := GetCurrent(parent)
	if isError(e) {
		return false
	}
	/*rule*/ _, err := GetRuleStep(rs.String())
	if isError(err) {
		return false
	}
	return false //Translate(rule, cursor)
}

func (ts SeqStep) RuleMatch(parent *SeqIterator) bool {
	var result bool
	cursor, _ /*element*/, e := GetCurrent(parent)
	if isError(e) {
		return false
	}
	checkChilds := func(curs *SeqIterator) bool {
		isLast := func(n int) bool { return n == len(*(ts.Childs()))-1 }
		for i, child := range *(ts.Childs()) {
			childResult := Translate(&child, curs)
			if ts.options[Choises] {
				if childResult {
					return true
				} else if isLast(i) {
					return false
				}
			} else {
				if childResult {
					if isLast(i) {
						return true
					}
				} else {
					return false
				}
			}
		}
		return false
	}

	if ts.options[Iterative] {
		Nrep := 0
		cycleCursor, error := InitIter(cursor)
		cycleCursor.buffer = *NewFacadeStep(ts) //string(models.ST_SEQ)
		if !isError(error) {
			for iterSucc := true; iterSucc; {
				iterSucc = checkChilds(cycleCursor)
				if iterSucc {
					Nrep++
				}
			}
		}
		if Nrep > 0 {
			cycleCursor.ApplyToParent()
			result = true
		} else {
			result = false
		}
	} else {
		result = checkChilds(cursor)
	}
	if ts.options[Optional] && !result {
		result = true
		cursor.buffer = *NewFacadeStep(SeqStep{})
	}

	return false
}

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
