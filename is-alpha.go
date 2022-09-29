package exam

func IsAlpha(c rune) bool {
	return (c >= 65 && c <= 90) || (c >= 97 && c <= 122)
}
