package removeDuplicatesMedium

// Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element
//appears at most twice. The relative order of the elements should be kept the same.
//
//Since it is impossible to change the length of the array in some languages, you must instead have the result be placed
//in the first part of the array nums. More formally, if there are k elements after removing the duplicates,
//then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
//
//Return k after placing the final result in the first k slots of nums.
//
//Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
//
//Custom Judge:
//
//The judge will test your solution with the following code:
//
//int[] nums = [...]; // Input array
//int[] expectedNums = [...]; // The expected answer with correct length
//
//int k = removeDuplicates(nums); // Calls your implementation
//
//assert k == expectedNums.length;
//for (int i = 0; i < k; i++) {
//    assert nums[i] == expectedNums[i];
//}
//If all assertions pass, then your solution will be accepted.

func removeDuplicates(nums []int) int {
	mapNums := make(map[int]int8)
	count := 0
	for i := 0; i < len(nums); i++ {
		val, ok := mapNums[nums[i]]
		if !ok {
			mapNums[nums[i]]++
			nums[count] = nums[i]
			count++
		} else {
			if val < 2 {
				mapNums[nums[i]]++
				nums[count] = nums[i]
				count++
			}
		}
	}
	return count
}

func removeDuplicatesVer2(nums []int) int {
	count := 1
	dupCount := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			dupCount++
		} else {
			dupCount = 1
		}
		if dupCount <= 2 {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func removeDuplicatesVer3(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	count := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[count-2] {
			nums[count] = nums[i]
			count++
		}
	}

	return count
}
