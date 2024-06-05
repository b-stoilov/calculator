package commands

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	_const "main/const"
	"main/utils"
	"net/http"
)

var ErrorsCmd = &cobra.Command{
	Use:   "errors",
	Short: "Get data for calculation errors",
	Args:  cobra.ExactArgs(0),
	Run:   fetchErrors,
}

func fetchErrors(cmd *cobra.Command, args []string) {
	response, err := http.Get(_const.URL + "/errors")

	if err != nil {
		log.Fatalf("Failed to send GET request: %v", err)
	}

	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	var jsonResponse []map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		log.Fatalf("Failed to parse JSON response: %v", err)
	}

	fmt.Printf("Response: %s\n", utils.PrettyPrintJSON(jsonResponse))
}
