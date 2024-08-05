package hIndex

import (
	"sort"
	"sync"
)

// Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper,
// return the researcher's h-index.
//
// According to the definition of h-index on Wikipedia: The h-index is defined as the maximum value of h such that the given
// researcher has published at least h papers that have each been cited at least h times.

func hIndex(citations []int) int {

	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	count := 0

	for i, j := range citations {
		if i >= j {
			break
		}
		count++
	}

	return count
}

func hIndex2(citations []int) int {
	n := len(citations)
	counts := make([]int, n+1)

	for _, c := range citations {
		if c >= n {
			counts[n]++
		} else {
			counts[c]++
		}
	}

	h := 0
	total := 0

	for i := n; i >= 0; i-- {
		total += counts[i]
		if total >= i {
			h = i
			break
		}
	}

	return h
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func parallelSort(citations []int) []int {
	if len(citations) <= 1 {
		return citations
	}
	mid := len(citations) / 2
	var left, right []int
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		left = parallelSort(citations[:mid])
	}()

	go func() {
		defer wg.Done()
		right = parallelSort(citations[mid:])
	}()

	wg.Wait()
	return merge(left, right)
}

func hIndex3(citations []int) int {
	sortedCitations := parallelSort(citations)

	hIndexChan := make(chan int)
	var wg sync.WaitGroup
	numGorutines := 4
	chunkSize := (len(sortedCitations) + numGorutines - 1) / numGorutines

	for i := 0; i < numGorutines; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			for j := start; j < start+chunkSize && j < len(sortedCitations); j++ {
				if sortedCitations[j] <= j {
					hIndexChan <- j
					return
				}
			}
		}(i * chunkSize)
	}

	go func() {
		wg.Wait()
		close(hIndexChan)
	}()

	hIndex := len(sortedCitations)
	for idx := range hIndexChan {
		if idx < hIndex {
			hIndex = idx
		}
	}

	return hIndex
}
