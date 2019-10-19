package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	var stroka1: = ''
	var stroka = []rune(str)
	for i := range stroka {
		stroka1 := stroka1 + stroka[i]
	}
	return string(stroka1)
}
