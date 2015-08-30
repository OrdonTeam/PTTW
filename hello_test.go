package HelloGo

import "testing"
import "github.com/assertgo/assert"

func TestThatEmptyWorldStaysEmpty(t *testing.T) {
	assert := assert.New(t)
	assert.That(Tick(make(map[int]bool))).IsNotNil()
}
