package pttw
import (
	"testing"
	"github.com/assertgo/assert"
)

func TestParamsShouldReturnEmptyStringWhenNotFound(t *testing.T) {
	assert := assert.New(t)
	param := FindParam("-in", []string{})
	assert.That(param).IsEqualTo("")
}

func TestParamsShouldReturnNextElementFromArray(t *testing.T) {
	assert := assert.New(t)
	param := FindParam("-in", []string{"-in", "in"})
	assert.That(param).IsEqualTo("in")
}