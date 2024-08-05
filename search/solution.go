package main

/* Поиск */

import (
	"sort"
)

// Линейный поиск
func linearAlgorithm(list []int, x int) int {
	for i := 0; i < len(list); i++ {
		if x == list[i] {
			return i
		}
	}

	return -1
}

// Бинарный поиск
func binaryAlgorithm(list []int, x int) int {
	l, r := 0, (len(list) - 1)
	sort.Ints(list)

	for l <= r {
		mid := (r + l) / 2

		if list[mid] == x {
			return mid
		} else if list[mid] > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func main() {
}
