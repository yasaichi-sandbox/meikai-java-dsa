package practice2_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	testCases := [][]int{
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{1, 1, 2, 3},
	}

	for _, testCase := range testCases {
		assert.Equal(t, 1, Min(testCase...))
	}
}
