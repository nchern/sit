package cli

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sit",
	Short: "Sit is a really efficient and a terminal-friendly issue tracker",
	// Long: TODO
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
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
