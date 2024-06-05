package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	_const "main/const"
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

func Post(cmd *cobra.Command, args []string, endpoint string) {
	data := args[0]
	requestBody := map[string]string{"expression": data}

	jsonReqBody, err := json.MarshalIndent(requestBody, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Request:\n%s\n", string(jsonReqBody))

	response, err := http.Post(_const.URL+endpoint, "application/json", bytes.NewBuffer(jsonReqBody))
	if err != nil {
		log.Fatalf("Failed to send POST request: %v", err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	PrintResponse(response, responseBody)
}
