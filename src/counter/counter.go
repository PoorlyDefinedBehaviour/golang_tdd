package counter

import "sync"

type T struct {
	value int
	mutex sync.RWMutex
}

func New(initialValue int) *T {
	return &T{
		value: initialValue,
		mutex: sync.RWMutex{},
	}
}

func (counter *T) Increment() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()

	counter.value++
}

func (counter *T) Value() int {
	counter.mutex.RLock()
	defer counter.mutex.RUnlock()

	return counter.value
}
