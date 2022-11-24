package piscine

func ToUpper(s string) string {
	var str string
	for i := 0; i < len(s); i++ {
		st := s[i]

		if st >= 97 && st <= 122 {
			str += string(st - 32)
		} else {
			str += string(st)
		}
	}
	return str
}
