/*
syncedData implements simple data types that are protected for concurrent use.
*/
package syncedData

import (
	"sync"
)

// Counter provides a simple counting interface that's made thread safe through
// use of a Mutex. Several methods are implemented to do typical functions.
// Incr() adds 1 to the counter. Decr() subtracts 1. Read() returns the
// current value of the Counter.
type Counter struct {
	m     sync.Mutex
	count int
}

func (c *Counter) Incr() {
	c.m.Lock()
	c.count++
	c.m.Unlock()
}

func (c *Counter) Decr() {
	c.m.Lock()
	c.count--
	c.m.Unlock()
}

func (c *Counter) Read() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.count
}
