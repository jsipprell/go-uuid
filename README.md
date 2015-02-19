Yet Another Go UUID Generator
=============================

_Because there can never be too many uuid generators or uuid types._

Usage
-----


```go
import "github.com/jsipprell/go-uuid"
import "fmt"

func main() {
  var id uuid.UUID

  id = uuid.New()
  fmt.Printf("first uuid is %v\n",id)

  id = uuid.New()
  fmt.Printf("second uuid is %v\n",id)

  id = uuid.New()
  fmt.Printf("third uuid is %v\n",id)
}
```

