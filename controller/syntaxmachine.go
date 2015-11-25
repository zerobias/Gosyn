package controller

import (
	//"errors"
	"fmt"
	"gosyn/lexer"
	"gosyn/models"
	"gosyn/output"
	. "gosyn/output/colorman"
	/*"log"
	"os"
	"os/exec"*/
	//"strconv"
)

var (
	rules []models.Rule
	//source models.ParsedSentence
	seq lexer.LexemeSequence
)

//Init and control of parsing
func SyntaxCycle(sequence *lexer.LexemeSequence, syntaxRules []models.Rule) DataFacade {
	//source = *sequence.Transform()
	seq = *sequence
	rules = syntaxRules
	/*output.PrintString(0, "\nSyntaxCycle SEQ")
	for _, lex := range []lexer.Lexeme(*sequence) {
		output.PrintString(0, lex.Name, (*(lex.Value)).String())
	}*/
	/*for _, ptr := range source.GetChilds() {
		fmt.Println((*ptr).Assocciated.Name)
	}*/
	startRuleName := "program"
	//output.PrintString(0, *GetRule(startRule)) //TODO catch nil
	initRule, err := GetRuleStep(startRuleName)
	if isError(err) {
		fmt.Errorf(err.Error(), startRuleName)
	}
	output.Par = *new(output.Paragraph)
	output.Par.Init()
	cursor := *NewIterator(0)
	Translate(initRule, &cursor)
	/*if len(cursor.buffer.ChildList) == 0 {
		output.Par.Send(0, COL_RED_BB, "ZERO")
	}*/

	//output.Par.Write()
	deep = 0

	output.Par = *new(output.Paragraph)
	output.Par.Init()
	//PrintTree(*cursor.buffer)
	output.Par.Write()

	return cursor.buffer
}

func isError(e error) bool {
	if e != nil {
		fmt.Errorf("Error!%v\n", e.Error())
		return true
	} else {
		return false
	}
}

//Main syntax parsing function
func Translate(step *Step, parentCursor *SeqIterator) (result bool) {
	deepInc()
	defer deepDec()

	cursor, element, error := GetCurrent(parentCursor)
	if isError(error) {
		return false
	}
	//fmt.Println("Element ", (*element.Value).String())
	cursor.buffer.step = step
	result = (*step).RuleMatch(parentCursor)
	/*switch step.StepType() {
	case models.ST_CLASS:
		result = lexer.GetClassByName(word.Value) == element.Cat
	case models.ST_TERM:
		result = element.Name == word.Value
	case models.ST_RULE:
		result = Translate(&GetRule(word.Value).TopWord, cursor)
	case models.ST_SEQ:
		checkChilds := func(curs *SeqIterator) bool {
			isLast := func(n int) bool { return n == len(word.Words)-1 }
			for i, child := range word.Words {
				childResult := Translate(&child, curs)
				if word.Choises {
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

		if word.Iterative {
			Nrep := 0
			cycleCursor, error := InitIter(cursor)
			cycleCursor.buffer.XMLName.Local = string(models.ST_SEQ)
			if !isError(error) {
				for iterSucc := true; iterSucc; {
					iterSucc = checkChilds(cycleCursor)
					output.Par.SendPair(deep, COL_RED, iterSucc,
						Nrep, cursor.int, cycleCursor.int)
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
		if word.Optional && !result {
			result = true
			cursor.buffer = models.NewTreeNode()
			cursor.buffer.XMLName.Local = string(models.ST_SEQ)
			output.Par.Send(0, COL_DEEP_LIM, "OPTIONAL")
		}
	}*/
	if result {
		if (*step).StepType() == models.ST_RULE ||
			(*step).StepType() == models.ST_SEQ {
			cursor.ApplyToParent()
		} else {
			fmt.Println("Add to parent", *step, "|", *element)
			cursor.AddToParent(step, element)
		}
	}
	/*if word.Type() == models.ST_SEQ {
		output.Par.SendPair(deep, COL_RED_BB, "result", cursor.buffer.StringElement.Type())
	} else {
		output.Par.SendPair(deep, COL_RED_BB, "result",
			cursor.buffer.StringElement.Type(), cursor.buffer.StringElement.Value)
	}*/
	return
}

func GetRule(name string) *models.Rule {
	if rules != nil {
		for _, val := range rules {
			if val.Name == name {
				return &val
			}
		}
	}
	return nil
}

func PrintTree(tree models.TreeNode) { //TODO make private or move to output (better)
	deepInc()
	defer deepDec()
	var valText string
	for _, val := range tree.ChildList {
		if val.Type == models.TREE_NODE {

			if val.AsNode().Type() == models.ST_SEQ {
				valText = ""
			} else {
				valText = val.AsNode().Value
			}
			output.Par.Send(deep, RESET+COL_RED, val.AsNode().Type(),
				valText)
			if deep < 12 {
				PrintTree(*val.AsNode())
			} else {
				output.Par.Send(deep, RESET,
					AppendColor("deep limit", COL_DEEP_LIM))
				break
			}
		} else {
			if val.AsValue() != nil {
				output.Par.Send(deep, RESET,
					val.AsValue().Type(),
					AppendColor(val.AsValue().Value.Name, COL_BLUE_U),
					val.AsValue().Value.Value)
			} else {
				output.Par.Send(deep, COL_RED_BB, "NIL")
			}
			// tree.GetTreeValue(*val).Value.Value)
		}
	}
}
