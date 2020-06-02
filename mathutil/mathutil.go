package mathutil

// MaxInt test
func MaxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// MinInt test
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
