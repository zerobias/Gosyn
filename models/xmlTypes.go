package models

import (
	"encoding/xml"
	//"reflect"
	//"fmt"
	//"strconv"
)

type Terminal struct {
	Name  string `xml:"name,attr"`
	Class string `xml:"class,attr"`
	Value string `xml:",chardata"`
} //todo move to XML processing, make private

//-------------------------------------------------------------

type DataType string

const (
	ST_SEQ   DataType = "seq"
	ST_RULE           = "rule"
	ST_TERM           = "term"
	ST_CLASS          = "class"
)

func (se *StringElement) Type() DataType {
	switch DataType(se.XMLName.Local) {
	case ST_SEQ:
		return ST_SEQ
	case ST_RULE:
		return ST_RULE
	case ST_TERM:
		return ST_TERM
	case ST_CLASS:
		return ST_CLASS
	}
	return DataType("NO TYPE")
}

type StringElement struct {
	XMLName   xml.Name
	Choises   bool            `xml:"choices,attr"`
	Iterative bool            `xml:"iterative,attr"`
	Optional  bool            `xml:"optional,attr"`
	Value     string          `xml:",chardata"`
	Words     []StringElement `xml:",any"`
}

type Rule struct {
	XMLName xml.Name      `xml:"RULE"`
	Name    string        `xml:"name,attr"`
	TopWord StringElement `xml:"seq"` //TODO как в XMLRoot Rules
	//TermNames []string `xml:"term"`
}

type XMLRoot struct {
	XMLName xml.Name   `xml:"ROOT"`
	Terms   []Terminal `xml:"TERMINALS>TERM"`
	Rules   []Rule     `xml:"RULES>RULE"`
}
