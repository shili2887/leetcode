package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		index := 0
		gasSum := 0
		costSum := 0
		for index < n {
			j := (i + index) % n
			gasSum += gas[j]
			costSum += cost[j]
			if costSum > gasSum {
				break
			}
			index++
		}

		if index == n {
			return i
		} else {
			i += index + 1
		}
	}

	return -1
}

func main() {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
}
