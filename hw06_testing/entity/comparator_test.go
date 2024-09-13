package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var annaKarenina = NewBook(
	1,
	"Анна Каренина",
	"Лев Толстой",
	1878,
	1124,
	4.75)

var martinEden = NewBook(
	2,
	"Мартин Иден",
	"Джек Лондон",
	1909,
	448,
	4.60)

func TestBookComparator_Compare(t *testing.T) {
	testCases := []struct {
		desc string
		bc   BookComparator
		b1   *Book
		b2   *Book
		want bool
		err  error
	}{
		{
			desc: "normal case by Year",
			bc:   BookComparator{CompareByYear},
			b1:   annaKarenina,
			b2:   martinEden,
			want: false,
			err:  nil,
		},
		{
			desc: "normal case by Size",
			bc:   BookComparator{CompareBySize},
			b1:   annaKarenina,
			b2:   martinEden,
			want: true,
			err:  nil,
		},
		{
			desc: "normal case by Rate",
			bc:   BookComparator{CompareByRate},
			b1:   annaKarenina,
			b2:   martinEden,
			want: true,
			err:  nil,
		},
		{
			desc: "second empty book",
			bc:   BookComparator{CompareByYear},
			b1:   annaKarenina,
			b2:   &Book{},
			want: true,
			err:  nil,
		},
		{
			desc: "multiple empty books",
			bc:   BookComparator{CompareBySize},
			b1:   &Book{},
			b2:   &Book{},
			want: false,
			err:  ErrEntityNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := tc.bc.Compare(*tc.b1, *tc.b2)
			require.ErrorIs(t, tc.err, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
