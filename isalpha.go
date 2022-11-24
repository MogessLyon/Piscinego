package piscine

func IsAlpha(s string) bool {
	if s == "" {
		return true
	}
	for _, value := range s {
		if byte(value) < 48 || (byte(value) > 57 && byte(value) < 65) || (byte(value) > 90 && byte(value) < 97) || byte(value) > 122 {
			return false
		}
	}
	return true
}
