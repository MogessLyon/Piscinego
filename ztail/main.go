package main

import (
	"fmt"
	"os"
)

func ConvertToInt(s string) int {
	result := 0

	for _, value := range s {
		result = result*10 + int(value-'0')
	}

	return result
}

func main() {
	arguments := os.Args[2:]
	if len(arguments[1:]) == 0 {
		return
	}

	ln := ConvertToInt(arguments[0])

	for idx, value := range arguments[1:] {
		file, err := os.ReadFile(value)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		affichage := string(file[len(file)-ln-1:])

		if len(arguments[1:]) == 1 {
			fmt.Printf("%s", affichage)
		} else {
			fmt.Printf("===> %s <===\n%s", value, affichage)
		}

		if len(arguments[1:])-1 != idx {
			fmt.Printf("\n")
		}

	}
}
