package cmd

import (
	"fmt"

	"strconv"

	"github.com/daniellowtw/xavier/cmd/service"
	"github.com/daniellowtw/xavier/db"
	"github.com/spf13/cobra"
)

var AddFilterCmd = &cobra.Command{
	Use:   "add-filter <id> <string> [delete|flag|read|save]",
	Short: "This will take a regex string and try to flag news that matches that filter",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return fmt.Errorf("incorrect arguments: %s", cmd.Use)
		}
		u := args[0]
		feedID, err := strconv.ParseInt(u, 10, 64)
		if err != nil {
			return err
		}
		str := args[1]
		action, err := mapToAction(args[2])
		if err != nil {
			return fmt.Errorf("unknown input")
		}
		s, err := service.NewServiceFromCmd(cmd)
		if err != nil {
			return err
		}
		return s.AddRegexFilter(feedID, str, action)
	},
}

func mapToAction(s string) (db.RuleAction, error) {
	switch s {
	case "flag":
		return db.Flag, nil
	case "read":
		return db.MarkAsRead, nil
	case "delete":
		return db.Delete, nil
	case "save":
		return db.Save, nil
	}
	return 0, fmt.Errorf("unknown")
}
