package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use: "delete <id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Need one argument: <urL>")
		}
		u := args[0]
		i, err := strconv.Atoi(u)
		if err != nil {
			return err
		}
		s, err := newServiceFromCmd(cmd)
		if err != nil {
			return err
		}
		return s.DeleteFeed(int64(i))
	},
}
