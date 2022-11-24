package piscine

func BasicJoin(elems []string) string {
	var str string
	for _, value := range elems {
		str += string(value)
	}
	return str
}
