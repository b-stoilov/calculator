package main

import (
	"encoding/json"
	"net/http"
	"server/service"
	"server/store"
)

type ExpressionRequestDTO struct {
	Expression string `json:"expression"`
}

func evaluateHandler(store *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ExpressionRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		resp, err := service.Evaluate(req.Expression, store, r.URL.String())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}
}

func validateHandler(store *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ExpressionRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		resp := service.Validate(req.Expression, store, r.URL.String())

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}
}

func errorsHandler(store *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := service.GetErrors(store)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}
}
