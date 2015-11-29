package pttw
import (
	"testing"
	"github.com/assertgo/assert"
)

func TestDBToLinearShouldConvert0(t *testing.T) {
	assert := assert.New(t)
	assert.That(SnrToLinear(0)).IsEqualTo(float64(1))
}

func TestDBToLinearShouldConvert20(t *testing.T) {
	assert := assert.New(t)
	assert.That(SnrToLinear(20)).IsEqualTo(float64(10))
}