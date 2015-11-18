package main

import (
	"fmt"
	"gosyn/controller"
	//"gosyn/models"
	"gosyn/lexer"
	"gosyn/output"
	"gosyn/output/htmlreport"
	"gosyn/utils"
)

func main() {
	output.InitWriter()
	fmt.Println("start")
	lexemeTypes, syntaxRules := utils.ParseXML(utils.Loadfile(utils.TermsFilePath, utils.TermsFileName)) // utils.ParseJSONLexemes("lexemes.json")
	sourceFile := "code1.txt"
	fileText := utils.LoadfileDefault(sourceFile)
	lexer.SetTokens(lexemeTypes)
	text := lexer.HandleSource(fileText)
	seq := lexer.ScanCycle(text)
	output.PrintString(0, []lexer.Lexeme(*seq))
	tree := controller.SyntaxCycle(seq, syntaxRules)
	output.CloseWriter()
	//fmt.Println(tree.ChildList[0].Type)
	htmlreport.MakeReport(tree)

}
