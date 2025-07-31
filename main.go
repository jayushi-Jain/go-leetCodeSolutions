package main

import (
	"fmt"
)



func main() {
	numbers := []int{1,2,3,4}
	target := 3

	left := 0
	right := len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			fmt.Println("num1: ", numbers[left])
			fmt.Println("num2: ", numbers[right])
			break
		}else if sum < 0 {
			left++
		}else {
			right--
		}
	}

	

}