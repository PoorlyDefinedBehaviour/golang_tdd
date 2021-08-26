package dictionary

type T map[string]string

func (dict *T) Search(word string) string {
	return (*dict)[word]
}
