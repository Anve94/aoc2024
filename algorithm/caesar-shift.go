package algorithm

func CaesarShift(amount int, str string) string {
	var shiftedStr string
	for _, c := range str {
		ascii := int(c)
		var asciiRangeStartingLetter rune
		if c >= 'a' && c <= 'z' {
			asciiRangeStartingLetter = 'a'
		} else if c >= 'A' && c <= 'Z' {
			asciiRangeStartingLetter = 'A'
		} else {
			shiftedStr += string(c)
			continue
		}

		ascii -= int(asciiRangeStartingLetter)
		ascii += amount
		ascii %= 26

		if ascii < 0 {
			ascii += 26
		}

		ascii += int(asciiRangeStartingLetter)
		shiftedStr += string(rune(ascii))
	}
	return shiftedStr
}
