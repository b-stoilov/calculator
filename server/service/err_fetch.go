package service

import (
	"server/models"
	"server/store"
)

func GetErrors(store *store.Store) []models.Error {
	return store.GetAllValues()
}
