package countdown

import (
	"fmt"
	"io"
)

func Count(writer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}

	fmt.Fprint(writer, "Go!")
}
