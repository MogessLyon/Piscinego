package piscine

func AlphaCount(s string) int {
	var count int
	for _, value := range s {
		a := rune(value)
		if a >= 65 && a <= 90 || a >= 97 && a <= 122 {
			count++
		}
	}
	return count
}
