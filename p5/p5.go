package p5

func Solve(start, end int) (result int) {
	result = 1
	for i := start; i <= end; i++ {
		result = lcm(result, i)
	}

	return result 	
}

// lcm returns the least common multiple of a, b
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
} 


// gcd returns the greatest commond divisor between a, b
func gcd(a, b int) int {
	// Extended Euclidean Algorithm
	switch {
	case a == 0:
		return b
	case b == 0:
		return a
	default:
		return gcd(b, a % b)
	}
}
