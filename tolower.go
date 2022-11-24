package piscine

func ToLower(s string) string {
	var str string
	for i := 0; i < len(s); i++ {
		st := s[i]

		if st >= 65 && st <= 90 {
			str += string(st + 32)
		} else {
			str += string(st)
		}
	}
	return str
}
