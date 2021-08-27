package counter

type Counter struct {
	value int
}

func New(initialValue int) *Counter {
	return &Counter{
		value: initialValue,
	}
}

func (counter *Counter) Increment() {
	counter.value++
}

func (counter *Counter) Value() int {
	return counter.value
}
