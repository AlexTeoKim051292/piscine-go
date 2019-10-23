package piscine

func ToLower(s string) string {

	strrune := []rune(s)
	for i, word := range strrune {
		if word >= 65 && word <= 90 {
			strrune[i] = word + 32
		}
	}
	return string(strrune)

}
