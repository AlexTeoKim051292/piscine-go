package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	k := 0
	strrune := []rune(s)
	for i := range strrune {
		k = i
	}
	if n > k+1 {
		return 0
	}
	return strrune[n-1]
}
