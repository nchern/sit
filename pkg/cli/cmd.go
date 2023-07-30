package cli

import (
	"log"

	"github.com/nchern/sit/pkg/issue"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sit",
	Short: "Sit is a really efficient and terminal-friendly issue tracker",
	// Long: TODO
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},

	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	log.SetFlags(0)
}

// Execute is an entry point of CLI
func Execute() {
	must(rootCmd.Execute())
}

func must(err error) {
	if err != nil {
		log.Fatal("fatal: ", err)
	}
}

func completeIDs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	ids, err := issue.FindIDs(toComplete)
	if err != nil {
		log.Printf("error: %s", err)
		return nil, cobra.ShellCompDirectiveError
	}

	return ids, cobra.ShellCompDirectiveDefault
}
