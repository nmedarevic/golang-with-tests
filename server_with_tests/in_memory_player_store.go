package main

type InMemoryStore struct {
	store map[string]int
}

func (store *InMemoryStore) GetPlayerScore(name string) int {
	return store.store[name]

}

func (store *InMemoryStore) RecordScore(name string) {
	store.store[name]++
}

func NewInMemoryPlayerStore() *InMemoryStore {
	return &InMemoryStore{map[string]int{}}
}
