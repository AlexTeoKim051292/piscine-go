package piscine

func Join(strs []string, sep string) string {

	var result string
	k := 0
	for i, _ := range strs {
		k = i
	}

	for i, _ := range strs {
		if i != k {
			strs[i] = strs[i] + sep
		}

	}
	for i, _ := range strs {
		result += strs[i]
	}
	return result
}
