package main

import (
	"cli/commands"
	"cli/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os/exec"
)

const url = "http://localhost:8080"

func main() {
	var rootCmd = &cobra.Command{Use: "ww"}

	//var serverStartCmd = &cobra.Command{
	//	Use:   "server start",
	//	Short: "Start HTTP server",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		log.Println("Server is running on port 8080")
	//		log.Fatal(http.ListenAndServe(":8080", nil))
	//
	//	},
	//}

	var getErrorCmd = &cobra.Command{
		Use:   "errors",
		Short: "HTTP GET request to fetch errors during validation and evaluation",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			response, err := http.Get(url + "/errors")

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
		},
	}

	var serverStartCmd = &cobra.Command{
		Use:   "start server",
		Short: "Run another Go application",
		Run: func(cmd *cobra.Command, args []string) {
			appPath := "C:/Users/boyan/GolandProjects/Calculator/server/server"

			command := exec.Command(appPath)

			if err := command.Start(); err != nil {
				log.Fatalf("Failed to run application: %v", err)
			}

			log.Printf("Server started successfully")
		},
	}

	rootCmd.AddCommand(serverStartCmd, getErrorCmd, commands.EvaluateCmd)

	err := rootCmd.Execute()
	if err != nil {
		return
	}

}
