package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		read := io.TeeReader(os.Stdin, os.Stdout)
		ioutil.ReadAll(read)
		os.Stdin.Close()
		os.Stdout.Close()
	} else {
		for _, value := range arguments {
			PrintStr(value)
		}
	}
}

func PrintStr(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		for _, value := range "ERROR: " + err.Error() {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
		os.Exit(1)
	}

	for _, value := range file {
		z01.PrintRune(rune(value))
	}
}
