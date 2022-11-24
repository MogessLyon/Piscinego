package piscine

func Join(strs []string, sep string) string {
	var str string
	for i, value := range strs {
		str += string(value)
		if i != len(strs)-1 {
			str += sep
		}
	}
	return str
}
