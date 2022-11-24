package piscine

func IsUpper(s string) bool {
	for _, value := range s {
		if byte(value) < 65 || byte(value) > 90 {
			return false
		}
	}
	return true
}
