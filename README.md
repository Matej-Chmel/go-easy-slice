# Easy slice
A simple library that extends the functionality of a built-in slice.

## Example
```go
package main

import (
    "fmt"

    gostack "github.com/Matej-Chmel/go-easy-slice"
)

func main() {
    s := eas.NewSlice[int32](4)
	s.Set(0, 4)
	s.Set(1, 8)
	s.Set(2, 20)
	s.Set(3, 87)

	fmt.Println(s.First())
	fmt.Println(s.Get(1))
	fmt.Println(s.Get(2))
	fmt.Println(s.Last())

	s.Append(17)
	s.AppendMore(78, 988, 901)

    repr := s.String()
    fmt.Println(repr) // [4 8 20 87 17 78 988 901]
}
```
