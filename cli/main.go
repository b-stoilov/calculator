package main

import (
	"github.com/spf13/cobra"
	"main/commands"
)

func main() {
	var rootCmd = &cobra.Command{Use: "calc"}

	rootCmd.AddCommand(
		commands.ServerStartCmd,
		commands.ErrorsCmd,
		commands.EvaluateCmd,
		commands.ValidateCmd)

	err := rootCmd.Execute()
	if err != nil {
		return
	}

}
