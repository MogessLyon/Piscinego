package piscine

func DivMod(a int, b int, div *int, mod *int) {
	divide := a / b
	reminder := a % b

	*div = divide
	*mod = reminder
}
