package canCompleteCircuit

import "sync"

// There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].
//
// You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station.
// You begin the journey with an empty tank at one of the gas stations.
//
// Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit
// once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if gas[i] < cost[i] {
			continue
		}
		if completeCircuit(gas, cost, i) {
			return i
		}
	}
	return -1
}

func completeCircuit(gas []int, cost []int, start int) bool {
	tank := 0
	for i := start; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			return false
		}
	}
	for i := 0; i < start; i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			return false
		}
	}
	return true
}

func canCompleteCircuitVer2(gas []int, cost []int) int {
	totalGas, totalCost, tank, start := 0, 0, 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		tank += gas[i] - cost[i]
		if tank < 0 {
			tank = 0
			start = i + 1
		}
	}
	if totalGas < totalCost {
		return -1
	}
	return start
}

func canCompleteCircuitConcurrent(gas []int, cost []int) int {
	n := len(gas)
	resultChan := make(chan int, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			if completeCircuitVer2(gas, cost, start) {
				resultChan <- start
			} else {
				resultChan <- -1
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		if result != -1 {
			return result
		}
	}

	return -1
}

func completeCircuitVer2(gas []int, cost []int, start int) bool {
	tank := 0
	n := len(gas)
	for i := 0; i < n; i++ {
		idx := (start + i) % n
		tank += gas[idx] - cost[idx]
		if tank < 0 {
			return false
		}
	}
	return true
}
