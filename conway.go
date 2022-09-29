package exam

func Conway(index int) int {
	str := "1"
	for i := 1; i < index; i++ {
		str = ConwayStr(str)
	}
	return Atoi(str)
}

func ConwayStr(t string) string {
	fstr := ""
	tr := []rune(t)
	lastOne := tr[0]
	currentNb := 0
	for _, c := range tr {
		if lastOne == c {
			currentNb += 1
		} else {
			fstr += string(rune(currentNb+48)) + string(lastOne)
			currentNb = 1
		}
		lastOne = c
	}
	fstr += string(rune(currentNb+48)) + string(lastOne)
	return fstr
}

func Atoi(s string) int {
	negative := false
	total := 0
	for _, c := range s {
		if c >= 48 && c <= 57 {
			total = total*10 + int(c-48)
			continue
		} else if c == '-' || c == '+' && total == 0 {
			negative = c == '-'
			continue
		}
		return 0
	}
	return map[bool]int{true: total * -1, false: total}[negative]
}
