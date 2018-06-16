package practice1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	testCases := [][]int{
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{2, 1, 4, 3},
		{2, 3, 4, 4},
	}

	for _, testCase := range testCases {
		assert.Equal(t, 4, Max(testCase...))
	}
}
