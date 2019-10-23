package piscine

func TrimAtoi(s string) int {

	output := 0
	len := 0
	positive := false
	signed := false
	zero := false
	strrune := []rune(s)
	newstrrune := []rune("")
	for _, word := range strrune {
		for i := '1'; i <= '9'; i++ {
			if i == word {
				positive = true
				zero = true
				newstrrune = append(newstrrune, word)
				len++

			}
		}
		if word == '-' && positive != true {
			signed = true
		}
		if word == '0' && zero == true {
			newstrrune = append(newstrrune, word)
			len++
		}
	}
	if len == 0 {
		return 0
	}
	coef := 1
	for i := len - 1; i >= 0; i-- {

		output += coef * (int(newstrrune[i]) - 48)
		coef *= 10
	}
	if signed == true {
		output *= -1
	}
	return output
}
