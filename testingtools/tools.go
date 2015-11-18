package testingtools

import (
	"testing"
)

func PrintExpected(t *testing.T, exp, got interface{}, send ...interface{}) { // (string, []interface{}) {
	const printformat = "\nexp :\t%v\ngot :\t%v\nsend:\t%v"
	t.Errorf(printformat, exp, got, send)
}
