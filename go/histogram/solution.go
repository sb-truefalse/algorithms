// Algorithm: Max rectangle area in the histogram
package algo

// O(n^2), V(n)
func NaiveSolution(list []int) int {
	maxArea := 0
	// On right
	for i := 0; i < len(list); i++ {
		countOfCurrentItem := 1
		for j := i + 1; j < len(list); j++ {
			if list[j] >= list[i] {
				countOfCurrentItem++
			} else {
				break
			}
		}

		// On left
		for j := i - 1; j >= 0; j-- {
			if list[j] >= list[i] {
				countOfCurrentItem++
			} else {
				break
			}
		}

		tempArea := list[i] * countOfCurrentItem
		if tempArea > maxArea {
			maxArea = tempArea
		}
	}

	return maxArea
}
