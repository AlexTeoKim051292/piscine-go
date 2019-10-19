package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) int {
	var stroka = []rune(str)
	for i := range stroka {
		z01.PrintRune(stroka)
	}
}
