package piscine

func ConcatParams(args []string) string {

	result := ""
	k := 0
	for i := range args {
		k = i
	}
	k++
	for i := range args {
		if i != k-1 {
			result += args[i] + "\n"
		} else {
			result += args[i]
		}
	}

	return result

}
