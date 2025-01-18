// Tests
package algo

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	x := 5
	got1, got2 := TwoSum(list, x)
	if got1 == -1 && got2 == -1 {
		t.Errorf("simpleAlgorithm(%#v, %d) = %d %d; want 2 3", list, x, got1, got2)
	}
}

func TestTwoSumWith2Pointers(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	x := 5
	got1, got2 := TwoSumWith2Pointers(list, x)
	if got1 == -1 && got2 == -1 {
		t.Errorf("simpleAlgorithm(%#v, %d) = %d %d; want 2 3", list, x, got1, got2)
	}
}

func TestTwoSumWithHash(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	x := 5
	got1, got2 := TwoSumWithHash(list, x)
	if got1 == -1 && got2 == -1 {
		t.Errorf("simpleAlgorithm(%#v, %d) = %d %d; want 2 3", list, x, got1, got2)
	}
}
