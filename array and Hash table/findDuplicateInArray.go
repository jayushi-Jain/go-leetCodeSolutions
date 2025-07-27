package main

import (
	"fmt"
)

func findDuplicate(nums []int) bool{
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++

		if counts[num] > 1 {
			return true
		} 
	}

    return false
}
