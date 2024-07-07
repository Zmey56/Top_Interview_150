package removeDuplicates

//Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique
//element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
//
//Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
//
//Change the array nums such that the first k elements of nums contain the unique elements in the order they were
//present in nums initially. The remaining elements of nums are not important as well as the size of nums.
//Return k.
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
	mapNums := make(map[int]struct{})
	count := 0
	for i := 0; i < len(nums); i++ {
		_, ok := mapNums[nums[i]]
		if !ok {
			mapNums[nums[i]] = struct{}{}
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func removeDuplicatesVer2(nums []int) int {
	count := 1
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] != nums[i-1] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func removeDuplicatesVer3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Create a channel to communicate between goroutines
	ch := make(chan int)

	// Goroutine to iterate through the array and send unique elements to the channel
	go func() {
		defer close(ch)
		last := nums[0]
		ch <- last

		for _, num := range nums[1:] {
			if num != last {
				ch <- num
				last = num
			}
		}
	}()

	// Read unique elements from the channel and update the original array
	k := 0
	for unique := range ch {
		nums[k] = unique
		k++
	}

	return k
}
