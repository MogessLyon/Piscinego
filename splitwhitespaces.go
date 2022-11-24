package piscine

func SplitWhiteSpaces(s string) []string {
	st := []rune(s)
	var str []string
	var w string
	for idx, value := range st {
		if value == ' ' {
			if w != "" {
				str = append(str, w)
			}
			w = ""

		} else {
			w += string(value)
		}

		if idx == len(st)-1 {
			str = append(str, w)
		}
	}
	return str
}
