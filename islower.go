package piscine

func IsLower(s string) bool {
	for _, value := range s {
		if byte(value) < 97 || byte(value) > 122 {
			return false
		}
	}
	return true
}
