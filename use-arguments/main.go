package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for index, a := range args {
		if ContainNotNumeric(a) {
			fmt.Printf("Argument %d : %s\n", index+1, a)
		}
	}
}

func ContainNotNumeric(str string) bool {
	for _, c := range str {
		if c >= 48 && c <= 57 {
			return false
		}
	}
	return true
}
