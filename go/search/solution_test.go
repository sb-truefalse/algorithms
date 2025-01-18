// Tests
package algo

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	x := 5
	got := LinearSearch(list, x)
	if got != 4 {
		t.Errorf("LinearSearch(%#v, %d) = %d; want 4", list, x, got)
	}
}

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	x := 5
	got := BinarySearch(list, x)
	if got != 4 {
		t.Errorf("BinarySearch(%#v, %d) = %d; want 4", list, x, got)
	}
}
