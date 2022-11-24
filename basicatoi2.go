package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, value := range s {
		if byte(value) >= 48 && byte(value) <= 57 {
			result = result*10 + int(value-'0')
		}
	}

	for _, value := range s {
		if byte(value) < 48 || byte(value) > 57 {
			result = 0
		}
	}
	return result
}
