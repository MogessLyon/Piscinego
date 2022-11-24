package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	arr1 := "01234567"
	arr2 := "12345678"
	arr3 := "23456789"

	for i := 0; i < 8; i++ {
		for j := i; j < 8; j++ {
			for k := j; k < 8; k++ {

				z01.PrintRune(rune(arr1[i]))
				z01.PrintRune(rune(arr2[j]))
				z01.PrintRune(rune(arr3[k]))

				if i+k+j != 21 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}

			}
		}
	}
	z01.PrintRune('\n')
}
