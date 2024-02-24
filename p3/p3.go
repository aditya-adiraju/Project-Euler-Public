package p3 

import (
  "math"
)

func Solve(num int) int {
	/*
		IDEA: Try everything backwards from sqrt(n) and return the first prime number that 
		divides num.
	*/

	for p := int(math.Sqrt(float64(num))); p >= 0; p-- {
		if num % p == 0 && isPrime(p) { 
			return p
		}
	}
	return 0 
}



// Util: Primality test
func isPrime(num int) bool {
	if num > 1 && num <= 3 {
		return true
	}
	// a fun optimization from the primality test wiki
	if num % 2 == 0 || num % 3 == 0 || num <= 1 { 
		return false
	}
	for i := 5; i*i < num; i += 6 {
		if num % i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}
