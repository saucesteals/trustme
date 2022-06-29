# Trust your TLS/SSL environment

Example

```go
package main

import (
	"fmt"
	"strings"
	"github.com/saucesteals/trustme"
)

var (
	Version = "v0.0.0"
)

func main() {
	if !strings.HasSuffix(Version, "development") {
		trustme.Trust() // The error returned is generally safe to ignore
		fmt.Println("[!] Now in a safe environment")
	}
}
```
