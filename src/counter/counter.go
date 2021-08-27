package counter

import "sync"

type Counter struct {
	value int
	mutex sync.RWMutex
}

func New(initialValue int) *Counter {
	return &Counter{
		value: initialValue,
		mutex: sync.RWMutex{},
	}
}

func (counter *Counter) Increment() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()

	counter.value++
}

func (counter *Counter) Value() int {
	counter.mutex.RLock()
	defer counter.mutex.RUnlock()

	return counter.value
}
