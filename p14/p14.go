package p14

func Solve(limit int) int {
	result := 1
	maxCount := 0
	for i := 1; i < limit; i++ {
		num := i
		count := 1
		for num != 1 {
			count++
			if num % 2 == 0 {
				num = num / 2
			} else {
				num = (3 * num) + 1
			}
		}
		if count > maxCount {
			maxCount = count
			result = i
		} 
	}
	return result
}
