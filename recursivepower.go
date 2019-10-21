package piscine

func RecursivePower(nb int, power int) int {

	result := 1

	if power < 0 {
		result = 0
	} else if power = 1 {
		result = nb
	} else if power > 1 {
		result = nb * RecursivePower(nb, power-1)
	}
	return result
}
