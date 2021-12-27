# Paddle

## Synopsis

Paddle is a utility for padding data to the closest length that is a multiple of the provided block amount.

## Specification

### Padding

1. The data has a byte with a value of one appended to it.
2. The data has zero-ed bytes appended to it until the length of the data is a multiple of the block.

### Unpadding

Remove padding by removing zero-ed bytes from the end of the data until the byte with a value of one is found.

## Distribution

Paddle is released as a package written in [Go](https://golang.org) under the MIT License. Paddle is currently at version 1.0.0, and will likely undergo an update that adds a constant-time implementation in order to be resistant to side-channel attacks.

## Example

```
package main
import (
	"fmt"
	"github.com/cs-bic/paddle"
)
func main() {
	block := 100
	data := []byte("Hello, world!")
	fmt.Println("ORIGINAL LENGTH: ", len(data))
	fmt.Println("ORIGINAL: ", data)
	data, issue := paddle.Pad(block, data)
	if issue != nil {
		panic(issue)
	}
	fmt.Println("PADDED LENGTH: ", len(data))
	fmt.Println("PADDED: ", data)
	data, issue = paddle.Unpad(data)
	if issue != nil {
		panic(issue)
	}
	fmt.Println("UNPADDED LENGTH: ", len(data))
	fmt.Println("UNPADDED: ", data)
}
```
