package syncedData

import (
	"sync"
)

type Counter struct {
	m			sync.Mutex
	count	int
}

func (c *Counter) Incr() {
	c.m.Lock()
	c.count++
	c.m.Unlock()
}

func (c *Counter) Read() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.count
}
