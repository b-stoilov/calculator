package commands

import (
	"github.com/spf13/cobra"
	"main/utils"
)

var EvaluateCmd = &cobra.Command{
	Use:   "evaluate [expression]",
	Short: "Evaluate a mathematical expression",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.Post(cmd, args, "/evaluate")
	},
}
