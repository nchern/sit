package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	issueCmd.AddCommand(issueCloseCmd)
	issueCmd.AddCommand(issueReopenCmd)

	rootCmd.AddCommand(issueCmd)
}

var (
	issueCmd = &cobra.Command{
		Use:   "issue",
		Short: "A subcommand to work with individual issues",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	issueCloseCmd = &cobra.Command{
		Use:  "close",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			fmt.Println("Should close", id)
		},
	}

	issueReopenCmd = &cobra.Command{
		Use:  "reopen",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			fmt.Println("Should reopen", id)
		},
	}
)
