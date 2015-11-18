package controller

import (
	//"container/list"
	"fmt"
	"strconv"
)

func printError(errorText string) {
	const textPrefix = "==========ERROR==========\n"
	fmt.Println(textPrefix, errorText)
}

func SyntaxError(lastIndex int, buffer []rune) {
	const textPrefix = "SYNTAX ERROR\n"
	printError(textPrefix + "ind " + strconv.Itoa(lastIndex) + " buf " + string(buffer))
}
