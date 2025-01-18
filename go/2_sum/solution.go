// Algorithm: Two Sum
package algo

import (
	"sort"
)

// Method: Naive
func TwoSum(numbers []int, x int) (int, int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				// next
			} else if numbers[i]+numbers[j] == x {
				return i, j
			}
		}
	}

	return -1, -1
}

// Method: 2-pointer
func TwoSumWith2Pointers(numbers []int, x int) (int, int) {
	sort.Ints(numbers)

	i := 0
	j := len(numbers) - 1

	for i < j {
		sum := numbers[i] + numbers[j]

		if sum == x {
			return i, j
		} else if sum > x {
			j -= 1
		} else {
			i += 1
		}
	}

	return -1, -1
}

// Method: Hash table
func TwoSumWithHash(numbers []int, x int) (int, int) {
	hsh := make(map[int]int)

	for i := 0; i < len(numbers); i++ {
		hsh[x-numbers[i]] = i
	}

	for i := 0; i < len(numbers); i++ {
		value, ok := hsh[numbers[i]]

		if ok {
			return i, value
		}
	}

	return -1, -1
}
