package figure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTriangle_Area(t *testing.T) {
	testCases := []struct {
		desc string
		fig  Triangle
		area float32
		err  error
	}{
		{
			desc: "normal triangle",
			fig:  Triangle{8, 6},
			area: 24,
			err:  nil,
		},
		{
			desc: "zero triangle",
			fig:  Triangle{0, 5},
			area: 0,
			err:  nil,
		},
		{
			desc: "negative triangle",
			fig:  Triangle{-1, 3},
			area: 0,
			err:  ErrNegativeBaseHeight,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := tc.fig.Area()
			require.ErrorIs(t, tc.err, err)
			assert.Equal(t, tc.area, got)
		})
	}
}
