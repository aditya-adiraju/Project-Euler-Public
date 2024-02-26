package p8

import (
	"log"
	"strconv"
)

func Solve(digits string, n int) int {
	maximum := getNProduct(digits, 0, n - 1)
	product := maximum

	for start, end := 0, n; end < len(digits); start, end = start + 1, end + 1 {
		conv, err := strconv.Atoi(string(digits[start]))
		if err != nil {
			log.Fatal(err)
		}

		if conv != 0 {
			product = product / conv
		} else {
			product = getNProduct(digits, start + 1, end - 1)
		}

		conv, err = strconv.Atoi(string(digits[end]))
		if err != nil {
			log.Fatal(err)
		}

		product *= conv

		maximum = max(product, maximum)
	} 

	return maximum
}

func getNProduct(digits string, start int,  end int) int{	
	result := 1

	for i := start; i <= end; i++ {
		num, err := strconv.Atoi(string(digits[i]))
		if err != nil {
			log.Fatal(err)
		}
		result *= num
	}

	return result
}
