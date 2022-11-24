package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	convertRune(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	convertRune(points.y)
	z01.PrintRune('\n')
}

func convertRune(nb int) {
	k := '0'
	for i := 0; i < nb/10; i++ {
		k++
	}
	z01.PrintRune(k)

	r := '0'
	for i := 0; i < nb%10; i++ {
		r++
	}
	z01.PrintRune(r)
}
