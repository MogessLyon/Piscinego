package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args

	for i := len(arr) - 1; i >= 1; i-- {
		a := arr[i]

		for j := 0; j < len(a); j++ {
			z01.PrintRune(rune(a[j]))
		}
		z01.PrintRune('\n')
	}
}
