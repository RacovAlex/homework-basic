package figure

import (
	"errors"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
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
			err:  errors.New("радиус не может быть отрицательным"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := tC.fig.Area()
			if err != nil {
				assert.EqualError(t, err, tC.err.Error())
			}
			assert.Equal(t, tC.want, got)
		})
	}
}
