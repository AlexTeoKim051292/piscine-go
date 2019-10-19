package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {

	var stroka = []rune(str)
	for _, element := range stroka {
		z01.PrintRune(element)
	}
}
