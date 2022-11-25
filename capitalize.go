package piscine

func IsNumerique(s rune) bool {
	if !(s >= 48 && s <= 57) && !(s >= 65 && s <= 90) && !(s >= 97 && s <= 122) {
		return false
	}

	return true
}

func Capitalize(s string) string {
	str := []rune(s)
	first := true

	for idx, value := range str {
		if IsNumerique(value) && first {
			if value >= 'a' && value <= 'z' {
				str[idx] -= 32
			}
			first = false
		} else if value >= 'A' && value <= 'Z' {
			str[idx] += 32
		} else if !IsNumerique(value) {
			first = true
		}
	}
	return string(str)
}
