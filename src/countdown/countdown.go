package countdown

import (
	"fmt"
	"io"
	"time"
)

func Count(writer io.Writer) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(writer, i)
	}

	fmt.Fprint(writer, "Go!")
}
