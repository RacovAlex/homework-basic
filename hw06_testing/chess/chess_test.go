package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChessBoard(t *testing.T) {
	testCases := []struct {
		desc string
		size int
		want string
	}{
		{
			desc: "normal case",
			size: 3,
			want: "# #\n # #\n# #\n",
		},
		{
			desc: "zero size",
			size: 0,
			want: "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := chessBoard(tc.size)
			assert.Equal(t, tc.want, got)
		})
	}
}
