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
	slice := Generator(arrayBasedRandomProvider([]float64{0.1, 1}))(2)
	assert.That(slice[0]).IsEqualTo(float64(0.1))
	assert.That(slice[1]).IsEqualTo(float64(1))
}

//func TestGeneratorShouldNormalizeOutput(t *testing.T) {
//	assert := assert.New(t)
//	slice := Generator(arrayBasedRandomProvider([]float64{1, 2}))(2)
//	assert.That(slice[0]).IsEqualTo(float64(0.5))
//	assert.That(slice[1]).IsEqualTo(float64(1))
//}
//
//func TestGeneratorShouldNormalizeOutputWithNegativeValues(t *testing.T) {
//	assert := assert.New(t)
//	slice := Generator(arrayBasedRandomProvider([]float64{1, -4, 2}))(3)
//	assert.That(slice[0]).IsEqualTo(float64(0.25))
//	assert.That(slice[1]).IsEqualTo(float64(-1))
//	assert.That(slice[2]).IsEqualTo(float64(0.5))
//}

func arrayBasedRandomProvider(array []float64) func() float64 {
	index := 0
	return func() float64 {
		index++
		return array[index - 1]
	}
}