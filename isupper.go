package piscine

func IsUpper(str string) bool {

	check := false
	strrune := []rune(str)
	word := 'a'
	for i := range strrune {
		word = strrune[i]
		if word >= 'A' && word <= 'Z' {
			check = true
		} else {
			return false
		}
	}
	return check
}
