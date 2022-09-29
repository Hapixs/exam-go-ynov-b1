package exam

func Decipher(s string) string {
	str := ""
	for _, r := range s {
		realValue := r
		if IsAlpha(r) {
			realValue -= 16
			if r >= 'A' && realValue < 'A' {
				realValue = 'Z' - ('A' - (realValue + 1))
			} else if r >= 'a' && realValue < 'a' {
				realValue = 'z' - ('a' - (realValue + 1))
			}
		}
		str += string(rune(realValue))
	}
	return str
}
