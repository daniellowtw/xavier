package feed

import (
	"github.com/daniellowtw/xavier/cmd/service"
	"github.com/spf13/cobra"
	"fmt"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle <feedID>",
	Short: "Toggle a feed to be active or not",
	RunE: func(cmd *cobra.Command, args []string) error {
		i, err := extractInt1(args, "feed ID")
		if err != nil {
			return err
		}
		s, err := service.NewServiceFromCmd(cmd)
		if err != nil {
			return err
		}
		isActive, err := s.ToggleActive(i)
		if err != nil {
			return err
		}
		fmt.Printf("Feed %d is now Active: %t\n", i, isActive)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(toggleCmd)
}
