package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func PrintResponse(response *http.Response, responseBody []byte) {
	if response.StatusCode == http.StatusOK {
		var jsonResponse interface{}
		if err := json.Unmarshal(responseBody, &jsonResponse); err == nil {
			fmt.Printf("Response:\n%s\n", PrettyPrintJSON(jsonResponse))
		} else {
			fmt.Printf("Failed to parse JSON response: %v\n", err)
		}
	} else {
		fmt.Printf("Response (Text format if response code 500): %s\n", string(responseBody))
	}
}

func PrettyPrintJSON(jsonResponse interface{}) string {
	jsonFormatted, err := json.MarshalIndent(jsonResponse, "", "  ")

	if err != nil {
		log.Fatalf("Failed to generate pretty JSON: %v", err)
	}

	return string(jsonFormatted)
}
