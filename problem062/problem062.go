package problem062

func CountWay(rowCount int, columnCount int) int {
	return combination(rowCount+columnCount-2, rowCount-1)
}

func combination(n int, m int) int {
	dividendChannel := make(chan int, 1)
	divisorChannel := make(chan int, 1)
	go func() {
		a := 1
		for b := n; b > m; b-- {
			a *= b
		}
		dividendChannel <- a
	}()
	go func() {
		a := 1
		for b := m; b > 0; b-- {
			a *= b
		}
		divisorChannel <- a
	}()

	result := 1
	select {
	case dividend := <-dividendChannel:
		result *= dividend
	}
	select {
	case divisor := <-divisorChannel:
		result /= divisor
	}
	return result
}
