package piscine

func IsAlpha(str string) bool {

	check := false
	strrune := []rune(str)
	word := 'a'
	for i := range strrune {
		word = strrune[i]
		if (word >= 'a' && word <= 'z') || (word >= 'A' && word <= 'Z') || (word >= '0' && word <= '9') {
			check = true
		} else {
			return false
		}
	}
	return check
}
