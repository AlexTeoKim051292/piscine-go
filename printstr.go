package piscine

func PrintStr(str string) int {
	var line := []rune(str)
	for i := range line {
		counter = i
	}
	return counter + 1
}
