package exam

func RecursiveSequence(current int, max int) int {
	if current <= max {
		next := map[bool]int{true: current + 5, false: current*3 + 1}[current%2 == 0]
		if next > max {
			return current
		}
		return RecursiveSequence(next, max)
	}
	return 0
}
