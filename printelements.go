package exam

func PrintElements(tab []int, str []string) string {
	fstr := ""
	for index, i := range tab {
		if i < len(str) {
			fstr += str[i]
			if index < len(tab)-1 {
				fstr += " "
			}
		}
	}
	return fstr
}
