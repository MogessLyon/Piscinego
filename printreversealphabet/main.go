package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var str string = "abcdefghijklmnopqrstuvwxyz"
	for i := 25; i >= 0; i-- {
		z01.PrintRune(rune(str[i]))
	}
	z01.PrintRune('\n')
}
