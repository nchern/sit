package cli

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
		Use:   "close",
		Short: "Closes a given issue",
		// TODO: add bulk close (?)
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			fmt.Println("not implemented: should close", id)
		},
	}

	issueReopenCmd = &cobra.Command{
		Use:   "reopen",
		Short: "Re-opens a given issue",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			fmt.Println("not implemented: should reopen", id)
		},
	}
)
