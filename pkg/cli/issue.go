package cli

import (
	"fmt"

	"github.com/nchern/sit/pkg/issue"
	"github.com/nchern/sit/pkg/model"
	"github.com/spf13/cobra"
)

func init() {
	issueCmd.AddCommand(issueCloseCmd)
	issueCmd.AddCommand(issueOpenCmd)
	issueCmd.AddCommand(issueDoneCmd)

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
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completeIDs,

		RunE: func(cmd *cobra.Command, args []string) error {
			return changeTicketState(args[0], model.StateClosed)
		},
	}

	issueOpenCmd = &cobra.Command{
		Use:               "open",
		Short:             "Re-opens a given issue",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completeIDs,

		RunE: func(cmd *cobra.Command, args []string) error {
			return changeTicketState(args[0], model.StateOpen)
		},
	}

	issueDoneCmd = &cobra.Command{
		Use:               "done",
		Short:             "Makes a given issue done",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: completeIDs,

		RunE: func(cmd *cobra.Command, args []string) error {
			return changeTicketState(args[0], model.StateDone)
		},
	}
)

func changeTicketState(partialID string, state model.TicketState) error {
	ids, err := issue.FindIDs(partialID)
	if err != nil {
		return err
	}
	if len(ids) > 1 {
		return issue.NewMultipleResulsError(partialID, ids)
	}
	if len(ids) == 0 {
		return fmt.Errorf("No issues found with ids like '%s'", partialID)
	}
	id := ids[0]
	t, err := issue.FetchTicket(id)
	if err != nil {
		return err
	}
	t.State = state
	return issue.Update(t)
}
