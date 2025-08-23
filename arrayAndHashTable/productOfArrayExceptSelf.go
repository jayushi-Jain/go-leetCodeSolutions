func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	//left numbers calculation
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	//right numbers calculations
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] = result[i] * rightProduct
		rightProduct *= nums[i]
	}

	return result
}