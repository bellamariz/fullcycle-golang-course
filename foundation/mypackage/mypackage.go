package mypackage

func MySum[T int | float64](a, b T) T {
	return a + b
}
