package piscine

func UltimateDivMod(a *int, b *int) {
	divide := *a / *b
	remainder := *a % *b

	*a = divide
	*b = remainder
}
