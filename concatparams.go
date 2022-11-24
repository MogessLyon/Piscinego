package piscine

func ConcatParams(args []string) string {
	var str []rune

	for idx, value := range args {
		st := []rune(value)
		str = append(str, st...)
		if idx != len(args)-1 {
			str = append(str, '\n')
		}
	}
	return string(str)
}
