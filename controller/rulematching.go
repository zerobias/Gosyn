package controller

import (
	//"fmt"
	"gosyn/lexer"
	//"gosyn/models"
)

func (ts TermStep) RuleMatch(parent *SeqIterator) bool {
	_, element, e := GetCurrent(parent)
	if isError(e) {
		return false
	}
	return ts.String() == element.Name
}

func (cs ClassStep) RuleMatch(parent *SeqIterator) bool {
	_, element, e := GetCurrent(parent)
	if isError(e) {
		return false
	}
	return lexer.GetClassByName(cs.String()) == element.Cat
}

func (rs RuleStep) RuleMatch(parent *SeqIterator) bool {
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
		for _, child := range *(ts.Childs()) {
			childResult := Translate(&child, curs)
			if ts.options[Choises] {
				if childResult {
					return true
				} else if IsLast(&child) {
					return false
				}
			} else {
				if childResult {
					if IsLast(&child) {
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
