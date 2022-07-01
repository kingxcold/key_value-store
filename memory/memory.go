package memory

import "errors"

type MemoryStore map[string]any

var InMemoryStore MemoryStore = make(MemoryStore)

func (s MemoryStore) Set(key string, value any) {
	s[key] = value
}

func (s MemoryStore) Get(key string) (any, error) {
	data, exist := s[key]
	if !exist {
		return "", errors.New("Data not found")
	}
	return data, nil
}

func (s MemoryStore) IsEmpty() bool {
	return len(s) == 0
}

func (s MemoryStore) Delete(key string) {
	delete(s, key)
}
