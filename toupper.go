package piscine

func ToUpper(s string) string {

	strrune := []rune(s)
	for i, word := range strrune {
		if word >= 97 && word <= 122 {
			strrune[i] = word - 32
		}
	}
	return string(strrune)

}
