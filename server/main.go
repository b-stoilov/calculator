package main

import (
	"fmt"
	"net/http"
	"server/store"
)

func main() {
	store := store.NewStore()

	http.HandleFunc("/evaluate", evaluateHandler(store))
	http.HandleFunc("/validate", validateHandler(store))
	http.HandleFunc("/errors", errorsHandler(store))

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

}
