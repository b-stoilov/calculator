package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	_const "main/const"
	"main/utils"
	"net/http"
)

const url = "http://localhost:8080"

type EvaluateResponse struct {
	Result int `json:"result"`
}

type EvaluateRequest struct {
	Expression string `json:"expression"`
}

var EvaluateCmd = &cobra.Command{
	Use:   "evaluate [expression]",
	Short: "Evaluate a mathematical expression",
	Args:  cobra.ExactArgs(1),
	Run:   evaluate,
}

func evaluate(cmd *cobra.Command, args []string) {
	data := args[0]
	requestBody := map[string]string{"expression": data}

	jsonReqBody, err := json.MarshalIndent(requestBody, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post(_const.URL+"/evaluate",
		"application/json",
		bytes.NewBuffer(jsonReqBody))

	fmt.Printf("Request:\n%s\n", string(jsonReqBody))

	if err != nil {
		log.Fatalf("Failed to send POST request: %v", err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	utils.PrintResponse(response, responseBody)
}
