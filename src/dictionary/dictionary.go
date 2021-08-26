package dictionary

type T map[string]string

func (dict *T) Search(word string) string {
	return (*dict)[word]
}

func (dict *T) Contains(word string) bool {
	_, ok := (*dict)[word]

	return ok
}
