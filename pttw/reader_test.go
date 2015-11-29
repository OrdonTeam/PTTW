package pttw

import (
	"github.com/assertgo/assert"
	"testing"
	"strings"
)

func TestReaderShouldReadFloatSlice(t *testing.T) {
	assert := assert.New(t)
	reader := strings.NewReader("4 5")
	read, _ := Read(reader)
	assert.That(read[0]).IsEqualTo(float64(4))
}

func TestReaderShouldReadFloatSliceWithTowValues(t *testing.T) {
	assert := assert.New(t)
	reader := strings.NewReader("4 5")
	read, _ := Read(reader)
	assert.That(read[0]).IsEqualTo(float64(4))
	assert.That(read[1]).IsEqualTo(float64(5))
}

func TestReaderShouldReadError(t *testing.T) {
	assert := assert.New(t)
	reader := strings.NewReader("")
	_, err := Read(reader)
	assert.ThatBool(err == nil).IsFalse()
}
