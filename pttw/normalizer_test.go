package pttw

import (
	"testing"
	"github.com/assertgo/assert"
)

func TestNormalizerShouldNormalizeSingleElement(t *testing.T) {
	assert := assert.New(t)
	slice := []float64{3}
	normalized := Normalize(slice)
	assert.That(normalized[0]).IsEqualTo(float64(1))
}

func TestNormalizerShouldNormalizeSameElements(t *testing.T) {
	assert := assert.New(t)
	slice := []float64{3, 3}
	normalized := Normalize(slice)
	assert.That(normalized[0]).IsEqualTo(float64(1))
	assert.That(normalized[1]).IsEqualTo(float64(1))
}

func TestNormalizerShouldNormalizeDifferentElements(t *testing.T) {
	assert := assert.New(t)
	slice := []float64{3, 6}
	normalized := Normalize(slice)
	assert.That(normalized[0]).IsEqualTo(float64(0.5))
	assert.That(normalized[1]).IsEqualTo(float64(1))
}