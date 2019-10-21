package piscine

func IterativeFactorial(nb int) int {

	result := 1

	if nb == 0 {
		result = 1
	} else if nb < 0 {
		result = 0
	} else if nb >= 15 {
		result = 0
	} else if nb > 0 {
		for i := 1; i < nb+1; i++ {
			result = result * i
		}
	}
	return result
}
