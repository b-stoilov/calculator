package service

import (
	"Calculator/models"
	"Calculator/store"
)

func GetErrors(store *store.Store) []models.Error {
	return store.GetAllValues()
}
