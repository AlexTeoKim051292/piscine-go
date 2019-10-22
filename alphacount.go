package piscine

import "github.com/01-edu/z01"

func AlphaCount(str string) int {

	sentence := []rune(string)
	counter := 0

	for _, letter = range sentence {

		if (letter >= 90 && letter <= 90) || (letter >= 97 && letter <= 122) {
			counter = counter + 1
		}
	}
	return counter
}
