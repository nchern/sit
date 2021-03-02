package cmd

import (
	"fmt"
	"os"

	"github.com/nchern/sit/pkg/issue"
	"github.com/spf13/cobra"
)

func init() {
	issueCmd.AddCommand(issueCloseCmd)
	issueCmd.AddCommand(issueReopenCmd)
	issueCmd.AddCommand(issuePrintCmd)

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
			fmt.Println("not implemented: should close", id)
		},
	}

	issueReopenCmd = &cobra.Command{
		Use:  "reopen",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			fmt.Println("not implemented: should reopen", id)
		},
	}

	issuePrintCmd = &cobra.Command{
		Use:   "print",
		Args:  cobra.ExactArgs(1),
		Short: "Outputs a given issue",
		Run: func(cmd *cobra.Command, args []string) {
			must(issue.WriteByPartialID(args[0], os.Stdout))
		},
	}
)
