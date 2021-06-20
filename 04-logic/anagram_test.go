package anagram

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupingAnagram(t *testing.T) {
	params := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	fmt.Println(GetGroupAnagram(params))
	assert.Equal(t,
		GetGroupAnagram(params),
		[][]string{
			{"kita", "atik", "tika"},
			{"aku", "kua"},
			{"kia"},
			{"makan"},
		})
}
