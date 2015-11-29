package pttw

import (
	"github.com/assertgo/assert"
	"testing"
	"bytes"
)

func TestWriterShouldWriteSlice(t *testing.T) {
	assert := assert.New(t)
	var buffer bytes.Buffer
	Write([]float64{1}, &buffer)
	assert.That(true).IsEqualTo(true)
	assert.That(buffer.String()).IsEqualTo("1")
}
