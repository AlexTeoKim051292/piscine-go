package piscine

func IsNumeric(str string) bool {

	check := false
	strrune := []rune(str)
	word := 'a'
	for i := range strrune {
		word = strrune[i]
		if word >= '0' && word <= '9' {
			check = true
		} else {
			return false
		}
	}
	return check
}
