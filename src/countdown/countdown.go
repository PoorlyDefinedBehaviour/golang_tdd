package countdown

import (
	"fmt"
	"io"
)

// NOTE:
// why create an interface for it?
// it is clearly just a function
type Sleeper interface {
	Sleep()
}

// we are injecting dependencies to make it
// testable
func Count(writer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}

	sleeper.Sleep()
	fmt.Fprint(writer, "Go!")
}
