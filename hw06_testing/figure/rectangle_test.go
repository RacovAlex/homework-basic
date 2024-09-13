package figure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRectangle_Area(t *testing.T) {
	testCases := []struct {
		desc string
		fig  Rectangle
		area float32
		err  error
	}{
		{
			desc: "normal rectangle",
			fig:  Rectangle{10, 5},
			area: 50,
			err:  nil,
		},
		{
			desc: "zero rectangle",
			fig:  Rectangle{0, 10},
			area: 0,
			err:  nil,
		},
		{
			desc: "negative rectangle",
			fig:  Rectangle{10, -5},
			area: 0,
			err:  ErrNegativeHeightWidth,
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
