// Tests
package algo

import (
	"testing"
)

func TestIsPrimeNumber(t *testing.T) {
	x := 5
	got := IsPrimeNumber(x)
	if !got {
		t.Errorf("IsPrimeNumber(%d) = %t; want got!", x, got)
	}
}

func TestIsPrimeNumberWithSqrt(t *testing.T) {
	x := 5
	got := IsPrimeNumberWithSqrt(x)
	if !got {
		t.Errorf("IsPrimeNumberWithSqrt(%d) = %t; want got!", x, got)
	}
}
