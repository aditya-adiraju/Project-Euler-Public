package p4

import (
	"fmt"
	"strconv"
)

func Solve() int {
	maximum := 0
	for i := 999; i >= 100; i -- {
		for j := 999; j >= i; j-- {
			if i*j > maximum && isPalindrome(i * j) {
				maximum = i * j
				fmt.Println(i, j)

			}
		}
	}
	return maximum
}

func isPalindrome(num int) bool {	
	numStr := strconv.Itoa(num)	
	for i, j := 0, len(numStr) - 1; i < j; i , j =  i + 1, j - 1 {
		if numStr[i] != numStr[j] {
			return false
		}
	}
	return true 
}
