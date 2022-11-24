package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0]

	arr := []rune(args)
	for i, value := range arr {
		if i > 1 {
			z01.PrintRune(rune(value))
		}
	}
	z01.PrintRune('\n')
}
