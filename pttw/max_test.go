package pttw

import (
	"testing"
	"github.com/assertgo/assert"
)

func TestMaxAbsShouldFindMax(t *testing.T) {
	assert := assert.New(t)
	slice := []float64{6}
	max := MaxAbs(slice)
	assert.That(max).IsEqualTo(float64(6))
}

func TestMaxAbsShouldFindMaxFromMultipleElements(t *testing.T) {
	assert := assert.New(t)
	slice := []float64{3, 7, 2}
	max := MaxAbs(slice)
	assert.That(max).IsEqualTo(float64(7))
}

func TestMaxAbsShouldFindMaxAbsWhenMaxIsNegative(t *testing.T) {
	assert := assert.New(t)
	slice := []float64{3, 7, 2, -8}
	max := MaxAbs(slice)
	assert.That(max).IsEqualTo(float64(8))
}