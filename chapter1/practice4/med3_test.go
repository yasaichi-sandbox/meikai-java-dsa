package practice4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMed3(t *testing.T) {
	testCases := [13][4]int{
		{3, 2, 1, 2},
		{3, 2, 2, 2},
		{3, 1, 2, 2},
		{3, 2, 3, 3},
		{2, 1, 3, 2},
		{3, 3, 2, 3},
		{3, 3, 3, 3},
		{2, 2, 3, 2},
		{2, 3, 1, 2},
		{2, 3, 2, 2},
		{1, 3, 2, 2},
		{2, 3, 3, 3},
		{1, 2, 3, 2},
	}

	for _, testCase := range testCases {
		a := testCase[0]
		b := testCase[1]
		c := testCase[2]
		expected := testCase[3]

		t.Logf("Med3(%d,%d,%d) = %d\n", a, b, c, Med3(a, b, c))
		assert.Equal(t, expected, Med3(a, b, c))
	}
}
