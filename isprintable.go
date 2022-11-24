package piscine

func IsPrintable(s string) bool {
	if s == "" {
		return true
	}
	for _, value := range s {
		if byte(value) < 32 || byte(value) > 127 {
			return false
		}
	}
	return true
}
