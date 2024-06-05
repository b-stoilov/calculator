package commands

import (
	"github.com/spf13/cobra"
	"main/utils"
)

var ValidateCmd = &cobra.Command{
	Use:   "validate [expression]",
	Short: "Validate a mathematical expression",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.Post(cmd, args, "/validate")
	},
}
