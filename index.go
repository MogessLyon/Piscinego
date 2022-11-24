package piscine

func Index(s string, toFind string) int {
	l := len(toFind)
	index := -1
	if l == 0 {
		return 0
	}

	for i := 0; i < len(s); i++ {
		if string(s[i]) == string(toFind[0]) {
			a := true

			for j, k := i, 0; k < l; j, k = j+1, k+1 {
				if string(s[j]) != string(toFind[k]) {
					a = false
					index = -1
					break
				}
			}

			if a {
				index = i
				break
			}
		}
	}
	return index
}
