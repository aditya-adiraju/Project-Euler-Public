package p6


// Naive approach, but the optimal one just uses the formula so that's boring
func Solve(num int) int {
	squareSum := 0
	polynomialSum := 0
	for i := 1; i <= num; i++ {
		squareSum += i*i
		polynomialSum += i
	}
	return (polynomialSum * polynomialSum) - squareSum
}
