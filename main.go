package main

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	temp := make(map[int][]int)

	for _, num := range nums {
		temp[num] = append(temp[num], num)
	}

}
