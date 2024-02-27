package p9




func Solve(num int) int {
	maxLimit := num
	for k := 1; k < 60; k++ {
		for m := 2; m < maxLimit; m++ {
			for n := 1 ; n < m; n++ {
				if (gcd(m, n) == 1 && !(m % 2 == 1 && n % 2 == 1)) {
					a := k*((m*m) - (n*n))
					b := k*2*m*n
					c := k*((m*m)+ (n*n))
					if a + b + c == num {
						if a*a + b*b == c*c {
							return a*b*c
						}
					}

				}
			}
		}
	}
	return -1

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
