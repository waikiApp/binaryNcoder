binaryNCoder: Encode/Decode string to/from binary digits

#### Example
```golang
package main
import (
  bn "github.com/waikiApp/binaryNcoder"
  "fmt"
)

func main() {
  encodedStr := bn.Encode("example")
  fmt.Printf("%s\n", encodedStr)
  decodedStr := bn.Decode(encodedStr)
  fmt.Printf("%s\n", decodedStr)
}
```

`go run example.go`

`01100101011110000110000101101101011100000110110001100101
example`
