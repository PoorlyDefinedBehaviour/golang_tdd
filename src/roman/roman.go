package roman

import "strings"

func FromArabic(n int) string {
	builder := strings.Builder{}

	for i := 0; i < n; i++ {
		builder.WriteRune('I')
	}

	return builder.String()
}
