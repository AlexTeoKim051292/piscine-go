package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	argument := os.Args
	//added perm
	strrune := []rune(argument[0])
	for _, word := range strrune {

		z01.PrintRune(word)
	}
	z01.PrintRune(10)
}
