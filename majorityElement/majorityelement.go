package majorityElement

import (
	"math/rand"
	"sort"
	"time"
)

//Given an array nums of size n, return the majority element.
//
//The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element
//always exists in the array.

func majorityElement(nums []int) int {
	numMaps := make(map[int]int)
	for _, n := range nums {
		numMaps[n]++
	}

	majorELem := 0
	maxCount := 0

	for key, value := range numMaps {
		if value > maxCount {
			majorELem = key
			maxCount = value
		}
	}

	return majorELem
}

func majorityElementVer2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElementVer3(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
		if counts[num] > len(nums)/2 {
			return num
		}
	}
	return -1 // Placeholder, as the problem statement guarantees a majority element exists
}

func majorityElementVer4(nums []int) int {
	rand.Seed(time.Now().UnixNano())

	for {
		candidate := nums[rand.Intn(len(nums))]

		count := 0
		for _, num := range nums {
			if num == candidate {
				count++
			}
		}

		if count > len(nums)/2 {
			return candidate
		}
	}
}

func majorityElementVer5(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	mid := len(nums) / 2
	leftMajor := majorityElement(nums[:mid])
	rightMajor := majorityElement(nums[mid:])

	if leftMajor == rightMajor {
		return leftMajor
	}

	leftCount := countOccurrences(nums, leftMajor)
	rightCount := countOccurrences(nums, rightMajor)

	if leftCount > rightCount {
		return leftMajor
	} else {
		return rightMajor
	}
}

func countOccurrences(nums []int, target int) int {
	count := 0
	for _, num := range nums {
		if num == target {
			count++
		}
	}
	return count
}

// OptimizedMooreVoting finds the majority element in nums using the optimized Boyer-Moore Voting Algorithm.
func majorityElementVer6(nums []int) int {
	candidate := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}

	if count > len(nums)/2 {
		return candidate
	}

	return -1
}
