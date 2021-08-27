package greet

import (
	"fmt"
	"io"
)

// This is supposed to exemplify dependency injection
// NOTE: why not just make the function pure and test it?
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
