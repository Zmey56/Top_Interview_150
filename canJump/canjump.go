package canJump

// You are given an integer array nums. You are initially positioned at the array's first index, and each element in
// the array represents your maximum jump length at that position.
//
// Return true if you can reach the last index, or false otherwise.
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	jump := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && i == len(nums)-2 {
			return false
		}

		if i+nums[i] == jump && nums[i] == 0 {
			return false
		}

		if i+nums[i] > jump {
			jump = i + nums[i]
		}
		if jump >= len(nums)-1 {
			return true
		}

	}
	return false
}

func canJump2(nums []int) bool {
	maxJump := 0
	for i, v := range nums {
		if i > maxJump {
			return false
		}
		maxJump = max(maxJump, i+v)
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump3(nums []int) bool {
	maxJump := 0
	for i := 0; i < len(nums); i++ {
		if i > maxJump {
			return false
		}
		maxJump = max(maxJump, i+nums[i])

		if maxJump >= len(nums)-1 {
			return true
		}
	}

	return false
}
