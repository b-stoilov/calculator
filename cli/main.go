package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

const url = "localhost:8080"

func main() {
	var rootCmd = &cobra.Command{Use: "cli"}

	fmt.Println("cli started")

	var getCmd = &cobra.Command{
		Use:   "errors",
		Short: "Send a GET request",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			response, err := http.Get(url)

			if err != nil {
				log.Fatalf("Failed to send GET request: %v", err)
			}

			defer response.Body.Close()
			body, _ := io.ReadAll(response.Body)

			var jsonResponse map[string]interface{}
			if err := json.Unmarshal(body, &jsonResponse); err != nil {
				log.Fatalf("Failed to parse JSON response: %v", err)
			}

			fmt.Printf("Response: %s\n", prettyPrintJSON(jsonResponse))
		},
	}

	rootCmd.AddCommand(getCmd)
	err := rootCmd.Execute()
	if err != nil {
		return
	}

}

func prettyPrintJSON(jsonResponse map[string]interface{}) string {
	jsonFormatted, err := json.MarshalIndent(jsonResponse, "", "  ")

	if err != nil {
		log.Fatalf("Failed to generate pretty JSON: %v", err)
	}

	return string(jsonFormatted)
}
