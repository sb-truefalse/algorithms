package main

/* Two Sum */

import (
	"fmt"
	"sort"
)

// Naive
func simpleAlgorithm(numbers []int, x int) (int, int) {
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

//  2-pointer
func tpmOptimizedAlgorithm(numbers []int, x int) (int, int) {
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

// Memory+ (hash table)
func mapOptimizedAlgorithm(numbers []int, x int) (int, int) {
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

func test() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	x := 5
	fmt.Println(simpleAlgorithm(list, x))
	fmt.Println(tpmOptimizedAlgorithm(list, x))
	fmt.Println(mapOptimizedAlgorithm(list, x))
}

func main() {
}
