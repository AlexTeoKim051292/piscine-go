package piscine

func IterativePower(nb int, power int) int {

	result := 1

	if power < 0 {
		result = 0
	} else {
		for i := 1; i < power+1; i++ {
			result = result * nb
		}
	}
	return result
}
