package colorman

import (
	//"io"
	//"bufio"
	//"gosyn/debug"
	//"gosyn/models"
	"strconv"
	"strings"
)

const (
	BLACK ColorType = iota
	RED
	GREEN
	YELLOW
	BLUE
	PURPLE
	CYAN
	WHITE
	_
	DEFAULT
)
const (
	REGULAR FontType = iota
	UNDERLINE
	BACKGROUND
	BRIGHT_BACKGROUND
)

var (
	RESET       ColorValue = "\x1b[0m"
	COL_GREY    ColorValue = ColorGetter(BLACK, REGULAR)
	COL_GREEN              = ColorGetter(GREEN, REGULAR)
	COL_RED                = ColorGetter(RED, REGULAR)
	COL_RED_BB             = ColorGetter(RED, BRIGHT_BACKGROUND)
	COL_BLUE_B             = ColorGetter(BLUE, BACKGROUND)
	COL_BLUE_U             = ColorGetter(BLUE, UNDERLINE)
	COL_GREEN_B            = ColorGetter(GREEN, BACKGROUND)
	COL_CYAN               = ColorGetter(CYAN, REGULAR)
)

type ColorValue string
type ColorType int
type FontType int
type ColorStruct struct {
	color ColorType
	font  FontType
}

type colorGenerator func(c ColorType) []int

var colorMap = map[FontType]colorGenerator{
	REGULAR:           func(c ColorType) []int { return []int{1, 30 + int(c)} },
	UNDERLINE:         func(c ColorType) []int { return []int{4, 30 + int(c)} },
	BACKGROUND:        func(c ColorType) []int { return []int{40 + int(c)} },
	BRIGHT_BACKGROUND: func(c ColorType) []int { return []int{0, 100 + int(c)} },
}

func (s *ColorStruct) GetValue() ColorValue {
	return ColorGetter(s.color, s.font)
}

func ColorGetter(col ColorType, font FontType) ColorValue {
	resultBuilder := func(parts []int) (result string) {
		values := make([]string, 0, 2)
		for _, n := range parts {
			values = append(values, strconv.Itoa(n))
		}
		return escapeSymbol + strings.Join(values, ";") + escapeEnd
	}
	return ColorValue(resultBuilder(colorMap[font](col)))
}

const (
	escapeSymbol = "\x1b["
	escapeEnd    = "m"
)

func IsColorized(word string) bool {
	return strings.HasPrefix(word, escapeSymbol)
}

/*const (
	GRAY   Color = "\x1b[1;30m"
	RED    Color = "\x1b[31;1m"
	GREEN  Color = "\x1b[1;32m"
	YELLOW Color = "\x1b[33m"
	//BLUE   Color = "\x1b[34;1m"
	BLUE       Color = "\x1b[30m"
	DEF        Color = "\x1b[39;0m"
	RED_BACK   Color = "\x1b[41;0m"
	WHITE_BACK Color = "\x1b[47m"
	DEF_BACK   Color = "\x1b[49m"
	RESET            = "\x1b[0m" //DEF + DEF_BACK
)*/

var (
	COL_BOOL_TRUE  ColorValue = RESET + ColorGetter(GREEN, REGULAR)
	COL_BOOL_FALSE            = RESET + ColorGetter(RED, REGULAR)
	COL_SETYPE                = ColorGetter(BLUE, BRIGHT_BACKGROUND) + ColorGetter(WHITE, REGULAR)
	COL_DEEP_LIM              = ColorGetter(BLACK, BRIGHT_BACKGROUND) + ColorGetter(BLACK, UNDERLINE)
)

func AppendColor(str string, col ColorValue) string {
	return string(col) + str // + string(RESET)
}

func AddColors(col ColorValue, str string) string {
	return string(col) + str //+ string(RESET)
}
