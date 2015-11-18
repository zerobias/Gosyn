package debug

import (
	"fmt"
	//"io"
	"gosyn/lexer"
	"gosyn/models"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/tabwriter"
	"unicode/utf8"
)

var (
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
func ConvertTypeLayer(obj ...interface{}) []string {
	return convertArgsInterface(obj...)
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
			case lexer.LexemeSequence:
				var result string
				cast := casted.Interface().(lexer.LexemeSequence)
				//fmt.Fprintln(tabWriter, "SLICESEQ", len(cast))
				for _, lex := range cast {
					result += (*lex.Value).String()
				}
				return result
			case []rune:
				return string(casted.Interface().([]rune))
			case []string:
				return joinWithTabs(casted.Interface().([]string))
			case []models.Rule:
				return "Rule[]" //models.GetString(obj)
			case models.Rule:
				fmt.Fprintln(tabWriter, "rule")
				return "Rule"
			default:
				fmt.Fprintln(tabWriter, "DEFAULT")
				fmt.Fprintln(tabWriter, "defObject: ", obj)
				fmt.Fprintln(tabWriter, "ValOfDef: ", sliceVal)
			}
			return "def"
		}
		j := typeCast(sliceVal)
		//fmt.Fprintln(tabWriter, j)
		return j
	case reflect.Struct:
		var structString string
		refStruct := reflect.ValueOf(obj)
		switch refStruct.Interface().(type) {
		case lexer.Lexeme:
			convStruct := refStruct.Interface().(lexer.Lexeme)
			fmt.Fprintln(tabWriter, "TypeOf", reflect.TypeOf(obj), "conv", convStruct.Value)
			structString = (*(convStruct.Value)).String()
		case models.Rule:
			fmt.Fprintln(tabWriter, "struct rule")
		default:
			structString = "Struct default"
		}

		return structString
	default:
		return "debug switch default"
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

func out(str string) {
	fmt.Fprintf(tabWriter, "%v\n", str)
	tabWriter.Flush()
}
