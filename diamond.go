package exam

import "fmt"

func Diamond(index int) {
	for i := 0; i <= index*2; i++ {
		if i < index {
			fmt.Println(string(BuildLine(index, index-i, index+i)))
		} else if i > index {
			fmt.Println(string(BuildLine(index, i-index, index+(index*2-i))))
		} else {
			fmt.Println(string(BuildLine(index, 0, index*2)))
		}
	}
}

func BuildLine(size, index1, index2 int) []rune {
	l := make([]rune, size*2+1)
	for indexL := range l {
		l[indexL] = map[bool]rune{true: '*', false: ' '}[indexL == index1 || indexL == index2]
	}
	return l
}
