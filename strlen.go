package piscine

func StrLen(str string) int {
	var stroka = []rune(str)
	var counter = 0
	for i := range stroka {
		counter = i
	}
	return counter + 1
}
