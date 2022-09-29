package exam

import "fmt"

func Diamond(index int) {
	index -= 1
	for i := 0; i < index; i++ {
		l := make([]rune, index*2)
		for indexL := range l {
			if indexL == index-i || indexL == index+i {
				l[indexL] = '*'
			} else {
				l[indexL] = ' '
			}
		}
		fmt.Println(string(l))
	}
	l := make([]rune, index*2+1)
	for indexL := range l {
		if indexL == 0 || indexL == len(l)-1 {
			l[indexL] = '*'
		} else {
			l[indexL] = ' '
		}
	}
	fmt.Println(string(l))
	for i := index - 1; i >= 0; i-- {
		l := make([]rune, index*2)
		for indexL := range l {
			if indexL == index-i || indexL == index+i {
				l[indexL] = '*'
			} else {
				l[indexL] = ' '
			}
		}
		fmt.Println(string(l))
	}
}
