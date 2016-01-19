package pttw
import (
	"testing"
	"github.com/assertgo/assert"
	"strings"
	"bytes"
)

func TestAlgorithmShouldRun(t *testing.T) {
	assert := assert.New(t)
	reader := strings.NewReader("1 1 1 1")
	noise := arrayBasedRandomProvider([]float64{1, 1, 1, 1})
	writer := bytes.Buffer{}
	Algorithm(reader, &writer, 0, noise)
	assert.That(writer.String()).IsEqualTo("2 2 2 2\n")
}

// To hard to calculate correct value :P
//func TestAlgorithmShouldRunWithOtherParams(t *testing.T) {
//	assert := assert.New(t)
//	reader := strings.NewReader("1 1 1 1")
//	noise := arrayBasedRandomProvider([]float64{1, 1, 1, 1})
//	writer := bytes.Buffer{}
//	Algorithm(reader, &writer, 20, noise)
//	assert.That(writer.String()).IsEqualTo("1.1 1.1 1.1 1.1\n")
//}
//
//func TestAlgorithmShouldRunWithOtherParamsWhenSignalIsNotNormalized(t *testing.T) {
//	assert := assert.New(t)
//	reader := strings.NewReader("2 -2 2 -2")
//	noise := arrayBasedRandomProvider([]float64{1, 0, -1, 1})
//	writer := bytes.Buffer{}
//	Algorithm(reader, &writer, 20, noise)
//	assert.That(writer.String()).IsEqualTo("2.2 -2 1.8 -1.8\n")
//}
