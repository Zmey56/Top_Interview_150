package merge

import (
	"sort"
	"sync"
)

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n,
//representing the number of elements in nums1 and nums2 respectively.
//
//Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//
//The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
//To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
//and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)

	sort.Ints(nums1)

}

func mergeVer2(nums1 []int, m int, nums2 []int, n int) {
	var wg sync.WaitGroup
	ch := make(chan int, m+n)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, v := range nums1[:m] {
			ch <- v
		}
	}()

	go func() {
		defer wg.Done()
		for _, v := range nums2 {
			ch <- v
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	nums1 = nums1[:0]
	for v := range ch {
		nums1 = append(nums1, v)
	}

	sort.Ints(nums1)
}
