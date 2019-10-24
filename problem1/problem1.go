package problem1

func Run(numbers []uint, k uint) bool {
    candidates := make(map[uint]bool)
    for _, n := range numbers {
        if n < k {
            if candidates[n] {
                return true
            }
            candidates[k-n] = true
        }
    }
    return false
}
