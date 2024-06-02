package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/evaluate", evaluateHandler)
	http.HandleFunc("/validate", validateHandler)
	http.HandleFunc("/errors", errorsHandler)

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
