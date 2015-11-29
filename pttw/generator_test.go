package pttw
import (
	"github.com/assertgo/assert"
	"testing"
)

func TestGeneratorShouldGenerate(t *testing.T) {
	assert := assert.New(t)
	assert.ThatBool(true).IsTrue()
}