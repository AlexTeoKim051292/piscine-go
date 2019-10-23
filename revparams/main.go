package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	argument := os.Args
	k := 0
	for i := range argument {
		k = i
	}
	for i := k; i >= 0; i-- {
		if i != 0 {
			strrune := []rune(argument[i])
			for j := range strrune {
				z01.PrintRune(strrune[j])
			}
			z01.PrintRune(10)
		}

	}
}
