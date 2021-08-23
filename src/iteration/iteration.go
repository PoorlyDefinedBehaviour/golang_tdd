package iteration

import "strings"

func Repeat(s string, times int) string {
	builder := strings.Builder{}

	for i := 0; i < times; i++ {
		_, _ = builder.WriteString(s)
	}

	return builder.String()
}
