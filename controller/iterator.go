package controller

import (
	"errors"
	"gosyn/lexer"
	"gosyn/models"
	"gosyn/output"
	. "gosyn/output/colorman"
	"strconv"
	"strings"
)

var (
	deep int = 1
)

func deepInc() { deep++ }
func deepDec() { deep = deep - 1 }

type SeqIterator struct {
	int
	parent *SeqIterator
	buffer models.PTreeNode
}

func InitIter(parent *SeqIterator) (*SeqIterator, error) {
	child := *NewIterator(parent.int)
	child.parent = parent
	if child.IsInRange() {
		return &child, nil
	} else {
		return &child, errors.New("OUT OF RANGE. InitIter, .int: " + strconv.Itoa(child.int))
	}
}

func NewIterator(value int) *SeqIterator {
	return &SeqIterator{value, nil, models.NewTreeNode()}
}

func (i *SeqIterator) IsInRange() bool {
	if i.int >= 0 && i.int < len(seq) {
		return true
	} else {
		return false
	}
}

func (i *SeqIterator) Inc() {
	if i.IsInRange() {
		i.int++
	} else {
		output.PrintString(deep, COL_RED_BB, "OUT OF RANGE Inc, .int: "+strconv.Itoa(i.int))
	}
}

func (i *SeqIterator) ApplyToParent() {
	i.parent.int = i.int
	(*(i.parent.buffer)).AddTreeNode(i.buffer)
	output.Par.Send(0, COL_DEEP_LIM, "APPLY", string(i.parent.buffer.Type()), strings.TrimSpace(i.parent.buffer.Value), i.buffer.Type())
}

func (i *SeqIterator) AddToParent(word *models.StringElement, lex *lexer.Lexeme) {
	i.Inc()
	i.parent.int = i.int
	(*(i.parent.buffer)).AddTreeValue(word, lex)
	output.Par.Send(0, COL_DEEP_LIM, "ADD", string(i.parent.buffer.Type()), strings.TrimSpace(i.parent.buffer.Value), i.parent.int, i.buffer.Value)
	//i.parent.Inc()
}

func (i *SeqIterator) GetElement() (*lexer.Lexeme, error) {
	if i.IsInRange() {
		//output.PrintString(0, "GetElement ", seq[i.int].Name, (*(seq[i.int].Value)).String())
		return &seq[i.int], nil
	} else {
		return nil, errors.New("OUT OF RANGE GetVal, .int: " + strconv.Itoa(i.int))
	}
}
