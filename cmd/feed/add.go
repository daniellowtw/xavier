package feed

import (
	"fmt"
	"net/url"

	"github.com/daniellowtw/xavier/cmd/service"
	"github.com/spf13/cobra"
)

//func NewAddCmd(s *api.Service) *cobra.Command {
var addCmd = &cobra.Command{
	Use: "add <url>",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Need one argument: <url>")
		}
		u := args[0]
		_, err := url.Parse(u)
		if err != nil {
			return err
		}
		s, err := service.NewServiceFromCmd(cmd)
		if err != nil {
			return err
		}
		return s.AddFeed(u)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
