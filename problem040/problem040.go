package problem040

func toBits(value int) []bool {
	bits := make([]bool, 32, 32)
	if value < 0 {
		for i := range bits {
			bits[i] = true
		}
	}
	for i := 0; (value >= 0 && (1<<i) <= value) || (value < 0 && (-1<<i) >= value); i++ {
		bits[i] = 1&(value>>i) > 0
	}
	return bits
}

func fromBits(bits []bool) int {
	value := 0
	for i, bit := range bits {
		if bit {
			value |= 1 << i
		}
	}
	return value
}

func FindSingleNumber(values []int) int {
	// suppose 32bit integer
	onesCount := make([]int, 32, 32)
	for _, value := range values {
		for i, bit := range toBits(value) {
			if bit {
				onesCount[i] += 1
			}
		}
	}

	singleNumberBits := make([]bool, 32, 32)
	for i, count := range onesCount {
		singleNumberBits[i] = count%3 == 1
	}
	return fromBits(singleNumberBits)
}
