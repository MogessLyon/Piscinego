package piscine

func Atoi(s string) int {
	result := 0
	for _, value := range s {
		if byte(value) >= 48 && byte(value) <= 57 {
			result = result*10 + int(value-'0')
		}
	}
	var str string
	if len(s) > 0 && s[0] == '-' {
		result = -result
	}

	for _, value := range s {
		if byte(value) >= 48 && byte(value) <= 57 || byte(value) == 43 || byte(value) == 45 {

			if value == '+' || value == '-' {
				str += string(value)
			}

			if len(str) >= 2 {
				result = 0
			}
			if len(str) == 1 {
				if s[0] != str[0] {
					result = 0
				}
			}

		} else {
			result = 0
		}
	}

	return result
}
