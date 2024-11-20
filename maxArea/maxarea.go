package maxArea

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of
// the ith line are (i, 0) and (i, height[i]).
//
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
//
// Return the maximum amount of water a container can store.
//
// Notice that you may not slant the container.

func maxArea(height []int) int {
	n := len(height)
	maxArea := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			h := min(height[i], height[j])
			width := j - i
			currentArea := h * width
			if currentArea > maxArea {
				maxArea = currentArea
			}
		}
	}
	return maxArea
}

func maxAreaVer2(height []int) int {
	maxArea := 0

	for i := 0; i < len(height); i++ {
		currentMax := 0
		for j := i + 1; j < len(height); j++ {
			h := min(height[i], height[j])
			width := j - i
			currentMax = max(currentMax, h*width)
		}
		maxArea = max(maxArea, currentMax)
	}

	return maxArea
}

func maxAreaVer3(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		h := min(height[left], height[right])
		width := right - left
		currentArea := h * width
		maxArea = max(maxArea, currentArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
