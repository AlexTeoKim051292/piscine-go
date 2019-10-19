package main

import "github.com/01-edu/z01"

func StrLen(str string) int {
	sum := 0
	for i := 0; i < len(str); i++ {
		sum := sum + 1
	}	
	return sum
}

func main() {
	str := "Hello World!"
	nb := z01.StrLen(str)
	z01.Println(nb)S
}
