package main

import (
	"Calculator/store"
	"fmt"
	"net/http"
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

//
//func random() {
//	reader := bufio.NewReader(os.Stdin)
//
//	fmt.Println("What to calculate?")
//
//	calculateRequest, _ := reader.ReadString('\n')
//
//	fmt.Println("Calculating...")
//
//	if strings.Contains(calculateRequest, "plus") {
//		expression := extractor.ExtractNumbers(calculateRequest)
//		fmt.Println("Your result is: ", calculator.Add(expression.X, expression.Y))
//	} else {
//		fmt.Println("No valid operation found.")
//	}
//}
