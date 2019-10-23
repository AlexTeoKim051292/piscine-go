package piscine

func BasicJoin(strs []string) string {

	result := ""
	for i := range strs {
		result += strs[i]
	}
	return result
}
