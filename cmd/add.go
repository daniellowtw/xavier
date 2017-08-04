package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

//func NewAddCmd(s *api.Service) *cobra.Command {
var AddCmd = &cobra.Command{
	Use: "add <url>",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Need one argument: <urL>")
		}
		u := args[0]
		_, err := url.Parse(u)
		if err != nil {
			return err
		}
		s, err := newServiceFromCmd(cmd)
		if err != nil {
			return err
		}
		return s.AddFeed(u)
	},
}
