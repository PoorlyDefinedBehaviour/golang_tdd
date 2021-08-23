package slices

func Sum(xs []int) int {
	out := 0

	for _, x := range xs {
		out += x
	}

	return out
}

func SumAll(xss ...[]int) int {
	out := 0

	for _, xs := range xss {
		out += Sum(xs)
	}

	return out
}
