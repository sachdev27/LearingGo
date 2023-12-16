package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// list represents the add command
var list = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			fmt.Println("list doesn't require any argument")
			return
		}


	},
}

func init() {
	rootCmd.AddCommand(list)
}
