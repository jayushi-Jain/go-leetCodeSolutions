func sum(numbers []int, target int) []int{	
	left := 0
	right := len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{number[left],numbers[right]}
		}else sum < 0 {
			left++
		}else {
			right--
		}
	}
}