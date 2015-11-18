package models

import (
	//"encoding/xml"
	"fmt"
	"llsp/lexer"
	"llsp/output/colorman"
	"reflect"
	"strconv"
)

func isSlice(obj interface{}) bool {
	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return true
	} else {
		return false
	}
}

func CastToStrings(obj ...interface{}) (result []string) {
	add := func(str ...string) {
		result = append(result, str...)
	}
	for _, val := range obj {
		/*if isSlice(val) {
			add(CastToStrings(reflect.ValueOf(val).Interface())...)
			return
		}*/
		if reflect.TypeOf(obj).Kind() == reflect.Struct {
			add(CastToStrings(reflect.ValueOf(obj).Interface())...)
			add("CtoSTr Struct")
		}
		if reflect.TypeOf(obj).Kind() == reflect.Ptr {
			//add(CastToStrings(reflect.ValueOf(obj).Interface())...)
			add("CtoSTr Poin")
		}
		switch val.(type) {
		case string:
			add(val.(string))
			/*case xml.Name:
			add(val.(xml.Name).Local)*/
		case DataType:
			add(colorman.AppendColor(string(val.(DataType)), colorman.COL_SETYPE))
		case byte:
			add(string(val.(byte)))
		case StringElement:
			add(val.(StringElement).XMLName.Local, val.(StringElement).Value)
		case int:
			add(strconv.Itoa(val.(int)))
		case lexer.LexemeSequence:
			for _, element := range val.(lexer.LexemeSequence) {
				add(CastToStrings(element)...)
			}
		case bool:
			if val.(bool) {
				add(colorman.AppendColor("true", colorman.COL_BOOL_TRUE))
			} else {
				add(colorman.AppendColor("false", colorman.COL_BOOL_FALSE))
			}
		case lexer.LiteralString:
			add((val.(lexer.LiteralString)).String())
		case lexer.Literal:
			add("Lit:" + strconv.Itoa((val.(lexer.Literal)).Index))
		case lexer.PLString:
			add((*(val.(lexer.PLString))).String())
		case lexer.Lexeme:
			add((*(val.(lexer.Lexeme)).Value).String())
		case []interface{}:
			for _, element := range val.([]interface{}) {
				add(CastToStrings(element)...)
			}
		case nil:
			add("Nil!")
		default:
			fmt.Println(val)
			add("CtoSTr DEF")
		}
	}
	return
}
