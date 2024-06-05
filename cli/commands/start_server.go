package commands

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var ServerStartCmd = &cobra.Command{
	Use:   "start server",
	Short: "Start the calculator server",
	Run:   startServer,
}

func startServer(cmd *cobra.Command, args []string) {
	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatalf(err.Error())
	}

	serverDir := filepath.ToSlash(filepath.Join(currentDir, "..", "server", "server"))
	command := exec.Command(serverDir)

	if err := command.Start(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}

	log.Printf("Server started successfully")
}
