package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 21 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	iterate := nb * RecursiveFactorial(nb-1)
	return iterate
}