package exam

func Average(a []int) int {
	total := 0
	st := 0
	for _, i := range a {
		if i >= 0 {
			total += i
			st++
		}
	}
	return total / st
}
