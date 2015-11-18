package debug

type EnumType int8

const (
	STATE EnumType = iota
	SYMB
	CATEGORY
)

var enumText = map[EnumType]map[int]string{
	STATE: {0: "S", 1: "Q01", 2: "Q02", 3: "Q03", 4: "Q04", 5: "Q05",
		6: "Q06", 7: "Q07", 8: "Q08", 9: "Q09", 10: "Q10",
		11: "Q11", 12: "Q12", 13: "Q13"},
	SYMB:     {0: "DIGIT", 1: "LETTER", 2: "UNDERSCORE", 3: "COMMA", 4: "LEXEME_END", 5: "ANOTHER"},
	CATEGORY: {0: "KEY", 1: "SIGN", 2: "DELIM", 3: "CLASS", 4: "NUM"},
}

func GetEnumText(enumType EnumType, id int) string {
	return enumText[enumType][id]
}
