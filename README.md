# cache

## Example #1

```go
package main

import (
"fmt"
"github.com/Nurik2k/cache"
)

func main() {

	cache := cache.NewCache()

	cache.Set("userId", 42)
	userId, _ := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId, _ = cache.Get("userId")

	fmt.Println(userId)

}
```