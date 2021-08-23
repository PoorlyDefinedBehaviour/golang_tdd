package slices

func Sum(xs []int) int {
	out := 0

	for _, x := range xs {
		out += x
	}

	return out
}
