package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	var base10 int
	power := 1

	for i := len(nbr) - 1; i >= 0; i-- {
		a := int(nbr[i] - '0')
		base10 = base10 + a*power
		power *= len(baseFrom)
	}
	str := ""
	for base10 != 0 {
		str = string(baseTo[base10%len(baseTo)]) + str
		base10 /= len(baseTo)
	}

	return str
}
