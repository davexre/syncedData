# syncedData

syncedData implements simple data types that are protected for concurrent use. Currently, the
package only contains a simple Counter, but will be extended in the future.

## Types

### type [Counter](/syncedData.go#L26)

`type Counter struct { ... }`

Counter provides a simple counting interface that's made thread safe through
use of a Mutex. Several methods are implemented to do typical functions.

Use:

```go
include "fmt"
include "github.com/davexre/syncedData"

func main() {
	var myCounter syncedData.Counter

	myCounter.Incr()
	myCounter.Incr()
	myCounter.Decr()

	fmt.Println("Who's number %!(NOVERB)v?", myCounter.Read())
}
```

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)