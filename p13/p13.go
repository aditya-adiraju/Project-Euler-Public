package p13

import (
	"fmt"
	"math"
	"strconv"
)

func Solve(nums [100]string) (result int){

  digitArray := addNumStr(makeDigitArray(nums[0]), makeDigitArray(nums[1])) 

  for i := 2; i < 100; i++ {
    digitArray = addNumStr(makeDigitArray(nums[i]), digitArray)
    fmt.Printf("Result: %#v\n", digitArray)
    fmt.Println(len(digitArray))
  }


  for i, n := len(digitArray) - 10, 0; i < len(digitArray); i, n = i+1, n+1 {
    result += int(math.Pow10(n)) * digitArray[i]
  }

  return result
  
}


func makeDigitArray(numStr string) (result []int) {
  
  for _, digit := range numStr {
    
    convertedInt, err := strconv.Atoi(string(digit))
    if err != nil {
      panic(err)
    }

    result = append(result, convertedInt)
  }

  for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
    result[i], result[j] = result[j], result[i]
  }


  fmt.Printf("DigitArray: %#v\n", result)
  return result
} 

// returns the sum of x and y as int slices
// NOTE: x, y and the result are arrays with the MSD being the end 
// This is largely a convenience thing
// ASSUMPTION: len(x) == 50 && len(y) >= len(x)

func addNumStr(x []int, y []int) (result []int) {
  carryOver := 0
  for i := 0; i < 50; i++ {
    result = append(result, (x[i] + y[i] + carryOver) % 10)
    carryOver = (x[i] + y[i] + carryOver)/ 10
  }

  if len(y) > 50 {
    j := 50
    for carryOver != 0  && j < len(y) {

      result = append(result, (y[j] + carryOver) % 10)
      carryOver = (y[j] + carryOver) / 10
      j++
    }

    if j < len(y) {
      result = append(result, y[j:]...)
    }
  } 

  if carryOver != 0 {
    result = append(result, carryOver)
  }
 
  return result
}
