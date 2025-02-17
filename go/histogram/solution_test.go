package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNaiveSolution(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want int
	}{
		{
			name: "First test",
			list: []int{1, 2, 3},
			want: 4,
		},
		{
			name: "Second test",
			list: []int{2, 7, 6, 9, 7, 5, 7, 3, 5},
			want: 30,
		},
		{
			name: "Third test",
			list: []int{2},
			want: 2,
		},
		{
			name: "Fourth test",
			list: []int{1, 1, 1, 1},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NaiveSolution(tt.list), tt.want)
		})
	}
}
