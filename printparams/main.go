package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	argument := os.Args
	for i := range argument {
		if i != 0 {
			strrune := []rune(argument[i])
			for _, word := range strrune {
				z01.PrintRune(word)
			}
			z01.PrintRune(10)
		}
	}
	//z01.PrintRune(10)
}
