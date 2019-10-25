package piscine

func SplitWhiteSpaces(str string) []string {

	result := [100]string{}
	strword := ""
	k := 0
	c := 0
	lengthword := 0
	for i := range str {
		k = i
	}
	k++

	for i, word := range str {
		if i == k-1 && string(word) != " " && string(word) != "\n" && string(word) != "\t" {
			strword += string(word)

			result[c] = strword
			c++
			lengthword++
			//result = append(result, strword)
		} else if string(word) != " " && string(word) != "\n" && string(word) != "\t" {

			strword += string(word)
			lengthword++
		} else if lengthword > 0 && (string(word) == " " || string(word) == "\n" || string(word) == "\t") {

			result[c] = strword
			c++
			lengthword = 0
			//result = append(result, strword)
			strword = ""
		}

	}
	result1 := make([]string, c)
	for i := range result1 {
		result1[i] = result[i]

	}
	return result1
}
