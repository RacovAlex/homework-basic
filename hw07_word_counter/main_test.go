package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var longString = `Эти тестовые случаи помогут убедиться в том, что функция 
	countWords работает корректно и обрабатывает различные сценарии ввода. 
	Тестирование на этих примерах позволит выявить возможные ошибки и улучшить 
	стабильность функции.`

func TestCountWords(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		args string
		want map[string]int
	}{
		{
			name: "empty string",
			args: "",
			want: map[string]int{},
		},
		{
			name: "normal string",
			args: "goodbye, my friend, goodbye",
			want: map[string]int{"goodbye": 2, "friend": 1, "my": 1},
		},
		{
			name: "register and symbols accounting",
			args: "Register accounting - is so IMPORTING!",
			want: map[string]int{"accounting": 1, "importing": 1, "is": 1, "register": 1, "so": 1},
		},
		{
			name: "repeat words and spices",
			args: "Words! words;    words",
			want: map[string]int{"words": 3},
		},
		{
			name: "long string",
			args: longString,
			want: map[string]int{
				"эти":          1,
				"тестовые":     1,
				"случаи":       1,
				"помогут":      1,
				"убедиться":    1,
				"в":            1,
				"том":          1,
				"что":          1,
				"функция":      1,
				"функции":      1,
				"countwords":   1,
				"работает":     1,
				"корректно":    1,
				"и":            2,
				"обрабатывает": 1,
				"различные":    1,
				"сценарии":     1,
				"ввода":        1,
				"тестирование": 1,
				"на":           1,
				"этих":         1,
				"примерах":     1,
				"позволит":     1,
				"выявить":      1,
				"возможные":    1,
				"ошибки":       1,
				"улучшить":     1,
				"стабильность": 1,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := countWords(tc.args)
			require.Equal(t, tc.want, got)
		})
	}
}
