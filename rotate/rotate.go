package rotate

import "sync"

//Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	result := make([]int, len(nums))
	result = append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, result)
}

func rotate2(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func rotate3(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 || n <= 1 {
		return
	}

	count := 0
	for start := 0; count < n; start++ {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % n
			nums[next], prev = prev, nums[next]
			current = next
			count++
			if start == current {
				break
			}
		}
	}
}

func rotate4(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	ch := make(chan int, len(nums)) // Buffered channel to hold all elements
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for _, v := range append(nums[len(nums)-k:], nums[:len(nums)-k]...) {
			ch <- v
		}
		close(ch)
	}()

	i := 0
	for v := range ch {
		nums[i] = v
		i++
	}

	wg.Wait()
}
