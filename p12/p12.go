package p12


// returns the first triangle number with more than thresholdCount divisors
func Solve(thresholdCount int) int {
	for i := 1;; i++ {
		if countDivisors(getTriangleNumber(i)) > thresholdCount {
			return getTriangleNumber(i)
		}
	} 
}

// returns the nth triangle number
func getTriangleNumber(n int) int {
	return (n * (n + 1))/ 2
}

// returns the number of divisors n has 
func countDivisors(n int) int {
	count := 0
	for i := 1; i*i <= n; i++ {
		if n % i == 0 {
			if i * i == n {
				count++
			} else {
				count += 2
			}
		}
	}

	return count
}
