package problem061

func PowSimple(base int, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func PowFaster(base int, exponent int) int {
	result := 1
	modifiedExponent, modifiedBase := exponent, base
	for modifiedExponent > 1 {
		if modifiedExponent&1 == 1 {
			result *= modifiedBase
		}
		modifiedExponent >>= 1
		modifiedBase *= modifiedBase
	}
	return result * modifiedBase
}
