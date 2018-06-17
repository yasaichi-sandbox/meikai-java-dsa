package practice9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumof(t *testing.T) {
	assert.Equal(t, 12, SumOf(3, 5))
	assert.Equal(t, 15, SumOf(6, 4))
	assert.Equal(t, -14, SumOf(1, -5))
}
