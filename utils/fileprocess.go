package utils

import (
	"fmt"
	//"io"
	"io/ioutil"
	//"log"
	//"bytes"
	"github.com/zerobias/gosyn/lexer"
	"github.com/zerobias/gosyn/models"
	"os"
	//"strings"
	//"unicode/utf8"
	"bufio"
	"encoding/xml"
	"github.com/zerobias/gosyn/output"
	. "github.com/zerobias/gosyn/output/colorman"
)

const space = " " //TODO delete this
const TermsFilePath = "C:/Users/LapTop/Documents/LiquidXMLGolangLLSP"
const TermsFileName = "lang.xml" //TODO remove these hardlinks

func LoadfileDefault(filename string) string {
	appFolder, _ := os.Getwd()
	return string(Loadfile(appFolder, filename))
}

func Loadfile(filepath, filename string) []byte {
	fmt.Printf("Load file %s path %s\n", filename, filepath)
	file, err := ioutil.ReadFile(filepath + "\\" + filename)
	if err != nil {
		fmt.Printf("file read error: %s\n", err.Error())
		return []byte("")
	}
	return file
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
func WriteFileDefault(filename, data string) { //TODO remove code duplication
	appFolder, _ := os.Getwd()
	Writefile(appFolder, filename, data)
}
func Writefile(filepath, filename, data string) {
	//d1 := []byte(data)
	//	err := ioutil.WriteFile(filepath+"/"+filename, d1, 0664)
	outfile, err := os.Create(filepath + "/" + filename)
	Check(err)
	defer outfile.Close()
	writer := bufio.NewWriter(outfile)
	_, err = writer.WriteString(data)
	Check(err)
	defer writer.Flush()
}

func ParseXML(rawData []byte) (tokens lexer.TokenList, rules []models.Rule) {
	var xmlFile models.XMLRoot
	parseError := xml.Unmarshal(rawData, &xmlFile)
	if parseError != nil {
		output.ColorPrint(COL_RED_BB, 2, parseError.Error())
	}
	for _, term := range xmlFile.Terms { //TODO rename "terms"
		tokens = append(tokens, *lexer.NewToken(term.Name, lexer.GetClassByName(term.Class), term.Value))
	}
	rules = xmlFile.Rules
	return
}
