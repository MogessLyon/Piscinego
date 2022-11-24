package piscine

func Swap(a *int, b *int) {
	valua := *a
	valub := *b

	*a = valub
	*b = valua
}
