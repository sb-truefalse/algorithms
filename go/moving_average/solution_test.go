// Tests
package algo

import (
	"reflect"
	"testing"
)

func TestMovingAvrg(t *testing.T) {
	timeseries := []int{1, 2, 3, 4, 5, 6}
	k := 3
	got := MovingAvrg(timeseries, k)
	if !reflect.DeepEqual(got, []float64{2, 3, 4, 5}) {
		t.Errorf("MovingAvrg(%#v, %d) = %#v; want []", timeseries, k, got)
	}
}

func TestMovingAvrgWith2Pointers(t *testing.T) {
	timeseries := []int{1, 2, 3, 4, 5, 6}
	k := 3
	got := MovingAvrgWith2Pointers(timeseries, k)
	if !reflect.DeepEqual(got, []float64{2, 3, 4, 5}) {
		t.Errorf("MovingAvrgWith2Pointers(%#v, %d) = %#v; want []", timeseries, k, got)
	}
}
