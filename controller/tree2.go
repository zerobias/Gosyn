package controller

import (
	"llsp/lexer"
	"llsp/models"
)

type SchemaData interface {
	DataType() models.DataType
	RuleMatch(parent *SeqIterator) bool
	Data() Container
}

type DataObject interface {
	SchemaData
	Value() string
}

type DataSet interface {
	//Optioner
	SchemaData
	Value() []Container
}

type Container interface {
	IsObject() bool
	Object() DataObject
	Array() DataSet
}

type DataContainer struct {
	object   DataObject
	set      DataSet
	dataType models.DataType
}

func (dc *DataContainer) IsObject() bool {

}

type value struct {
	value string
}

type (
	DClass value
	DTerm  value
	DRule  value
)

func (v *value) Value() string {
	return v.value
}

/*func (dc *DClass) Value() string {
	return dc.value
}

func (dt *DTerm) Value() string {
	return dt.value
}

func (dr *DRule) Value() string {
	return dr.value
}*/

//Flyweight Data Type. Does not consume memory
func (dc DClass) DataType() models.DataType {
	return models.ST_CLASS
}

//Flyweight Data Type. Does not consume memory
func (dt DTerm) DataType() models.DataType {
	return models.ST_TERM
}

//Flyweight Data Type. Does not consume memory
func (dr DRule) DataType() models.DataType {
	return models.ST_RULE
}

func (dc *DClass) RuleMatch(parent *SeqIterator) bool {
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
