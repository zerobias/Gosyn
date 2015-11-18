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
	"strconv"
)

var (
	rules []models.Rule
	//source models.ParsedSentence
	seq lexer.LexemeSequence
)

func SyntaxCycle(sequence *lexer.LexemeSequence, syntaxRules []models.Rule) models.PTreeNode {
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
	startRule := "program"
	output.PrintString(0, *GetRule(startRule)) //TODO catch nil
	output.Par = *new(output.Paragraph)
	output.Par.Init()
	cursor := *NewIterator(0)
	Translate(&GetRule(startRule).TopWord, &cursor)
	if len(cursor.buffer.ChildList) == 0 {
		output.Par.Send(0, COL_RED_BB, "ZERO")
	}
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

func Translate(word *models.StringElement, parentCursor *SeqIterator) (result bool) {
	deepInc()
	defer deepDec()

	cursor, element, error := GetCurrent(parentCursor)
	if isError(error) {
		return false
	}
	//fmt.Println("Element ", (*element.Value).String())
	cursor.buffer.StringElement = *word
	switch word.Type() {
	case models.ST_CLASS:
		output.Par.SendPair(deep-1, COL_BLUE_B, "CLASS", word.Value,
			"curs.i", strconv.Itoa(cursor.int), (*element.Value).String())
		result = lexer.GetClassByName(word.Value) == element.Cat
	case models.ST_TERM:
		output.Par.SendPair(deep-1, COL_BLUE_B, "TERM", word.Value, (*element.Value).String())
		result = element.Name == word.Value
	case models.ST_RULE:
		output.Par.Send(deep-1, COL_BLUE_B, "RULE", word.Value)
		result = Translate(&GetRule(word.Value).TopWord, cursor)
	case models.ST_SEQ:
		output.Par.SendPair(deep, COL_BLUE_U, "SEQ", deep,
			"O", word.Optional,
			"C", word.Choises, "I", word.Iterative)
		checkChilds := func(curs *SeqIterator) bool {
			isLast := func(n int) bool {
				if n == len(word.Words)-1 {
					return true
				} else {
					return false
				}
			}
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
	}
	if result {
		if word.Type() == models.ST_RULE ||
			word.Type() == models.ST_SEQ {
			cursor.ApplyToParent()
		} else {
			fmt.Println("Add to parent", *word, "|", *element)
			cursor.AddToParent(word, element)
		}
		output.Par.SendPair(deep, COL_GREEN, "RES", result, word.Type()) //TODO print output for failures too
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
