package piscine

func IsNumerique(s rune) bool {
	if !(s >= 48 && s <= 57) && !(s >= 65 && s <= 90) && !(s >= 97 && s <= 122) {
		return false
	}

	return true
}

func Capitalize(s string) string {
	st := []rune(s)
	firstletter := true

	for i := range st {
		if (st[i] >= 'a' && st[i] <= 'z') && firstletter {
			st[i] -= 32
			firstletter = false
		} else if (st[i] >= 48 && st[i] <= 57) || (st[i] >= 'A' && st[i] <= 'Z') && firstletter {
			firstletter = false
		} else if (st[i] >= 'A' && st[i] <= 'Z') && !firstletter {
			st[i] += 32
		} else if !firstletter && !(st[i] >= 48 && st[i] <= 57) && !IsNumerique(st[i]) {
			firstletter = true
		}
	}
	str := string(st)

	return str
}
