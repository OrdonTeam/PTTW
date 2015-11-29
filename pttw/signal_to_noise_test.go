package pttw

import (
	"testing"
	"github.com/assertgo/assert"
)

func TestSumShouldSum(t *testing.T) {
	assert := assert.New(t)
	first := []float64{1, 2, 3}
	second := []float64{2, 3, 4}
	sum := AddSignalToNoise(first, second, 1)
	assert.That(sum[0]).IsEqualTo(float64(3))
	assert.That(sum[1]).IsEqualTo(float64(5))
	assert.That(sum[2]).IsEqualTo(float64(7))
}
