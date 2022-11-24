package piscine

import (
	"github.com/01-edu/z01"
)

func IsNegative(arg int) {
	var str string = "TF"
	if arg < 0 {
		z01.PrintRune(rune(str[0]))
		z01.PrintRune('\n')
	} else {
		z01.PrintRune(rune(str[1]))
		z01.PrintRune('\n')
	}
}
