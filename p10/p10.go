package p10

// WAYYY BETTER THAN P7
// returns the nth prime
func Solve(n int) int {

	limit := n
	// initialize 
	primes := make([]int, limit)
	for i := range primes {
		primes[i] = i
	}

	for i, value := range primes {
		if value > 1 {
			primes = filterMultiples(primes, value)
		} else {
			primes[i] = 0
		}
	}

	return sum(primes)
}

func filterMultiples(nums []int, multiple int) []int {
	for i := 2 * multiple; i < len(nums); i += multiple {
		nums[i] = 0
	}
	return nums
}


func sum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum 
}
