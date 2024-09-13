package figure

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCircle_Area(t *testing.T) {
	testCases := []struct {
		desc string
		fig  Circle
		want float32
		err  error
	}{
		{
			desc: "normal circle",
			fig:  Circle{5},
			want: math.Pi * 5 * 5,
			err:  nil,
		},
		{
			desc: "zero circle",
			fig:  Circle{0},
			want: math.Pi * 0 * 0,
			err:  nil,
		},
		{
			desc: "negative radius",
			fig:  Circle{-1},
			want: 0.0,
			err:  ErrNegativeRadius,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := tc.fig.Area()
			//	if err != nil {
			//		assert.EqualError(t, err, tc.err.Error())
			//	}
			require.ErrorIs(t, err, tc.err)
			assert.Equal(t, tc.want, got)
		})
	}
}
