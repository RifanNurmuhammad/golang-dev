package refactor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstStringInBracket(t *testing.T) {

	type testCase struct {
		title, param, expected string
	}

	tests := []testCase{
		{
			title:    "#1 Test get first string",
			param:    "(kata pertama)",
			expected: "kata pertama",
		},
		{
			title:    "#2 Test get first string without close bracket",
			param:    "(kata pertama",
			expected: "",
		},
		{
			title:    "#3 blank parameter",
			param:    "",
			expected: "",
		},
		{
			title:    "#4 Test get first string in bracket",
			param:    "(kata pertama) dalam sebuah abjad",
			expected: "kata pertama",
		},
		{
			title:    "#5 Test get first string in bracket",
			param:    "ini (kata pertama) dalam sebuah abjad",
			expected: "kata pertama",
		},
		{
			title:    "#6 Test get first string in bracket",
			param:    "ini (kata pertama) dan (kata kedua)",
			expected: "kata pertama",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			assert.Equal(t, findFirstStringInBracket(tt.param), tt.expected)
		})
	}

}

func Benchmark1(b *testing.B) {
	for n := 0; n < 1000; n++ {
		findFirstStringInBracket("(kata pertama)")
	}
}
