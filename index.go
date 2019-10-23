package piscine

func Index(s string, toFind string) int {

	strrune := []rune(s)
	findrune := []rune(toFind)
	len1 := 0
	len2 := 0
	check := false
	for i := range strrune {

		len1 = i
	}
	for i := range findrune {
		len2 = i
		check = true
	}
	len1++
	len2++
	if check == false {
		return 0
	}
	for i := 0; i <= len1-len2; i++ {
		if toFind == s[i:len2+i] {

			return i
		}
	}

	return -1

}
