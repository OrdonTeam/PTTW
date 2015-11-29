package pttw

import (
	"testing"
	"github.com/assertgo/assert"
)

func TestSumShouldSum(t *testing.T) {
	assert := assert.New(t)
	first := []float64{0.5, 1, -1}
	second := []float64{-0.5, -1, 1}
	sum := AddSignalToNoise(first, second, 0)
	assert.That(sum[0]).IsEqualTo(float64(0))
	assert.That(sum[1]).IsEqualTo(float64(0))
	assert.That(sum[2]).IsEqualTo(float64(0))
}

func TestSumShouldSumWithSNR(t *testing.T) {
	assert := assert.New(t)
	first := []float64{0.5, 1, -1}
	second := []float64{0.5, 1, -1}
	sum := AddSignalToNoise(first, second, 20)
	assert.That(sum[0]).IsEqualTo(float64(0.55))
	assert.That(sum[1]).IsEqualTo(float64(1.1))
	assert.That(sum[2]).IsEqualTo(float64(-1.1))
}

func TestSumShouldSumWithSNRAndMaxAmplitudeOfSignal(t *testing.T) {
	assert := assert.New(t)
	first := []float64{1, 2, -2}
	second := []float64{0.5, 1, -1}
	sum := AddSignalToNoise(first, second, 20)
	assert.That(sum[0]).IsEqualTo(float64(1.1))
	assert.That(sum[1]).IsEqualTo(float64(2.2))
	assert.That(sum[2]).IsEqualTo(float64(-2.2))
}
