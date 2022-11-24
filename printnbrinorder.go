package piscine

import (
	"github.com/01-edu/z01"
)

func shortInteger(arr []int) {
	i := 1
	for i < len(arr) {
		if arr[i-1] > arr[i] {
			tmp1 := arr[i]
			tmp2 := arr[i-1]
			arr[i] = tmp2
			arr[i-1] = tmp1
			i = 1
		} else {
			i++
		}
	}
}

func PrintNbrInOrder(n int) {
	var arr []int
	if n == 0 {
		arr = append(arr, 0)
	}
	for i := 0; n > 0; i++ {
		arr = append(arr, n%10)
		n /= 10
	}
	shortInteger(arr)

	for i := 0; i < len(arr); i++ {
		z01.PrintRune(rune('0' + arr[i]))
	}
}
