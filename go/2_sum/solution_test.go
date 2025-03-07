// Tests
package algo

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	x := 5
	got1, got2 := TwoSum(list, x)
	if got1 != 0 && got2 != 3 {
		t.Errorf("TwoSum(%#v, %d) = %d %d; want 0 3", list, x, got1, got2)
	}
}

func TestTwoSumWith2Pointers(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	x := 5
	got1, got2 := TwoSumWith2Pointers(list, x)
	if got1 != 0 && got2 != 3 {
		t.Errorf("TwoSumWith2Pointers(%#v, %d) = %d %d; want 0 3", list, x, got1, got2)
	}
}

func TestTwoSumWithHash(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	x := 5
	got1, got2 := TwoSumWithHash(list, x)
	if got1 != 0 && got2 != 3 {
		t.Errorf("TwoSumWithHash(%#v, %d) = %d %d; want 0 3", list, x, got1, got2)
	}
}
