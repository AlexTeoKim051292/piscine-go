package piscine

func IsPrintable(str string) bool {

	check := false
	strrune := []rune(str)
	word := 'a'
	for i := range strrune {
		word = strrune[i]
		if word >= ' ' && word <= '~' {
			check = true
		} else {
			return false
		}
	}
	return check
}
