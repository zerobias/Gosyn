package output

import (
	//"io"
	//"bufio"
	//"fmt"
	//"gosyn/debug"
	"github.com/zerobias/gosyn/models"
	. "github.com/zerobias/gosyn/output/colorman"
	//"os"
	"regexp"
	"strconv"
	"strings"
	/*"reflect"

	"unicode/utf8"*/)

const (
	minWidth = 5
	tabWidth = 2
	maxTabs  = 48 //TODO write correct out of range
)

type ParamString struct {
	Words  []string
	Indent int
	col    ColorValue
}

func (s *ParamString) AddWord(word string) {
	s.Words = append(s.Words, word)
}

type Paragraph struct {
	lines  []ParamString
	spaces []int
}

func CleanLen(word string) int {
	result := len(word)
	reg := regexp.MustCompile("\x1b[^m]*.")
	for _, val := range reg.FindAllString(word, -1) {
		result = result - len(val)
	}
	return result
}
func (p *Paragraph) realign() {
	p.spaces = make([]int, 0, maxTabs)
	for _, line := range p.lines {
		for i, word := range line.Words {
			line.fillSpaces(i, &p.spaces)
			if len(p.spaces) > i+line.Indent {
				if CleanLen(word)+tabWidth > p.spaces[i+line.Indent] {
					p.spaces[i+line.Indent] = CleanLen(word) + tabWidth
				} /*else if CleanLen(word)+tabWidth == p.spaces[i+line.Indent] {
					p.spaces[i+line.Indent]++
				}*/
			}
		}
	}
}

func (s *ParamString) fillSpaces(index int, spaces *[]int) {
	size := func() int {
		if index+s.Indent < maxTabs {
			return index + s.Indent
		} else {
			return maxTabs
		}
	}
	if len(*spaces) <= size() {
		for i := 0; i <= size()-len(*spaces); i++ {
			*spaces = append(*spaces, minWidth)
		}
	}
}

func (p *Paragraph) Send(level int, col ColorValue, vars ...interface{}) {
	str := models.CastToStrings(vars)
	p.lines = append(p.lines, ParamString{str, level, col})
}

func (p *Paragraph) SendPair(level int, col ColorValue, vars ...interface{}) {
	str := models.CastToStrings(vars)
	var result []string
	for i, word := range str {
		if i%2 == 0 {
			result = append(result, word)
		} else {
			result[len(result)-1] += " " + word
		}
	}
	p.lines = append(p.lines, ParamString{result, level, col})
}

func (s *ParamString) applyColor() {
	//s.Words[0] = string(s.col) + s.Words[0]
	for i := range s.Words {
		if IsColorized(s.Words[i]) {
			//strings.Replace(s.Words[i], string(RESET), string(s.col), 1)
			s.Words[i] = s.Words[i] + string(s.col)
		} /*else {
			s.Words[i] = string(s.col) + s.Words[i] // + string(RESET)
		}*/
	}
	s.Words[len(s.Words)-1] += string(RESET)
}

func (p *Paragraph) Init() {
	p.lines = *new([]ParamString)

}
func (p *Paragraph) Write() {
	defer p.close()
	//p.realign()
	for _, line := range p.lines {
		if len(line.Words) > 1 {
			//p.writeLine(strconv.Itoa(len(line.Words[1])))
		}
		line.applyColor()

	}
	p.realign()
	for _, line := range p.lines {
		//p.writeLine(AddColors(line.col, line.concat(p.spaces)))
		p.writeLine(line.concat(p.spaces))
	}
	var spaceStr string
	for _, num := range p.spaces {
		spaceStr += strconv.Itoa(num) + " "
	}
	p.writeLine(AddColors(COL_GREEN, spaceStr))
}

func (line *ParamString) concat(spaces []int) (result string) {
	addSpaces := func(str string, maxLength int) string {
		count := maxLength - CleanLen(str)
		return str + strings.Repeat(" ", count)
	}
	for i := 0; i < line.Indent; i++ {
		result += string(line.col) + strings.Repeat(" ", spaces[i])
	}
	for i, word := range line.Words {
		if len(spaces) > i+line.Indent { //TODO crash without this checking. Any better ways?
			//if CleanLen(word) < spaces[i+line.Indent] {
			result += addSpaces(word, spaces[i+line.Indent])
			/*} else {
				result += word
			}*/
		}
	}
	return
}

func (p *Paragraph) writeLine(line string) {
	out(line)
}

func (p *Paragraph) close() {
	//p.tabWriter.Flush()
	writer.Flush()
}
