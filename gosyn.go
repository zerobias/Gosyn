package main

import (
	"fmt"
	"github.com/zerobias/gosyn/controller"
	//"gosyn/models"
	"github.com/zerobias/gosyn/lexer"
	"github.com/zerobias/gosyn/output"
	"github.com/zerobias/gosyn/output/htmlreport"
	"github.com/zerobias/gosyn/utils"
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

	classStep := controller.ClassStep(controller.NewStep("SIGN"))
	step := controller.Step(&classStep)
	facade := controller.NewFacade(&step, nil, nil)
	fmt.Println("Test ", lexer.GetClassByName(facade.StepValue()))

	ruleStep := controller.RuleStep(controller.NewStep("program"))
	step = controller.Step(&ruleStep)
	facade = controller.NewFacade(&step, nil, nil)
	fmt.Println(facade.StepType(), facade.StepValue())
	b, ch := facade.HasChilds()
	fmt.Println("Test ", b, *ch)
}
