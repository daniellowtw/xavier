package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListAllCmd = &cobra.Command{
	Use: "list",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := newServiceFromCmd(cmd)
		if err != nil {
			return err
		}
		feeds, err := s.ListAllFeeds()
		if err != nil {
			return err
		}
		for _, f := range feeds {
			fmt.Println(f.Id, f.Title, f.Description)
		}
		return nil
	},
}
