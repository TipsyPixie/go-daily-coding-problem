package problem1

func Run(numbers []int, k int) bool {
	candidates := make(map[int]bool)
	for _, n := range numbers {
		if candidates[n] {
			return true
		}
		candidates[k-n] = true
	}
	return false
}
