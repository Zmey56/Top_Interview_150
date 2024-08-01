package canJumpII

// You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].
//
// Each element nums[i] represents the maximum length of a forward jump from index i. In other words,
// if you are at nums[i], you can jump to any nums[i + j] where:
//
// 0 <= j <= nums[i] and
// i + j < n
// Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	jums := 0
	maxValue := 0
	currentEnd := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]+i > maxValue {
			maxValue = nums[i] + i
		}

		if i == currentEnd {
			jums++
			currentEnd = maxValue
			if maxValue >= len(nums)-1 {
				break
			}
		}
	}

	return jums
}

func jump2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	jumps := 0
	maxReach := 0
	currentEnd := 0
	for i := 0; i < len(nums); i++ {
		maxReach = max(maxReach, i+nums[i])

		if i == currentEnd {
			jumps++
			currentEnd = maxReach
			if maxReach >= len(nums)-1 {
				break
			}
		}
	}
	return jumps
}
