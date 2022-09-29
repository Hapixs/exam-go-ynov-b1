package exam

func RecursiveSequence(current int, max int) int {
	if current > max {
		return 0
	}
	next := map[bool]int{true: current + 5, false: current*3 + 1}[isEven(current)]
	if next > max {
		return current
	}
	return RecursiveSequence(next, max)
}

func isEven(i int) bool {
	return i%2 == 0
}
