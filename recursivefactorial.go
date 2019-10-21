package piscine

func RecursiveFactorial(nb int) int {

	result := 1

	if nb == 0 || nb == 1 {
		result = 1
	} else if nb < 0 {
		result = 0
	} else if nb >= 15 {
		result = 0
	} else if nb > 1 {
		result = result * RecursiveFactorial(nb-1)
	}
	return result
}

