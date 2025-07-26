package main

import (
	"fmt"
)

func twoSums(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, ok := seen[complement]; ok {
			return []int{index,i}
		}

		seen[num] = i
	}

	return nil

}

func main() {
	nums := []int{2, 7, 11, 15}
    target := 9

    output := twoSums(nums,target)
 	fmt.Println("Result: ", output)
}