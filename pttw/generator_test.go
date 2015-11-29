package pttw
import (
	"github.com/assertgo/assert"
	"testing"
	"math/rand"
)

func TestGeneratorShouldGenerateRandomSliceInGivenSize(t *testing.T) {
	assert := assert.New(t)
	slice := Generator(rand.Float64)(5)
	assert.That(len(slice)).IsEqualTo(5)
}

func TestGeneratorShouldUseRandomFunction(t *testing.T) {
	assert := assert.New(t)
	slice := Generator(arrayBasedRandomProvider([]float64{1, 2}))(2)
	assert.That(slice[0]).IsEqualTo(float64(1))
	assert.That(slice[1]).IsEqualTo(float64(2))
}

func arrayBasedRandomProvider(array []float64) func() float64 {
	index := 0
	return func() float64 {
		index++
		return array[index-1]
	}
}