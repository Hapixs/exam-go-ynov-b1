package main

import (
	"fmt"
	"os"
)

func main() {
	for index, a := range os.Args[1:] {
		if ContainNoNumeric(a) {
			fmt.Printf("Argument %d : %s\n", index+1, a)
		}
	}
}

func ContainNoNumeric(str string) bool {
	for _, c := range str {
		if c >= 48 && c <= 57 {
			return false
		}
	}
	return true
}
