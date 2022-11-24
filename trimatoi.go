package piscine

func TrimAtoi(s string) int {
	result := 0
	var str string
	for i := 0; i < len(s); i++ {
		a := s[i]
		if byte(a) >= 48 && byte(a) <= 57 || byte(a) == 45 {
			str += string(a)
		}
	}
	for i := 0; i < len(s); i++ {
		a := s[i]
		if byte(a) >= 48 && byte(a) <= 57 {
			result = result*10 + int(a-'0')
		}
	}
	if len(str) != 0 && str[0] == '-' {
		result = -result
	}

	return result
}
