package output

import (
	//"io"
	"bufio"
	"fmt"
	//"gosyn/debug"
	//"gosyn/models"
	"os"
	"strings"
	"text/tabwriter"
	/*"reflect"
	"strconv"

	"unicode/utf8"*/)

var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var tabWriter *tabwriter.Writer

func InitWriter() {
	tabWriter = tabwriter.NewWriter(os.Stdout, 14, 14, 0, '\t', 0)
}

func out(str string) {
	fmt.Fprintf(tabWriter, "%v\n", str)
	//writeLine(tabWriter, str)
}

func CloseWriter() {
	tabWriter.Flush()
}

func writeLine(w *tabwriter.Writer, s string) {
	tabs := strings.Count(s, "\t")
	if tabs < 2 {
		appendTabs := strings.Repeat("\t", 2-tabs)
		s = strings.Replace(s, "\n", appendTabs+"\n", 1)
	}
	fmt.Fprintf(w, "%v\n", s)
}

/*var (
	headers   []string
	tabWriter = tabwriter.NewWriter(os.Stdout, 10, 8, 0, '\t', 0)
)

func alreadyStored(header ...string) bool {
	joined := strings.Join(header, "\t")
	for _, head := range headers {
		if joined == head {
			return true
		}
	}
	headers = append(headers, joined)
	return false
}

func AddQuotes(val interface{}) string {
	return "'" + convertObject(val) + "'"
}

func convertObject(obj interface{}) string {
	switch reflect.TypeOf(obj).Kind() {
	case reflect.String:
		return reflect.ValueOf(obj).Interface().(string)
	case reflect.Int:
		return strconv.Itoa(reflect.ValueOf(obj).Interface().(int))
	case reflect.Int32:
		intStr := strconv.FormatInt(int64(reflect.ValueOf(obj).Interface().(int32)), 16)
		//out(intStr)
		decoded, _ := utf8.DecodeRuneInString(intStr)
		return string(decoded)
	case reflect.Bool:
		return strconv.FormatBool(reflect.ValueOf(obj).Interface().(bool))
	case reflect.Slice:
		sliceVal := reflect.ValueOf(obj)
		//fmt.Fprintln(tabWriter, "_________________________")
		typeCast := func(casted reflect.Value) string {
			switch casted.Interface().(type) {
			case models.LexemeSequence:
				var result string
				cast := casted.Interface().(models.LexemeSequence)
				//fmt.Fprintln(tabWriter, "SLICESEQ", len(cast))
				for _, lex := range cast {
					result += lex.Value
				}
				return result
			case []rune:
				return string(casted.Interface().([]rune))
			case []string:
				return joinWithTabs(casted.Interface().([]string))
			default:
				fmt.Fprintln(tabWriter, "DEFAULT")
			}
			return ""
		}
		j := typeCast(sliceVal)
		//fmt.Fprintln(tabWriter, j)
		return j
	case reflect.Struct:
		var structString string
		refStruct := reflect.ValueOf(obj)
		convStruct := refStruct.Interface().(models.Lexeme)
		fmt.Fprintln(tabWriter, "TypeOf", reflect.TypeOf(obj), "conv", convStruct.Value)
		return structString
	}
	return "NULL"
}

func convertArgsInterface(val ...interface{}) (result []string) {
	for _, obj := range val {
		result = append(result, convertObject(obj))
	}
	return
}

func joinWithTabs(text []string) string {
	return strings.Join(text, "\t")
}

type Color string

const (
	BOLDRED Color = "\x1b[31;1m"
	DEF     Color = "\x1b[0m"
	GRAY    Color = "\x1b[30;1m"
)

func ColorPrint(col Color, val ...interface{}) {
	str := joinWithTabs(convertArgsInterface(val...))
	out(addColors(col, str))
}

func addColors(col Color, str string) string {
	return string(col) + str + string(DEF)
}

func PrintString(val ...interface{}) {
	j := joinWithTabs(convertArgsInterface(val...))
	out(j)
}

func Print(header []string, val ...interface{}) {
	if !alreadyStored(header...) {
		fmt.Fprintln(tabWriter, addColors(BOLDRED, strings.Join(header, "\t")))
	}
	PrintString(val...)
}
*/
