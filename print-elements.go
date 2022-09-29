package exam

func PrintElements(tab []int, str []string) string {
	fstr := ""
	for index, i := range tab {
		if i < len(str) {
			fstr += map[bool]string{true: str[i] + " ", false: str[i]}[index < len(tab)-1]
		}
	}
	return fstr
}
