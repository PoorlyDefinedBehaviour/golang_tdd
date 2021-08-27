package dictionary

type T map[string]string

func (dict T) Search(key string) string {
	return dict[key]
}

func (dict T) Contains(key string) bool {
	_, ok := dict[key]

	return ok
}

func (dict T) Remove(key string) {
	delete(dict, key)
}
