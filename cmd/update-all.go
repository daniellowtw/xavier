package cmd

import (
	"github.com/spf13/cobra"
)

var UpdateAllCmd = &cobra.Command{
	Use: "update-all",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := newServiceFromCmd(cmd)
		_, err = s.UpdateAllFeeds()
		return err
	},
}
