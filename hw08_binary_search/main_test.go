package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name   string
		slice  []int
		target int
		want   int
	}{
		{
			name:   "normal case",
			slice:  []int{1, 2, 9, 13, 25, 36, 47, 84, 90},
			target: 84,
			want:   7,
		},
		{
			name:   "not sort slice case",
			slice:  []int{13, 75, 24, 124, 47, 31, 15, 100},
			target: 100,
			want:   6,
		},
		{
			name:   "no target in slice case",
			slice:  []int{1, 2, 9, 13, 25, 36, 47, 84, 90},
			target: 0,
			want:   -1,
		},
		{
			name:   "zero slice case",
			slice:  []int{},
			target: 12,
			want:   -1,
		},
		{
			name:   "short slice case",
			slice:  []int{9},
			target: 9,
			want:   0,
		},
		{
			name:   "duplicates case",
			slice:  []int{1, 2, 3, 4, 4, 5},
			target: 4,
			want:   4,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := BinarySearch(tc.slice, tc.target)
			assert.Equal(t, tc.want, res)
		})
	}
}
