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
	first := []float64{1}
	second := []float64{1}
	sum := AddSignalToNoise(first, second, 20)
	assert.That((sum[0] - 1) * (sum[0] - 1) < 0.1000001).IsEqualTo(true)
	assert.That((sum[0] - 1) * (sum[0] - 1) > 0.0999999).IsEqualTo(true)
}

func TestSumShouldSumWithSNRAndMaxAmplitudeOfSignal(t *testing.T) {
	assert := assert.New(t)
	first := []float64{2}
	second := []float64{-6}
	sum := AddSignalToNoise(first, second, 20)
	assert.That((sum[0] - 2) * (sum[0] - 2) < 0.4000001).IsEqualTo(true)
	assert.That((sum[0] - 2) * (sum[0] - 2) > 0.3999999).IsEqualTo(true)
}
