package piscine

func Split(s, sep string) []string {
	srune := []rune(s)
	seprune := []rune(sep)
	var str []string
	var st string

	idx := 0
	for idx < len(srune) {
		if srune[idx] == seprune[0] {
			find := true
			for j, k := idx, 0; k < len(sep); j, k = j+1, k+1 {
				if srune[j] != seprune[k] {
					find = false
					break
				}
			}

			if find {
				str = append(str, st)
				idx += len(sep)
				st = ""
			} else {
				st += string(srune[idx])
				idx += 1
			}
		} else {
			st += string(srune[idx])
			idx += 1
		}

		if idx == len(srune) {
			str = append(str, st)
		}
	}

	return str
}
