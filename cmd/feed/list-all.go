package feed

import (
	"fmt"

	"github.com/daniellowtw/xavier/cmd/service"
	"github.com/spf13/cobra"
)

var ListAllCmd = &cobra.Command{
	Use: "list",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := service.NewServiceFromCmd(cmd)
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

func init() {
	RootCmd.AddCommand(ListAllCmd)
}
