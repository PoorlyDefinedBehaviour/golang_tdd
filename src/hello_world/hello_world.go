package helloworld

import "fmt"

func Hello(name string) string {
	if name == "" {
		return "Hello, world"
	}

	return fmt.Sprintf("Hello, %s", name)
}
