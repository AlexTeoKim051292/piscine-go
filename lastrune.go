package piscine

func LastRune(s string) rune {

	k := 0
	strrune := []rune(s)
	for i := range strrune {
		k = i
	}
	return strrune[k]
}
