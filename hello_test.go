package HelloGo

import "testing"
import "github.com/assertgo/assert"

func TestThatBoolIsFalseHasNoErrors(t *testing.T) {
	assert := assert.New(t)
	assert.That(Tick()).IsNotNil()
}
