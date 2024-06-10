package store

import "server/models"

type Store struct {
	data map[string]models.Error
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]models.Error),
	}
}

func (store *Store) Set(key string, value models.Error) {
	if entry, exists := store.data[key]; exists {
		entry.Frequency += 1
		value = entry
	}

	store.data[key] = value
}

func (store *Store) Get(key string) models.Error {
	//value, exists :=
	return store.data[key]
}

func (store *Store) GetAllValues() []models.Error {
	var values []models.Error

	for _, entry := range store.data {
		values = append(values, entry)
	}

	return values
}

//func (info *Store) Exists(key string) bool {
//	_, exists := info.data[key]
//}
