package feed

import (
	"github.com/daniellowtw/xavier/cmd/service"
	"github.com/spf13/cobra"
)

var UpdateAllCmd = &cobra.Command{
	Use: "update-all",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := service.NewServiceFromCmd(cmd)
		_, err = s.UpdateAllFeeds()
		return err
	},
}

func init() {
	RootCmd.AddCommand(UpdateAllCmd)
}
