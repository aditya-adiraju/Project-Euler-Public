package p2


func Practice() []int {
	result := make([]int, 10)
	for i := range result {
		switch {
		case i == 0 || i == 1:
			result[i] = i + 1
		default:
			result[i] = result[i-1] + result[i-2]
		}
	}
	return result 
}

func Solve(n int) (sum int) {
	// invariant: Fib(curr - 2) 
	prev2 := 1
	// invariant: Fib(curr - 1)
	prev1 := 2

	sum = 0

	if prev1 % 2 == 0 { 
		sum += prev1
	}

	if prev2 % 2 == 0 { 
		sum += prev2
	}

	for prev1 + prev2 <= n {
		curr := prev1 + prev2
		if curr % 2 == 0 {
			sum += curr
		}
		prev2, prev1 = prev1, curr
	}

	return sum
}

