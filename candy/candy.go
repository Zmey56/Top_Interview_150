package candy

// There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.
//
// You are giving candies to these children subjected to the following requirements:
//
// Each child must have at least one candy.
// Children with a higher rating get more candies than their neighbors.
// Return the minimum number of candies you need to have to distribute the candies to the children.

func candy(ratings []int) int {
	candySlice := make([]int, len(ratings), len(ratings))
	for i := range candySlice {
		candySlice[i] = 1
	}

	changed := true

	for {
		changed = false
		for i := 0; i < len(ratings); i++ {
			if i > 0 && ratings[i] > ratings[i-1] && candySlice[i] <= candySlice[i-1] {
				candySlice[i] = candySlice[i-1] + 1
				changed = true
			}
			if i < len(ratings)-1 && ratings[i] > ratings[i+1] && candySlice[i] <= candySlice[i+1] {
				candySlice[i] = candySlice[i+1] + 1
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	candies := 0
	for _, v := range candySlice {
		candies += v
	}

	return candies

}

func candyVer2(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	candies := make([]int, n)
	// Initialize all candies with 1 since every child must have at least one candy
	for i := range candies {
		candies[i] = 1
	}

	// Left to right
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Right to left
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	// Sum up all candies
	totalCandies := 0
	for _, candy := range candies {
		totalCandies += candy
	}

	return totalCandies
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func candyVer3(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	totalCandies := 1
	up := 1
	down := 0
	peak := 1

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			up++
			peak = up
			down = 0
			totalCandies += up
		} else if ratings[i] == ratings[i-1] {
			up = 1
			down = 0
			peak = 1
			totalCandies += 1
		} else {
			down++
			up = 1
			totalCandies += down
			if down >= peak {
				totalCandies++
			}
		}
	}

	return totalCandies
}

func candyVer4(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	// Channels to communicate between goroutines
	left := make(chan int, n)
	right := make(chan int, n)

	// Goroutine for the left to right pass
	go func() {
		leftCandies := make([]int, n)
		leftCandies[0] = 1
		for i := 1; i < n; i++ {
			if ratings[i] > ratings[i-1] {
				leftCandies[i] = leftCandies[i-1] + 1
			} else {
				leftCandies[i] = 1
			}
		}
		for _, c := range leftCandies {
			left <- c
		}
		close(left)
	}()

	// Goroutine for the right to left pass
	go func() {
		rightCandies := make([]int, n)
		rightCandies[n-1] = 1
		for i := n - 2; i >= 0; i-- {
			if ratings[i] > ratings[i+1] {
				rightCandies[i] = rightCandies[i+1] + 1
			} else {
				rightCandies[i] = 1
			}
		}
		for _, c := range rightCandies {
			right <- c
		}
		close(right)
	}()

	// Calculate the total minimum candies required
	totalCandies := 0
	for i := 0; i < n; i++ {
		leftCandy := <-left
		rightCandy := <-right
		totalCandies += max(leftCandy, rightCandy)
	}

	return totalCandies
}
