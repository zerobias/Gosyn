package output

import (
	//"io"
	//"bufio"
	//"gosyn/debug"
	"gosyn/models"
	. "gosyn/output/colorman"
	//"strconv"
	//"encoding/xml"
	"strings"
	//"text/tabwriter"
	/*"os"
	"reflect"
	"unicode/utf8"*/)

var Par Paragraph

var elementColors = map[string][]ColorValue{
	"seq":   {COL_BLUE_B, COL_BLUE_U},
	"term":  {RESET, COL_GREEN},
	"rule":  {COL_CYAN, COL_GREY},
	"class": {COL_GREEN, RESET},
}

func printStringElement(level int, word models.StringElement) {
	if word.Type() == models.ST_SEQ {
		for _, child := range word.Words {
			Par.Send(level, COL_BLUE_B, AddColors(COL_BLUE_U, "seq"))
			printStringElement(level+1, child)
		}
	} else {
		Par.Send(level, elementColors[word.XMLName.Local][0], //TODO simplify type string call
			AddColors(elementColors[word.XMLName.Local][1], word.XMLName.Local), word.Value)
		//PrintString(level, word)
	}
}

func PrintRules(rules []models.Rule) {
	PrintSectionTitle("STORED", "RULES")
	Par = *new(Paragraph)
	Par.Init()
	for _, rule := range rules {
		//ColorPrint(RED, 0, "RULE", rule.Name)
		Par.Send(0, COL_RED, AddColors(COL_RED_BB, "RULE"), rule.Name)
		//PrintString(0, rule)
		printStringElement(0, rule.TopWord)
	}
	Par.Write()
}

func PrintSectionTitle(text ...string) {
	const line = "------"
	ColorPrint(COL_RED, 0, line, joinWithTabs(text...), line)
}

func ColorPrint(col ColorValue, level int, val ...interface{}) {
	str := getIndent(level/2) + joinWithTabs(models.CastToStrings(val...)...)
	out(AddColors(col, str))
}

func PrintString(level int, val ...interface{}) {
	ColorPrint(RESET, level, val)
}

func PrintInLine(val ...interface{}) {
	out(joinWithSpace(models.CastToStrings(val...)...))
}

func joinWithTabs(text ...string) string {
	return strings.Join(text, "\t")
}

func joinWithSpace(text ...string) string {
	return strings.Join(text, " ")
}

func getIndent(level int) (result string) {
	const tab = "\t"
	for i := 0; i < level; i++ {
		result += tab
	}
	return
}
