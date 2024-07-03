package removeElement

import "sync"

// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of
//the elements may be changed. Then return the number of elements in nums which are not equal to val.
//
//Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:
//
//Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
//The remaining elements of nums are not important as well as the size of nums.
//Return k.
//Custom Judge:
//
//The judge will test your solution with the following code:
//
//int[] nums = [...]; // Input array
//int val = ...; // Value to remove
//int[] expectedNums = [...]; // The expected answer with correct length.
//                            // It is sorted with no values equaling val.
//
//int k = removeElement(nums, val); // Calls your implementation
//
//assert k == expectedNums.length;
//sort(nums, 0, k); // Sort the first k elements of nums
//for (int i = 0; i < actualLength; i++) {
//    assert nums[i] == expectedNums[i];
//}
//If all assertions pass, then your solution will be accepted.

func removeElement(nums []int, val int) int {
	countNotVal := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[countNotVal] = nums[i]
			countNotVal++
		}
	}

	return countNotVal

}

func removeElementVer2(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	return i

}

func removeElementVer3(nums []int, val int) int {
	var wg sync.WaitGroup
	ch := make(chan int, 1) // Buffered channel with capacity 1

	wg.Add(1)

	go func() {
		defer wg.Done()
		for _, v := range nums {
			if v != val {
				ch <- v
			}
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	newNums := []int{}
	for v := range ch {
		newNums = append(newNums, v)
	}

	copy(nums, newNums) // Copy the elements back to nums
	return len(newNums)
}
