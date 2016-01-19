package pttw
import (
	"testing"
	"github.com/assertgo/assert"
)

func TestPowerShouldReturnPowerOfSingleSample(t *testing.T) {
	assert := assert.New(t)
	assert.That(PowerOfSignal([]float64{1.0})).IsEqualTo(1.0)
}

func TestPowerShouldReturnPowerOfSingleSampleOf2(t *testing.T) {
	assert := assert.New(t)
	assert.That(PowerOfSignal([]float64{2.0})).IsEqualTo(4.0)
}

func TestPowerShouldReturnAveragePowerOfTwoSamples(t *testing.T) {
	assert := assert.New(t)
	assert.That(PowerOfSignal([]float64{1.0, 2.0})).IsEqualTo(2.5)
}