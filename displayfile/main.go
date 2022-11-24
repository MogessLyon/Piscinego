package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("File name missing")
	}
	if len(arguments) > 1 {
		fmt.Println("Too many arguments")
	}
	if len(arguments) == 1 {
		file, _ := ioutil.ReadFile(arguments[0])
		str := string(file)
		fmt.Print(str)
	}
}
