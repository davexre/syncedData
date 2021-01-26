package synceddata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testCounter(t *testing.T) {
	assert := assert.New(t)
	var myCounter Counter

	assert.Equal(0, myCounter.Read(), "Counter initialization not 0")

	myCounter.Incr()
	assert.Equal(1, myCounter.Read(), "Counter not incremented properly")
	myCounter.Incr()
	assert.Equal(2, myCounter.Read(), "Counter not incremented properly")
	myCounter.Decr()
	assert.Equal(1, myCounter.Read(), "Counter not decremented properly")
	myCounter.Set(5)
	assert.Equal(5, myCounter.Read(), "Counter not set properly")

}
