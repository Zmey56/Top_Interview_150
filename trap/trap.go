package trap

// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}

func trapVer2(height []int) int {
	if len(height) < 3 {
		return 0
	}

	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	water := 0
	for i := 0; i < n; i++ {
		water += min(leftMax[i], rightMax[i]) - height[i]
	}

	return water
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trapVer3(height []int) int {
	if len(height) < 3 {
		return 0
	}

	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftChan := make(chan []int)
	rightChan := make(chan []int)

	go computeLeftMax(height, leftChan)
	go computeRightMax(height, rightChan)

	leftMax = <-leftChan
	rightMax = <-rightChan

	water := 0
	for i := 0; i < n; i++ {
		water += min(leftMax[i], rightMax[i]) - height[i]
	}

	return water
}

func computeLeftMax(height []int, ch chan []int) {
	n := len(height)
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	ch <- leftMax
}

func computeRightMax(height []int, ch chan []int) {
	n := len(height)
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	ch <- rightMax
}
