package cmd

import (
	"fmt"

	"strconv"

	"github.com/spf13/cobra"
	"github.com/daniellowtw/xavier/db"
)

var AddFilterCmd = &cobra.Command{
	Use: "add-filter <id> <string>",
	Short: "This will take a regex string and try to flag news that matches that filter",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("Need two arguments: <id> <string>")
		}
		u := args[0]
		feedID, err := strconv.ParseInt(u, 10, 64)
		if err != nil {
			return err
		}
		str := args[1]
		s, err := newServiceFromCmd(cmd)
		if err != nil {
			return err
		}
		return s.AddRegexFilter(feedID, str, db.Flag)
	},
}
