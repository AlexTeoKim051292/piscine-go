package piscine

func IsLower(str string) bool {

	check := false
	strrune := []rune(str)
	word := 'a'
	for i := range strrune {
		word = strrune[i]
		if word >= 'a' && word <= 'z' {
			check = true
		} else {
			return false
		}
	}
	return check
}
