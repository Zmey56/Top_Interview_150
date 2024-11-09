package twoSum

// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that
// they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
//
// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
//
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
//
// Your solution must use only constant extra space.

func twoSum(numbers []int, target int) []int {

	left := 0
	right := len(numbers) - 1

	result := []int{0, 0}

	for i := 0; i < len(numbers); i++ {
		if numbers[left]+numbers[right] == target {
			result[0] = left + 1
			result[1] = right + 1
			break
		}
		if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}

	return result

}

func twoSumVer2(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum > target {
			right--
		} else {
			left++
		}
	}

	return nil
}

func twoSumVer3(numbers []int, target int) []int {
	type result struct {
		index1 int
		index2 int
		found  bool
	}

	ch := make(chan result)
	left := 0
	right := len(numbers) - 1

	go func() {
		for left < right {
			sum := numbers[left] + numbers[right]
			if sum == target {
				ch <- result{left + 1, right + 1, true}
				return
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
		ch <- result{0, 0, false}
	}()

	res := <-ch
	if res.found {
		return []int{res.index1, res.index2}
	}
	return nil
}
