package piscine

func Fibonacci(index int) int {

	result := 1

	if index < 0 {
		result = -1
	} else if index <= 1 {
		result = index
	} else {
		result = Fibonacci(index-1) + Fibonacci(index-2)
	}
	return result
}
