package piscine

func IsPrime(nb int) bool {
	a := 2
	check := true
	if nb <= 1 {
		check = false
	}
	for a < nb-1 {
		if nb%a == 0 {
			check = false
			break
		} else {
			a++
		}
	}
	return check
}
