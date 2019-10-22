package piscine

func AlphaCount(str string) int {

	sentence := []rune(str)
	counter := 0

	for _, letter := range sentence {

		if letter >= 65 && letter <= 90 || letter >= 97 && letter <= 122 {
			counter++
		}
	}
	return counter
}
