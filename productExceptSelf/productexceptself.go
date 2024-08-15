package productExceptSelf

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// You must write an algorithm that runs in O(n) time and without using the division operation.

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	left, right := 1, 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			left *= nums[j]
		}
		for k := i + 1; k < len(nums); k++ {
			right *= nums[k]
		}
		result[i] = left * right
		left, right = 1, 1
	}
	return result
}

func productExceptSelfVer2(nums []int) []int {
	result := make([]int, len(nums))
	left, right := make([]int, len(nums)), make([]int, len(nums))
	left[0], right[len(nums)-1] = 1, 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := 0; i < len(nums); i++ {
		result[i] = left[i] * right[i]
	}
	return result
}

func productExceptSelfVer3(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Initialize the result array with 1s
	for i := range result {
		result[i] = 1
	}

	left := 1
	for i := 0; i < n; i++ {
		result[i] *= left
		left *= nums[i]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}

func productExceptSelfVer4(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	leftChan := make(chan []int)
	rightChan := make(chan []int)

	// Initialize the result array with 1s
	for i := range result {
		result[i] = 1
	}

	// Calculate left products in a goroutine
	go func() {
		left := 1
		leftProducts := make([]int, n)
		for i := 0; i < n; i++ {
			leftProducts[i] = left
			left *= nums[i]
		}
		leftChan <- leftProducts
	}()

	// Calculate right products in a goroutine
	go func() {
		right := 1
		rightProducts := make([]int, n)
		for i := n - 1; i >= 0; i-- {
			rightProducts[i] = right
			right *= nums[i]
		}
		rightChan <- rightProducts
	}()

	// Collect results from channels
	leftProducts := <-leftChan
	rightProducts := <-rightChan

	// Combine left and right products
	for i := 0; i < n; i++ {
		result[i] = leftProducts[i] * rightProducts[i]
	}

	return result
}

func productExceptSelfVer5(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	productExceptSelfHelper(nums, result, 0, 1)
	return result
}

func productExceptSelfHelper(nums, result []int, index, currentProduct int) int {
	if index == len(nums) {
		return 1
	}
	leftProduct := currentProduct
	rightProduct := productExceptSelfHelper(nums, result, index+1, currentProduct*nums[index])
	result[index] = leftProduct * rightProduct
	return rightProduct * nums[index]
}
