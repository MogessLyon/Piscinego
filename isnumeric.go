package piscine

func IsNumeric(s string) bool {
	if s == "" {
		return true
	}
	for _, value := range s {
		if byte(value) < 48 || byte(value) > 57 {
			return false
		}
	}
	return true
}
