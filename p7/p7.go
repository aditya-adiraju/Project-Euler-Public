package p7


// The Sieve of Erasthenes seems like the most practical solution for normal computer stuff

// returns the nth prime
func Solve(n int) int {

	// A guesstimate for the possible limti, ofc I'm too lazy to implement the true upper bound
	limit := 200000  
	// initialize
	nums := make([]int, limit - 2)
	for i := 0; i < limit - 2; i++ {
		nums[i] = i + 2
	}


	for i := 0; i < len(nums); i++ {
		nums = filterMultiples(nums, nums[i])
	}

	return nums[n - 1]

}

func filterMultiples(nums []int, multiple int) []int {
	filteredNums := make([]int, 0) 
	for _, value := range nums {
		if value % multiple != 0 || value == multiple {
			filteredNums = append(filteredNums, value)
		}
	}
	return filteredNums
} 
