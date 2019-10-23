package piscine

func Capitalize(s string) string {

	firstletter := true
	strrune := []rune(s)
	for i, word := range strrune {
		if (word >= 'a' && word <= 'z') || (word >= 'A' && word <= 'Z') || (word >= '0' && word <= '9') {
			if firstletter == true {
				if word >= 'a' && word <= 'z' {
					strrune[i] = strrune[i] - 32

				}
				firstletter = false
			} else if word >= 'A' && word <= 'Z' {
				strrune[i] = strrune[i] + 32
			}

		} else {
			firstletter = true
		}

	}
	return string(strrune)
}
