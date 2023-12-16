package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {

		

		fmt.Println("Added" + strings.Join(args[:]," ") + "to your task list.")
	},
}

func init() {
	rootCmd.AddCommand(add)
}
