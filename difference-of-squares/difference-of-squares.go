package diffsquares

func SquareOfSums(a int) int {
	i := 0
	for j := 0; j <= a; j++ {
		i += j
	}
	return i * i
}

func SumOfSquares(a int) int {
	i := 0
	for j := 0; j <= a; j++ {
		i += j * j
	}
	return i
}

func Difference(a int) int {
	return SquareOfSums(a) - SumOfSquares(a)
}
