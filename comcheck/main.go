package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	check := []string{"galaxy", "galaxy 01", "01"}

	checkin := false
	for _, value := range arguments {
		for _, val := range check {
			if val == value {
				checkin = true
				break
			}
		}
	}

	if checkin {
		fmt.Println("Alert!!!")
	}
}
