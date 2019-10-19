package piscine

func StrLen(str string) int {
	sum := 0
	for i := 0; i < len(str); i++ {
		sum := sum + 1
	}
	return sum
}
