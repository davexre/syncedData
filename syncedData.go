// Package synceddata implements simple data types that are protected for concurrent use. Currently, the
// package only contains a simple Counter, but will be extended in the future.
//
// Counter
//
// Counter provides a simple counting interface that's made thread safe through
// use of a Mutex. Several methods are implemented to do typical functions.
//
// Use:
//
// 	include "fmt"
// 	include "github.com/davexre/synceddata"
//
// 	func main() {
// 		var myCounter synceddata.Counter
//
// 		myCounter.Incr()
// 		myCounter.Incr()
// 		myCounter.Decr()
//
// 		fmt.Printf("Who's number %v?\n", myCounter.Read())
//
//		myCounter.Set(6)
//		myCounter.Incr()
//
//		fmt.Printf("%v is my lucky number!\n", myCounter.Read())
// 	}
package synceddata

import (
	"sync"
)

// Counter provides a simple counting interface that's made thread safe through
// use of a Mutex. Several methods are implemented to do typical functions.
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

func (c *Counter) Set(i int) {
	c.m.Lock()
	c.count = i
	c.m.Unlock()
}
