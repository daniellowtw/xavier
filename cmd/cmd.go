package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use: os.Args[0],
	}
)

func init() {
	RootCmd.PersistentFlags().String("db", "data.db", "the file of the sqlite3 database")
	RootCmd.PersistentFlags().Bool("show-sql", false, "print out the sql statements")
	RootCmd.PersistentFlags().Int("log-level", 3, "log level (0 for debug to 5 for fatal)")
	RootCmd.AddCommand(
		AddCmd,
		UpdateAllCmd,
		DebugCmd,
		UpdateFeedCmd,
		ListAllCmd,
		DeleteCmd,
		NewWebCmd(),
		ImportFeedSourceFromFileCmd,
		NewLearnCmd(),
		AddFilterCmd,
	)
}
