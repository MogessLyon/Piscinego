package main

import (
	"os"
)

func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, value := range s {
		if value < '+' || value > '+' && value < '-' || value > '-' && value < '0' || value > '9' {
			return false
		}
	}
	return true
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 3 {
		return
	}
	if !IsNumeric(arguments[0]) || !IsNumeric(arguments[2]) {
		return
	}

	firstop := Atoi(arguments[0])
	secondop := Atoi(arguments[2])
	b := arguments[1]
	result := 0
	limit := 9223372036854775807

	if firstop >= limit || secondop > limit || firstop <= -limit || secondop < -limit {
		return
	}

	switch b {
	case "+":
		result = firstop + secondop
	case "-":
		result = firstop - secondop
	case "*":
		result = firstop * secondop
	case "/":
		if secondop == 0 {
			str := []byte("No division by 0\n")
			os.Stdout.Write(str)
			return
		}
		result = firstop / secondop
	case "%":
		if secondop == 0 {
			str := []byte("No modulo by 0\n")
			os.Stdout.Write(str)
			return
		}
		result = firstop % secondop
	default:
		return
	}

	os.Stdout.Write(Itoi(result))
}

func Atoi(s string) int {
	result := 0

	for _, value := range s {
		if value >= '0' && value <= 57 {
			result = result*10 + int(value-'0')
		}
	}
	if s[0] == '-' {
		result = -result
	}
	str := ""
	for _, value := range s {
		if value >= '0' && value <= 57 || value == '+' || value == '-' {
			if value == '+' || value == '-' {
				str += string(value)
			}

			if len(str) >= 2 {
				return 0
			}

			if len(str) == 1 {
				if s[0] != str[0] {
					return 0
				}
			}
		} else {
			return 0
		}
	}
	return result
}

func Itoi(nbr int) []byte {
	var arr []byte
	var result []byte

	if nbr < 0 {
		result = append(result, '-')
		nbr = -nbr
	} else if nbr == 0 {
		result = append(result, '0')
	}

	for nbr != 0 {
		k := byte(nbr%10 + '0')
		nbr /= 10
		arr = append(arr, k)
	}
	for v := len(arr) - 1; v >= 0; v-- {
		result = append(result, arr[v])
	}
	result = append(result, '\n')
	return result
}
