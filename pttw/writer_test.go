package pttw

import (
	"github.com/assertgo/assert"
	"testing"
	"bytes"
)

func TestWriterShouldWriteSlice(t *testing.T) {
	assert := assert.New(t)
	buffer := bytes.Buffer{}
	Write([]float64{1}, &buffer)
	assert.That(true).IsEqualTo(true)
	assert.That(buffer.String()).IsEqualTo("1")
}

func TestWriterShouldWriteSliceDelimitedWithSpaces(t *testing.T) {
	assert := assert.New(t)
	buffer := bytes.Buffer{}
	Write([]float64{1, 2}, &buffer)
	assert.That(true).IsEqualTo(true)
	assert.That(buffer.String()).IsEqualTo("1 2")
}
