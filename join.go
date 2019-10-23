package piscine

func Join(strs []string, sep string) string {

	result := ""
	k := 0
	for i := range strs {
		k = i
	}

	for i := range strs {
		if i != k {
			strs[i] = strs[i] + sep
		}

	}
	for i := range strs {
		result += strs[i]
	}
	return result
}
