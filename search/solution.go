package main

/* Поиск */

import (
	"fmt"
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

func test() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := 10
	fmt.Println(linearAlgorithm(list, x) == 9)
	fmt.Println(binaryAlgorithm(list, x) == 9)
}

func main() {
}
