package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for n := '0'; n <= '9'; n++ {
					a := string(i)
					b := string(j)
					c := string(k)
					d := string(n)
					if a+b != c+d && a+b < c+d {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(n)
						if i != '9' || j != '8' || k != '9' || n != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}

				}
			}
		}
	}
	z01.PrintRune('\n')
}
