package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) int {
	var stroka = []rune(str)
	var stroka1 string = ""
	for i := range stroka {
		stroka1 = stroka + stroka[i]
	}
	return stroka1
}
