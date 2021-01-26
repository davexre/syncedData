# synceddata

Package synceddata implements simple data types that are protected for concurrent use. Currently, the
package only contains a simple Counter, but will be extended in the future.

## Counter

Counter provides a simple counting interface that's made thread safe through
use of a Mutex. Several methods are implemented to do typical functions.

Use:

```go
include "fmt"
include "github.com/davexre/synceddata"

func main() {
	var myCounter synceddata.Counter

	myCounter.Incr()
	myCounter.Incr()
	myCounter.Decr()

	fmt.Printf("Who's number %v?\n", myCounter.Read())

	myCounter.Set(6)
	myCounter.Incr()

	fmt.Printf("%v is my lucky number!\n", myCounter.Read())
}
```

