package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	if n < 0 {

		return
	}
	if n == 0 {

		z01.PrintRune('0')
		return
	}
	strrune := []rune("")
	for i := n - 1; i >= 0; i-- {
		i++
		strrune = append(strrune, rune(i%10)+48)
		i = i / 10
	}
	for index := range strrune {
		for j := range strrune {
			if strrune[index] < strrune[j] {
				strrune[index], strrune[j] = strrune[j], strrune[index]
			}
		}
	}

	for _, word := range strrune {

		z01.PrintRune(word)
	}

}
