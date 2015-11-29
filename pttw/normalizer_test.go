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