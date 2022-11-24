package piscine

func Rot14(s string) string {
	word := []rune(s)

	for idx, value := range word {
		if value >= 'a' && value <= 'z' {
			if value == 'm' || value == 'n' {
				word[idx] -= 12
			} else if value < 'm' {
				word[idx] += 14
			} else {
				word[idx] -= 12
			}
		}

		if value >= 'A' && value <= 'Z' {
			if value == 'M' || value == 'N' {
				word[idx] -= 12
			} else if value < 'M' {
				word[idx] += 14
			} else {
				word[idx] -= 12
			}
		}
	}

	return string(word)
}
